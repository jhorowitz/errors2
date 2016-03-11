package errors2

import "time"

type LogMessage struct {
	time    time.Time
	message string
}