package ui

import (
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/auth"
	pb "gophKeeper/protobuf/V1/users"
)

func CreateRegistrationForm(app *tview.Application) tview.Primitive {
	loginField := tview.NewInputField().SetLabel("Логин: ")
	passwordField := tview.NewInputField().SetLabel("Пароль: ").SetMaskCharacter('*')

	form := tview.NewForm().
		AddFormItem(loginField).
		AddFormItem(passwordField).
		AddButton("Регистрация", func() {
			login := loginField.GetText()
			password := passwordField.GetText()
			platform := pb.Platform_UNKNOWN
			if err := auth.RegisterUser(login, password, platform); err != nil {
				ShowModal(app, "Ошибка регистрации: "+err.Error())
			} else {
				ShowModal(app, "Регистрация прошла успешно!")
			}
		}).
		AddButton("Отмена", func() {
			app.SetRoot(CreateMainMenu(app), true)
		})

	return form
}

func CreateLoginForm(app *tview.Application) tview.Primitive {
	loginField := tview.NewInputField().SetLabel("Логин: ")
	passwordField := tview.NewInputField().SetLabel("Пароль: ").SetMaskCharacter('*')

	form := tview.NewForm().
		AddFormItem(loginField).
		AddFormItem(passwordField).
		AddButton("Войти", func() {
			login := loginField.GetText()
			password := passwordField.GetText()
			if _, err := auth.LoginUser(login, password); err != nil {
				ShowModal(app, "Ошибка авторизации: "+err.Error())
			} else {
				app.SetRoot(CreateDataTypeMenu(app), true)
			}
		}).
		AddButton("Отмена", func() {
			app.SetRoot(CreateMainMenu(app), true)
		})

	return form
}
