package controller

import (
	"sync"
)

const (
	WORKERS = 4
	JOB_QUEUE_MAX = 200
	MAX_HOLDING_JOB = 1000
)

type jobItem struct {
	id    string
	code  string
	stdin string
}

type jobState struct {
	status string // queued|running|finished|failed
	stdout string
	stderr string
}

var (
	onceStart sync.Once
	queue chan jobItem

	storeMu sync.RWMutex
	store = map[string]jobState{}
	jobOrder []string
)

func StartWorkers() {
	onceStart.Do(func() {
		queue = make(chan jobItem, JOB_QUEUE_MAX)
		for boxID := 0; boxID < WORKERS; boxID++ {
			go workerLoop(boxID)
		}
	})
}

// EnqueueJob returns false if queue is full.
func EnqueueJob(jobID, code, stdin string) bool {
	select {
	case queue <- jobItem{id: jobID, code: code, stdin: stdin}:
		return true
	default:
		return false
	}
}

func workerLoop(boxID int) {
	for it := range queue {
		jobStoreUpdate(it.id, func(s jobState) jobState {
			s.status = "running"
			return s
		})

		out, errText := runCInIsolate(boxID, it.code, it.stdin)

		jobStoreUpdate(it.id, func(s jobState) jobState {
			if errText != "" {
				s.status = "failed"
				s.stderr = errText
				s.stdout = out
			} else {
				s.status = "finished"
				s.stdout = out
				s.stderr = ""
			}
			return s
		})
	}
}

/* ---- store helpers (private) ---- */

func jobStoreGet(id string) (jobState, bool) {
	storeMu.RLock()
	defer storeMu.RUnlock()
	s, ok := store[id]
	return s, ok
}

func jobStoreSet(id string, s jobState) {
	storeMu.Lock()
	defer storeMu.Unlock()

	// If new job ID, append to order
	if _, exists := store[id]; !exists {
		jobOrder = append(jobOrder, id)

		// Evict oldest if exceed limit
		if len(jobOrder) > MAX_HOLDING_JOB {
			oldest := jobOrder[0]
			jobOrder = jobOrder[1:]
			delete(store, oldest)
		}
	}

	store[id] = s
}

func jobStoreDelete(id string) {
	storeMu.Lock()
	defer storeMu.Unlock()

	delete(store, id)

	for i, v := range jobOrder {
		if v == id {
			jobOrder = append(jobOrder[:i], jobOrder[i+1:]...)
			break
		}
	}
}

func jobStoreUpdate(id string, fn func(jobState) jobState) {
	storeMu.Lock()
	defer storeMu.Unlock()

	s := store[id]
	store[id] = fn(s)
}
