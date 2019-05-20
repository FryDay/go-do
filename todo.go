package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// ToDo represents a task
type ToDo struct {
	Name          string    `json:"name,omitempty"`
	TimeCreated   time.Time `json:"time_created,omitempty"`
	TimeCompleted time.Time `json:"time_completed,omitempty"`
}

// NewToDo takes a name and returns a ToDo
func NewToDo(name string) *ToDo {
	return &ToDo{
		Name:        name,
		TimeCreated: time.Now(),
	}
}

// Edit ToDo
func (td *ToDo) Edit(name string) {
	td.Name = name
	go td.WriteFile()
}

// Complete sets TimeCompleted to the current time
func (td *ToDo) Complete() {
	td.TimeCompleted = time.Now()
	go td.WriteFile()
}

// Reopen clears TimeCompleted
func (td *ToDo) Reopen() {
	td.TimeCompleted = time.Time{}
	go td.WriteFile()
}

// Filename returns a filename based on ToDo values
func (td *ToDo) Filename() string {
	return filepath.Join(configDir, fmt.Sprintf("go-do_%s.json", td.TimeCreated.Format("060102150405.999")))
}

// WriteFile ...
func (td *ToDo) WriteFile() {
	b, _ := json.Marshal(td)
	ioutil.WriteFile(td.Filename(), b, 0600)
}

// DeleteFile ...
func (td *ToDo) DeleteFile() {
	os.Remove(td.Filename())
}

// ToDos is a list of ToDo
type ToDos []*ToDo

// Add a Todo
func (tds ToDos) Add(todo *ToDo) ToDos {
	go todo.WriteFile()

	tds = append(tds, todo)
	return tds
}

// Delete a ToDo by index
func (tds ToDos) Delete(i int) ToDos {
	go tds[i].DeleteFile()

	copy(tds[i:], tds[i+1:])
	tds[len(tds)-1] = nil
	return tds[:len(tds)-1]
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
