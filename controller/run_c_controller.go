package controller

import (
	"github.com/google/uuid"
)

// RunCController: enqueue job, return job_id
func RunCController(req RunC) RunCResponse {
	jobID := uuid.NewString()

	// store initial job state
	jobStoreSet(jobID, jobState{status: "queued"})

	// enqueue (no package-level error var)
	if !EnqueueJob(jobID, req.Code, req.Stdin) {
		jobStoreDelete(jobID)
		return RunCResponse{JobId: ""}
	}

	return RunCResponse{JobId: jobID}
}

// JobFromIDController: return stdout/stderr only
func JobFromIDController(req JobFromId) JobFromIdResponse {
	j, ok := jobStoreGet(req.JobId)
	if !ok {
		return JobFromIdResponse{Stdout: "", Stderr: "job not found"}
	}
	return JobFromIdResponse{Stdout: j.stdout, Stderr: j.stderr}
}
