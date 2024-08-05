package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/auth"
	"gophKeeper/internal/models/entities"
)

func CreateMainMenu(app *tview.Application) tview.Primitive {
	registerButton := tview.NewButton("Регистрация").
		SetSelectedFunc(func() {
			app.SetRoot(CreateRegistrationForm(app), true)
		})

	loginButton := tview.NewButton("Авторизация").
		SetSelectedFunc(func() {
			app.SetRoot(CreateLoginForm(app), true)
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

func CreateDataTypeMenu(app *tview.Application) tview.Primitive {
	list := tview.NewList().
		AddItem("Логин/пароль", "Управление логинами и паролями", 'l', func() {
			ShowLoginPasswordDetails(app)
		}).
		AddItem("Текстовые данные", "Управление произвольными текстовыми данными", 't', func() {
			ShowDataTypeDetails(app, entities.DataTypeTextData)
		}).
		AddItem("Бинарные данные", "Управление произвольными бинарными данными", 'b', func() {
			ShowDataTypeDetails(app, entities.DataTypeBinaryData)
		}).
		AddItem("Банковские карты", "Управление данными банковских карт", 'c', func() {
			ShowDataTypeDetails(app, entities.DataTypeBankCardData)
		}).
		AddItem("Выйти", "Выйти из аккаунта", 'q', func() {
			err := auth.Logout()
			if err == nil {
				app.SetRoot(CreateMainMenu(app), true)
			} else {
				ShowModal(app, "Ошибка при выходе: "+err.Error())
			}
		})

	return list
}

func ShowDataTypeDetails(app *tview.Application, dataType int) {
	var message string
	switch dataType {
	case entities.DataTypeLoginPassword:
		message = "Управление логинами и паролями"
	case entities.DataTypeTextData:
		message = "Управление текстовыми данными"
	case entities.DataTypeBinaryData:
		message = "Управление бинарными данными"
	case entities.DataTypeBankCardData:
		message = "Управление данными банковских карт"
	}
	ShowModal(app, message)
}
