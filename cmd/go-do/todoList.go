package main

import (
	"fmt"
	"sort"

	godo "github.com/FryDay/go-do"
	"github.com/gdamore/tcell/v2"
	"gitlab.com/tslocum/cview"
)

// ToDoList ...
type ToDoList struct {
	*cview.Box
	ToDos       godo.ToDos
	CurrentToDo int
}

// NewToDoList returns a new todo list.
func NewToDoList() *ToDoList {
	return &ToDoList{
		Box: cview.NewBox(),
	}
}

// InputHandler ...
func (t *ToDoList) InputHandler() func(event *tcell.EventKey, setFocus func(p cview.Primitive)) {
	return t.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p cview.Primitive)) {
		switch event.Key() {
		case tcell.KeyUp:
			t.selectUp()
		case tcell.KeyDown:
			t.selectDown()
		}
		switch event.Rune() {
		case 'k':
			t.selectUp()
		case 'j':
			t.selectDown()
		case 'K':
			td := t.ToDos[t.CurrentToDo]
			td.Up()
			writeToDoFile(td)
			t.sort()
		case 'J':
			td := t.ToDos[t.CurrentToDo]
			td.Down()
			writeToDoFile(td)
			t.sort()
		case ' ':
			td := t.ToDos[t.CurrentToDo]
			if td.TimeCompleted.IsZero() {
				td.Complete()
				writeToDoFile(td)
			} else {
				td.Reopen()
				writeToDoFile(td)
			}
			t.CurrentToDo = 0
			t.sort()
		}
	})
}

// Draw ...
func (t *ToDoList) Draw(screen tcell.Screen) {
	t.Box.Draw(screen)
	x, y, width, height := t.GetInnerRect()

	for i, td := range t.ToDos {
		if i >= height {
			break
		}
		checkbox := "[lime]\u2610"
		if !td.TimeCompleted.IsZero() {
			checkbox = "[red]\u2612"
		}
		line := fmt.Sprintf(` %s[default]   %s`, checkbox, td.ToString())
		if i == t.CurrentToDo {
			line = fmt.Sprintf(` %s[default]   [::r]%s`, checkbox, td.ToString())
		}
		cview.Print(screen, []byte(line), x, y+i, width, cview.AlignLeft, tcell.ColorDefault)
	}
}

// Add a new item
func (t *ToDoList) Add(item *godo.ToDo) {
	t.ToDos = t.ToDos.Add(item)
	writeToDoFile(item)
	t.sort()
}

// Delete the currently selected item
func (t *ToDoList) Delete() {
	deleteToDoFile(t.ToDos[t.CurrentToDo])
	t.ToDos = t.ToDos.Delete(t.CurrentToDo)
	t.selectUp()
}

func (t *ToDoList) selectUp() {
	t.CurrentToDo--
	if t.CurrentToDo < 0 {
		t.CurrentToDo = 0
	}
}

func (t *ToDoList) selectDown() {
	t.CurrentToDo++
	if t.CurrentToDo >= len(t.ToDos) {
		t.CurrentToDo = len(t.ToDos) - 1
	}
}

func (t *ToDoList) sort() {
	toDoRef := t.ToDos[t.CurrentToDo]
	sort.Sort(t.ToDos)
	for i, td := range t.ToDos {
		if td == toDoRef {
			t.CurrentToDo = i
			break
		}
	}
}
