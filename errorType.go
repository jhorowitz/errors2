package errors2

import (
	"time"
	"fmt"
)

type LogMessage struct {
	time    time.Time
	message string
}

type FunctionTracker struct {
	name               string
	startTime, endTime time.Time
	started, ended     bool
}

func (f FunctionTracker) Report () (report string) {
	if len(f.name) == 0 {
		f.name = "anonymous"
	}
	if !f.started {
		return "Tracking was never started."
	}
	if !f.ended {
		return fmt.Printf("Tracking was never completed for %s. It's been %s since it started.", f.name, time.Since(f.startTime).String())
	}
	report += fmt.Printf("[%s]: completed in %s", f.name, f.endTime.Sub(f.startTime).String())
	return
}

func (f FunctionTracker) Begin(name string) {
	f.started = true
	f.startTime = time.Now()
}

func (f FunctionTracker) Complete(name string) {
	f.ended = true
	f.endTime = time.Now()
}

type Error struct {
	failureReason string
	failed        bool
}

func (e Error) Fail(reason string) {
	e.failed = true
	e.failureReason = reason
}

func (e Error) ToError() error {
	if !e.failed {
		return nil
	}
	return e
}

func (e Error) Error() string {
	return e.failureReason
}
