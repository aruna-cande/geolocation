package service

import "time"

type Statistics struct {
	TimeElapsed time.Duration
	Accepted int64
	Discarded int64
}