package registration

import (
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/auth"
	"gophKeeper/internal/client/ui/modals"
	"gophKeeper/internal/client/ui/navigation"
	pb "gophKeeper/protobuf/V1/users"
)

func createLoginField() *tview.InputField {
	return tview.NewInputField().
		SetLabel("Логин: ")
}

func createPasswordField() *tview.InputField {
	return tview.NewInputField().
		SetLabel("Пароль: ").
		SetMaskCharacter('*')
}

func handleRegistration(app *tview.Application, loginField, passwordField *tview.InputField) {
	login := loginField.GetText()
	password := passwordField.GetText()
	platform := pb.Platform_UNKNOWN

	authService, err := auth.GetAuthService()
	if err != nil {
		modals.ShowModal(app, "Ошибка инициализации сервиса: "+err.Error())
		return
	}

	if err := authService.RegisterUser(login, password, platform); err != nil {
		modals.ShowModal(app, "Ошибка регистрации: "+err.Error())
	} else {
		modals.ShowModal(app, "Регистрация прошла успешно!")
	}
}

func CreateRegistrationForm(app *tview.Application) tview.Primitive {
	loginField := createLoginField()
	passwordField := createPasswordField()

	form := tview.NewForm().
		AddFormItem(loginField).
		AddFormItem(passwordField).
		AddButton("Регистрация", func() {
			handleRegistration(app, loginField, passwordField)
		}).
		AddButton("Отмена", func() {
			navigation.PopPage(app)
		})

	return form
}
