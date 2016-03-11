package errors2

import (
	"sync"
	"time"
)

const (
	lvldbug = iota
	lvlinfo
	lvlerr
	lvlcrit
)

type logMessage struct {
	time    time.Time
	level   int
	message string
}

type Logger struct {
	messages []*logMessage
	Timer    Timer
}

func (l *Logger) Debug(message string) {
	l.log(&logMessage{time: time.Now(), level: lvldbug, message: message})
}

func (l *Logger) Info(message string) {
	l.log(&logMessage{time: time.Now(), level: lvlinfo, message: message})
}

func (l *Logger) Error(message string) {
	l.log(&logMessage{time: time.Now(), level: lvlerr, message: message})
}

func (l *Logger) Critical(message string) {
	l.log(&logMessage{time: time.Now(), level: lvlcrit, message: message})
}

var lock = &sync.Mutex{}

func (l *Logger) log(message *logMessage) {
	lock.Lock()
	l.messages = append(l.messages, message)
	lock.Unlock()
}
