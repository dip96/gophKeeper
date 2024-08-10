package login_data

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/service/login_data"
	"gophKeeper/internal/client/ui/modals"
	"gophKeeper/internal/client/ui/navigation"
)

func createLoginField(initialValue string) *tview.InputField {
	return tview.NewInputField().
		SetLabel("Login: ").
		SetText(initialValue)
}

func createPasswordField(initialValue string) *tview.InputField {
	return tview.NewInputField().
		SetLabel("Password: ").
		SetText(initialValue).
		SetMaskCharacter('*')
}

func ShowLoginPasswordForm(app *tview.Application, id string) tview.Primitive {
	var initialLogin, initialPassword string
	repository := login_data.LoginDataRepository{}

	if id != "" {
		resp, err := repository.GetData(id)
		if err != nil {
			modals.ShowModal(app, "Ошибка при загрузке данных: "+err.Error())
			return nil
		}
		initialLogin = resp.Login
		initialPassword = resp.Password
	}

	loginField := createLoginField(initialLogin)
	passwordField := createPasswordField(initialPassword)

	form := tview.NewForm().
		AddFormItem(loginField).
		AddFormItem(passwordField)

	form.AddButton("Сохранить", func() {
		login := loginField.GetText()
		password := passwordField.GetText()

		var err error
		dataRequest := login_data.LoginDataRequest{
			Login:    login,
			Password: password,
			Id:       id,
		}

		if id == "" {
			_, err = repository.SaveData(dataRequest)
		} else {
			_, err = repository.EditData(dataRequest)
		}

		if err != nil {
			modals.ShowModal(app, "Ошибка при сохранении данных: "+err.Error())
		} else {
			modals.ShowModal(app, "Запись успешно сохранена")
		}
	})

	form.AddButton("Отмена", func() {
		navigation.PopPage(app)
	})

	form.SetBorder(true).SetTitle("Ввод логина и пароля").SetTitleAlign(tview.AlignCenter)

	formItems := []tview.Primitive{
		loginField,
		passwordField,
		form.GetButton(0),
		form.GetButton(1),
	}

	currentFocus := 0

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			currentFocus = (currentFocus + 1) % len(formItems)
			app.SetFocus(formItems[currentFocus])
			return nil
		case tcell.KeyBacktab:
			currentFocus = (currentFocus - 1 + len(formItems)) % len(formItems)
			app.SetFocus(formItems[currentFocus])
			return nil
		}
		return event
	})

	app.SetFocus(loginField)

	return form
}
