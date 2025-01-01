package tracker

// USED JUST FOR HOLDING DATA

import "time"

type Tracker struct {
	Runs    []*Timer
	Name    string
	Amount  int
	Average time.Duration
}

func NewTracker(name string) *Tracker {
	return &Tracker{
		Runs:    []*Timer{},
		Name:    name,
		Amount:  0,
		Average: time.Duration(0),
	}
}

func (t *Tracker) addRun(timer *Timer) {
	t.Runs = append(t.Runs, timer)
}
