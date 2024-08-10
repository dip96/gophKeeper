package index

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gophKeeper/internal/client/ui/forms/login"
	"gophKeeper/internal/client/ui/forms/registration"
	"gophKeeper/internal/client/ui/navigation"
)

func CreateMainMenu(app *tview.Application) tview.Primitive {
	registerButton := tview.NewButton("Регистрация").
		SetSelectedFunc(func() {
			navigation.PushPage(app, registration.CreateRegistrationForm(app), nil)
		})

	loginButton := tview.NewButton("Авторизация").
		SetSelectedFunc(func() {
			navigation.PushPage(app, login.CreateLoginForm(app), nil)
		})

	registerButton.SetLabelColor(tcell.ColorBlack).
		SetBackgroundColor(tcell.ColorWhite)
	loginButton.SetLabelColor(tcell.ColorBlack).
		SetBackgroundColor(tcell.ColorWhite)

	buttons := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(registerButton, 0, 1, true).
		AddItem(loginButton, 0, 1, false)

	buttons.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp, tcell.KeyDown:
			if app.GetFocus() == registerButton {
				loginButton.SetBackgroundColor(tcell.ColorBlue)
				registerButton.SetBackgroundColor(tcell.ColorBlack)
				app.SetFocus(loginButton)
			} else {
				registerButton.SetBackgroundColor(tcell.ColorRed)
				loginButton.SetBackgroundColor(tcell.ColorBlack)
				app.SetFocus(registerButton)
			}
			return nil
		}
		return event
	})

	frame := tview.NewFrame(buttons).
		SetBorders(0, 0, 0, 0, 1, 1)

	return frame
}
