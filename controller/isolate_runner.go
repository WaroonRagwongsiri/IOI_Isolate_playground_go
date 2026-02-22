package controller

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	ISOLATE = "isolate"
	CC = "/usr/bin/gcc"
)

// runCInIsolate returns (stdout, stderrText).
// If compile fails, stderrText will contain the compile error.
// If runtime writes to stderr, stderrText will contain that runtime stderr.
func runCInIsolate(boxID int, code string, stdin string) (string, string) {
	// cleanup -> init
	_ = exec.Command(ISOLATE, fmt.Sprintf("--box-id=%d", boxID), "--cleanup").Run()

	initCmd := exec.Command(ISOLATE, fmt.Sprintf("--box-id=%d", boxID), "--init")
	var initErr bytes.Buffer
	initCmd.Stderr = &initErr
	if err := initCmd.Run(); err != nil {
		if initErr.String() != "" {
			return "", initErr.String()
		}
		return "", err.Error()
	}
	defer func() {
		_ = exec.Command(ISOLATE, fmt.Sprintf("--box-id=%d", boxID), "--cleanup").Run()
	}()

	box := fmt.Sprintf("/var/local/lib/isolate/%d/box", boxID)

	// write code + optional stdin to the box dir
	if err := os.WriteFile(filepath.Join(box, "main.c"), []byte(code), 0644); err != nil {
		return "", err.Error()
	}
	hasStdin := stdin != ""
	if hasStdin {
		if err := os.WriteFile(filepath.Join(box, "input.txt"), []byte(stdin), 0644); err != nil {
			return "", err.Error()
		}
	}

	// compile inside isolate
	compileCmd := exec.Command(
		ISOLATE,
		fmt.Sprintf("--box-id=%d", boxID),
		"--processes=64",
		"--run",
		"--env=PATH=/usr/bin:/bin",
		"--chdir=/box",
		"--",
		CC,
		"main.c",
		"-o",
		"main",
		"-B/usr/bin",
	)
	var compErr bytes.Buffer
	compileCmd.Stderr = &compErr
	if err := compileCmd.Run(); err != nil {
		if compErr.String() != "" {
			return "", compErr.String()
		}
		return "", err.Error()
	}

	// ensure output exists
	if _, err := os.Stat(filepath.Join(box, "main")); err != nil {
		return "", "compile failed: output binary not found"
	}

	// run inside isolate; redirect program output to files
	args := []string{
		fmt.Sprintf("--box-id=%d", boxID),
		"--time=1",
		"--mem=262144",
		"--stdout=out.txt",
		"--stderr=err.txt",
	}
	if hasStdin {
		args = append(args, "--stdin=input.txt")
	}
	args = append(args, "--run", "--", "./main")

	_ = exec.Command(ISOLATE, args...).Run()

	outB, _ := os.ReadFile(filepath.Join(box, "out.txt"))
	errB, _ := os.ReadFile(filepath.Join(box, "err.txt"))

	return string(outB), string(errB)
}
