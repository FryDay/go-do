package godo

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

// ToDo represents a task
type ToDo struct {
	Name          string    `json:"name,omitempty"`
	Note          string    `json:"note,omitempty"`
	Priority      uint8     `json:"priority,omitempty"`
	TimeCreated   time.Time `json:"time_created,omitempty"`
	TimeCompleted time.Time `json:"time_completed,omitempty"`
}

// NewToDo takes a name and note and returns a ToDo
func NewToDo(name, note string) *ToDo {
	return &ToDo{
		Name:        name,
		Note:        note,
		Priority:    0,
		TimeCreated: time.Now(),
	}
}

// Edit ToDo's name and note
func (td *ToDo) Edit(name, note string) {
	td.Name = name
	td.Note = note
}

// Up a ToDo's priority
func (td *ToDo) Up() {
	log.Println(td.Priority)
	td.Priority++
	log.Println(td.Priority)
}

// Down a ToDo's priority
func (td *ToDo) Down() {
	if td.Priority == 0 {
		return
	}
	td.Priority--
}

// Complete sets TimeCompleted to the current time
func (td *ToDo) Complete() {
	td.TimeCompleted = time.Now()
}

// Reopen clears TimeCompleted
func (td *ToDo) Reopen() {
	td.TimeCompleted = time.Time{}
}

// Bytes returns ToDo as a json formatted slice of bytes
func (td *ToDo) Bytes() []byte {
	b, _ := json.Marshal(*td)
	return b
}

// ToString ...
func (td *ToDo) ToString() string {
	s := td.Name

	if td.Priority > 0 {
		s += " ("
		for i := uint8(0); i < td.Priority; i++ {
			s += "+"
		}
		s += ")"
	}

	return s
}

// ToDos is a list of ToDo
type ToDos []*ToDo

// Add a Todo
func (tds ToDos) Add(todo *ToDo) ToDos {
	tds = append(tds, todo)
	return tds
}

// Delete a ToDo by index
func (tds ToDos) Delete(i int) ToDos {
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
	if !tds[i].TimeCompleted.IsZero() && !tds[j].TimeCompleted.IsZero() {
		iName := strings.ToLower(tds[i].Name)
		jName := strings.ToLower(tds[j].Name)
		if iName > jName {
			return false
		}
		if iName < jName {
			return true
		}
	}
	if tds[i].TimeCompleted.After(tds[j].TimeCompleted) {
		return false
	}
	if tds[i].TimeCompleted.Before(tds[j].TimeCompleted) {
		return true
	}

	if tds[i].Priority < tds[j].Priority {
		return false
	}
	if tds[i].Priority > tds[j].Priority {
		return true
	}

	iName := strings.ToLower(tds[i].Name)
	jName := strings.ToLower(tds[j].Name)
	if iName > jName {
		return false
	}
	if iName < jName {
		return true
	}

	return tds[i].TimeCreated.Before(tds[j].TimeCreated)
}
