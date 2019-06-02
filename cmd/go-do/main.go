package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	godo "github.com/FryDay/go-do"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	app       = tview.NewApplication()
	header    = tview.NewTextView()
	footer    = tview.NewTextView()
	flex      = tview.NewFlex()
	pages     = tview.NewPages()
	list      = NewToDoList()
	configDir = filepath.Join(os.Getenv("HOME"), ".config", "go-do")
	lockFile  = filepath.Join(configDir, ".lock")
	logFile   = filepath.Join(configDir, "go-do.log")
	errors    = make(chan error)

	editMode bool
)

const (
	mainPage    = "main"
	itemPage    = "item"
	confirmPage = "confirm"
)

func init() {
	if _, err := os.Stat(lockFile); err == nil {
		fmt.Println("another instance is running")
		os.Exit(1)
	}

	go func() {
		l, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			panic(err)
		}
		defer l.Close()

		log.SetOutput(l)
		for {
			err := <-errors
			if err != nil {
				log.Println(err.Error())
			}
		}
	}()

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
	err := os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(lockFile, nil, 0600)

	var files []string
	err = filepath.Walk(configDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		todo := &godo.ToDo{}
		err = json.Unmarshal(b, todo)
		if err != nil {
			panic(err)
		}
		list.Add(todo)
	}
}

func main() {
	defer os.Remove(lockFile)
	defer app.Stop()

	app.SetRoot(pages, true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
