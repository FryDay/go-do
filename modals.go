package main

import (
	"fmt"

	"github.com/rivo/tview"
)

// NewItemModal returns a new ItemModal primitive
func NewItemModal(t *ToDo) tview.Primitive {
	title := "New Item"
	name := ""
	if t != nil {
		title = "Edit Item"
		name = t.Name
	}

	form := tview.NewForm().
		SetButtonsAlign(tview.AlignCenter).
		AddInputField("Name", name, 20, nil, func(text string) { name = text })
	form.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignCenter)

	if t != nil {
		form.AddButton("Edit", func() {
			if name == "" {
				return
			}
			t.Edit(name)
			pages.RemovePage(itemPage)
			editMode = false
		})
	} else {
		form.AddButton("Add", func() {
			if name == "" {
				return
			}
			list.Add(NewToDo(name))
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
			AddItem(form, 7, 1, true).
			AddItem(nil, 0, 1, false), 30, 1, true).
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
