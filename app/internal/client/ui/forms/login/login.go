package login

import (
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/auth"
	"gophKeeper/internal/client/ui/modals"
	"gophKeeper/internal/client/ui/navigation"
	"gophKeeper/internal/client/ui/pages/list_entities"
)

func createLoginField() *tview.InputField {
	return tview.NewInputField().SetLabel("Логин: ")
}

func createPasswordField() *tview.InputField {
	return tview.NewInputField().SetLabel("Пароль: ").SetMaskCharacter('*')
}

func CreateLoginForm(app *tview.Application) tview.Primitive {
	loginField := createLoginField()
	passwordField := createPasswordField()

	form := tview.NewForm().
		AddFormItem(loginField).
		AddFormItem(passwordField).
		AddButton("Войти", func() {
			handleLogin(app, loginField, passwordField)
		}).
		AddButton("Отмена", func() {
			navigation.PopPage(app)
		})

	return form
}

func handleLogin(app *tview.Application, loginField, passwordField *tview.InputField) {
	login := loginField.GetText()
	password := passwordField.GetText()

	authService, err := auth.GetAuthService()
	if err != nil {
		modals.ShowModal(app, "Ошибка инициализации сервиса: "+err.Error())
		return
	}

	if _, err := authService.LoginUser(login, password); err != nil {
		modals.ShowModal(app, "Ошибка авторизации: "+err.Error())
	} else {
		navigation.PushPage(app, list_entities.CreateDataTypeMenu(app), nil)
	}
}
