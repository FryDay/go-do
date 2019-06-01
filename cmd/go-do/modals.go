package main

import (
	"fmt"

	godo "github.com/FryDay/go-do"
	"github.com/rivo/tview"
)

// NewItemModal returns a new ItemModal primitive
func NewItemModal(td *godo.ToDo) tview.Primitive {
	title := "New Item"
	var name, note string
	if td != nil {
		title = "Edit Item"
		name = td.Name
		note = td.Note
	}

	form := tview.NewForm().
		SetButtonsAlign(tview.AlignCenter).
		AddInputField("Name", name, 51, nil, func(text string) { name = text }).
		AddInputField("Note", note, 51, nil, func(text string) { note = text })
	form.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignCenter)

	if td != nil {
		form.AddButton("Edit", func() {
			if name == "" {
				return
			}
			td.Edit(name, note)
			writeToDoFile(td)
			pages.RemovePage(itemPage)
			editMode = false
		})
	} else {
		form.AddButton("Add", func() {
			if name == "" {
				return
			}
			newToDo := godo.NewToDo(name, note)
			writeToDoFile(newToDo)
			list.Add(newToDo)
			pages.RemovePage(itemPage)
			editMode = false
		})
	}

	form.AddButton("Cancel", func() {
		pages.RemovePage(itemPage)
		editMode = false
	})

	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(form, 9, 1, true).
			AddItem(nil, 0, 1, false), 60, 1, true).
		AddItem(nil, 0, 1, false)
}

//NewConfirmationModal returns a new ConfirmationModal primitive
func NewConfirmationModal(action string) tview.Primitive {
	title := fmt.Sprintf("Confirm %s", action)

	form := tview.NewForm().
		SetButtonsAlign(tview.AlignCenter)
	form.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignCenter)

	form.AddButton("Confirm", func() {
		list.Delete()
		pages.RemovePage(confirmPage)
	}).
		AddButton("Cancel", func() {
			pages.RemovePage(confirmPage)
		})

	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(form, 5, 1, true).
			AddItem(nil, 0, 1, false), 30, 1, true).
		AddItem(nil, 0, 1, false)
}
