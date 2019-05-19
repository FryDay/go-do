package main

import (
	"os"
	"path/filepath"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	app    = tview.NewApplication()
	header = tview.NewTextView()
	footer = tview.NewTextView()
	flex   = tview.NewFlex()
	pages  = tview.NewPages()
	list   = NewToDoList()

	editMode bool
)

const (
	mainPage    = "main"
	itemPage    = "item"
	confirmPage = "confirm"
)

func init() {
	app.SetBeforeDrawFunc(func(s tcell.Screen) bool {
		s.Clear()
		return false
	})

	app.SetInputCapture(input)

	header.SetScrollable(false)
	header.SetDynamicColors(true)
	header.SetText("[::b]Go-Do")

	list.SetBackgroundColor(tcell.ColorDefault)

	footer.SetScrollable(false)
	footer.SetDynamicColors(true)
	footer.SetText("[::b][a[]Add Item  [e[]Edit Item  [d[]Delete  [q[]Quit")

	flex.SetDirection(tview.FlexRow)
	flex.AddItem(header, 1, 1, false)
	flex.AddItem(list, 0, 1, true)
	flex.AddItem(footer, 1, 1, false)

	pages.SetBackgroundColor(tcell.ColorDefault)
	pages.AddAndSwitchToPage(mainPage, flex, true)

	load()
}

func input(event *tcell.EventKey) *tcell.EventKey {
	if editMode {
		return event
	}
	switch event.Rune() {
	case 'a':
		modal := NewItemModal(nil)
		editMode = true
		pages.AddPage(itemPage, modal, true, true)
	case 'e':
		modal := NewItemModal(list.ToDos[list.CurrentToDo])
		editMode = true
		pages.AddPage(itemPage, modal, true, true)
	case 'd':
		if len(list.ToDos) > 0 {
			modal := NewConfirmationModal("Delete")
			pages.AddPage(confirmPage, modal, true, true)
		}
	case 'q':
		app.Stop()
	}
	return event
}

func load() {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", "go-do")
	err := os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	//TODO: Load Files
}

func main() {
	defer app.Stop()

	app.SetRoot(pages, true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
