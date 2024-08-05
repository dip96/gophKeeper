package ui

import (
	"github.com/rivo/tview"
)

func ShowModal(app *tview.Application, message string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			app.SetRoot(CreateMainMenu(app), true)
		})
	app.SetRoot(modal, true)
}
