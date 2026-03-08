package service

import "time"

// Statistics holds the result metrics from an import operation.
type Statistics struct {
	TimeElapsed time.Duration
	Accepted    int64
	Discarded   int64
}
