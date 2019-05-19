package main

import (
	"fmt"
	"sort"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// ToDoList ...
type ToDoList struct {
	*tview.Box
	ToDos       ToDos
	CurrentToDo int
}

// NewToDoList returns a new todo list.
func NewToDoList() *ToDoList {
	return &ToDoList{
		Box: tview.NewBox(),
	}
}

// InputHandler ...
func (t *ToDoList) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return t.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
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
		case ' ':
			if t.ToDos[t.CurrentToDo].TimeCompleted.IsZero() {
				t.ToDos[t.CurrentToDo].Complete()
			} else {
				t.ToDos[t.CurrentToDo].Reopen()
			}
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
		line := fmt.Sprintf(` %s[white]   %s`, checkbox, td.Name)
		if i == t.CurrentToDo {
			line = fmt.Sprintf(` %s[white]   [::r]%s`, checkbox, td.Name)
		}
		tview.Print(screen, line, x, y+i, width, tview.AlignLeft, tcell.ColorDefault)
	}
}

// Add a new item
func (t *ToDoList) Add(item *ToDo) {
	t.ToDos = t.ToDos.Add(item)
	t.sort()
}

// Delete the currently selected item
func (t *ToDoList) Delete() {
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
