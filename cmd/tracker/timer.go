package tracker

// USED FOR RUNNING TIME TASKS

import (
	"time"
)

type Timer struct {
	start    time.Time
	duration time.Duration
	name     string
	tracker  *Tracker
}

func NewTimer(name string, tracker *Tracker) *Timer {
	return &Timer{
		start:    time.Time{},
		duration: time.Duration(0),
		name:     name,
		tracker:  tracker,
	}
}

func (t *Timer) StartTimer() {
	t.start = time.Now()
}

func (t *Timer) EndTimer() time.Duration {
	t.duration = time.Since(t.start)
	t.tracker.addRun(t)
	return time.Since(t.start)
}
