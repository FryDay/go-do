package main

import (
	"fmt"

	godo "github.com/FryDay/go-do"
	"github.com/gdamore/tcell/v2"
	"gitlab.com/tslocum/cview"
)

// NewItemModal returns a new ItemModal primitive
func NewItemModal(td *godo.ToDo) cview.Primitive {
	title := "New Item"
	var name, note string
	if td != nil {
		title = "Edit Item"
		name = td.Name
		note = td.Note
	}

	form := cview.NewForm()
	form.SetButtonsAlign(cview.AlignCenter)
	form.AddInputField("Name", name, 51, nil, func(text string) { name = text })
	form.AddInputField("Note", note, 51, nil, func(text string) { note = text })
	form.SetBorder(true)
	form.SetTitle(title)
	form.SetTitleAlign(cview.AlignCenter)
	form.SetTitleColor(tcell.ColorDefault)
	form.SetBorderColor(tcell.ColorBlue)
	form.SetFieldBackgroundColor(tcell.ColorBlue)
	form.SetFieldBackgroundColorFocused(tcell.ColorBlack)
	form.SetLabelColor(tcell.ColorBlue)
	form.SetButtonBackgroundColor(tcell.ColorBlue)
	form.SetButtonBackgroundColorFocused(tcell.ColorBlack)
	form.SetButtonTextColor(tcell.ColorBlack)
	form.SetButtonTextColorFocused(tcell.ColorBlue)

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

	newFlex := cview.NewFlex()
	newFlex.SetDirection(cview.FlexRow)
	newFlex.AddItem(nil, 0, 1, false)
	newFlex.AddItem(form, 9, 1, true)
	newFlex.AddItem(nil, 0, 1, false)

	newModal := cview.NewFlex()
	newModal.AddItem(nil, 0, 1, false)
	newModal.AddItem(newFlex, 60, 1, true)
	newModal.AddItem(nil, 0, 1, false)

	return newModal
}

//NewConfirmationModal returns a new ConfirmationModal primitive
func NewConfirmationModal(action string) cview.Primitive {
	title := fmt.Sprintf("Confirm %s", action)

	form := cview.NewForm()
	form.SetButtonsAlign(cview.AlignCenter)
	form.SetBorder(true)
	form.SetTitle(title)
	form.SetTitleAlign(cview.AlignCenter)
	form.SetTitle(title)
	form.SetTitleAlign(cview.AlignCenter)
	form.SetTitleColor(tcell.ColorDefault)
	form.SetBorderColor(tcell.ColorBlue)
	form.SetFieldBackgroundColor(tcell.ColorBlue)
	form.SetFieldBackgroundColorFocused(tcell.ColorBlack)
	form.SetLabelColor(tcell.ColorBlue)
	form.SetButtonBackgroundColor(tcell.ColorBlue)
	form.SetButtonBackgroundColorFocused(tcell.ColorBlack)
	form.SetButtonTextColor(tcell.ColorBlack)
	form.SetButtonTextColorFocused(tcell.ColorBlue)

	form.AddButton("Confirm", func() {
		list.Delete()
		pages.RemovePage(confirmPage)
	})
	form.AddButton("Cancel", func() {
		pages.RemovePage(confirmPage)
	})

	newFlex := cview.NewFlex()
	newFlex.SetDirection(cview.FlexRow)
	newFlex.AddItem(nil, 0, 1, false)
	newFlex.AddItem(form, 5, 1, true)
	newFlex.AddItem(nil, 0, 1, false)

	newConfirmation := cview.NewFlex()
	newConfirmation.AddItem(nil, 0, 1, false)
	newConfirmation.AddItem(newFlex, 30, 1, true)
	newConfirmation.AddItem(nil, 0, 1, false)

	return newConfirmation
}
