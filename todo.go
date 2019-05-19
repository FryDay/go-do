package main

import (
	"time"
)

// ToDo represents a task
type ToDo struct {
	Name          string    `json:"name,omitempty"`
	TimeCreated   time.Time `json:"time_created,omitempty"`
	TimeCompleted time.Time `json:"time_completed,omitempty"`
}

// ToDos is a list of ToDo
type ToDos []*ToDo

// NewToDo takes a name and returns a ToDo
func NewToDo(name string) *ToDo {
	return &ToDo{
		Name:        name,
		TimeCreated: time.Now(),
	}
}

// Complete sets TimeCompleted to the current time
func (td *ToDo) Complete() {
	td.TimeCompleted = time.Now()
}

// Reopen clears TimeCompleted
func (td *ToDo) Reopen() {
	td.TimeCompleted = time.Time{}
}

func (tds ToDos) Len() int {
	return len(tds)
}

func (tds ToDos) Swap(i, j int) {
	tds[i], tds[j] = tds[j], tds[i]
}

func (tds ToDos) Less(i, j int) bool {
	if tds[i].TimeCompleted.After(tds[j].TimeCompleted) {
		return false
	}
	if tds[i].TimeCompleted.Before(tds[j].TimeCompleted) {
		return true
	}

	return tds[i].TimeCreated.Before(tds[j].TimeCreated)
}
