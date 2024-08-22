package modals

import (
	"github.com/rivo/tview"
	"gophKeeper/internal/client/ui/navigation"
)

func ShowModal(app *tview.Application, message string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			navigation.PopPage(app)
		})
	navigation.PushPage(app, modal, nil)
}
