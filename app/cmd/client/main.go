package main

import (
	"context"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"google.golang.org/grpc"
	auth "gophKeeper/internal/lib/auth/keyring"
	pb "gophKeeper/protobuf/V1/users"
)

func main() {
	app := tview.NewApplication()
	mainMenu := createMainMenu(app)
	if err := app.SetRoot(mainMenu, true).Run(); err != nil {
		panic(err)
	}
}

func createMainMenu(app *tview.Application) tview.Primitive {
	registerButton := tview.NewButton("Регистрация").
		SetSelectedFunc(func() {
			app.SetRoot(createRegistrationForm(app), true)
		})

	loginButton := tview.NewButton("Авторизация").
		SetSelectedFunc(func() {
			app.SetRoot(createLoginForm(app), true)
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

func createRegistrationForm(app *tview.Application) tview.Primitive {
	loginField := tview.NewInputField().SetLabel("Логин: ")
	passwordField := tview.NewInputField().SetLabel("Пароль: ").SetMaskCharacter('*')

	form := tview.NewForm().
		AddFormItem(loginField).
		AddFormItem(passwordField).
		AddButton("Регистрация", func() {
			login := loginField.GetText()
			password := passwordField.GetText()
			platform := pb.Platform_UNKNOWN
			if err := registerUser(login, password, platform); err != nil {
				showModal(app, "Ошибка регистрации: "+err.Error())
			} else {
				showModal(app, "Регистрация прошла успешно!")
			}
		}).
		AddButton("Отмена", func() {
			app.SetRoot(createMainMenu(app), true)
		})

	return form
}

func createLoginForm(app *tview.Application) tview.Primitive {
	loginField := tview.NewInputField().SetLabel("Логин: ")
	passwordField := tview.NewInputField().SetLabel("Пароль: ").SetMaskCharacter('*')

	form := tview.NewForm().
		AddFormItem(loginField).
		AddFormItem(passwordField).
		AddButton("Войти", func() {
			login := loginField.GetText()
			password := passwordField.GetText()
			if _, err := loginUser(login, password); err != nil {
				showModal(app, "Ошибка авторизации: "+err.Error())
			} else {
				showModal(app, "Авторизация успешна!")
				app.SetRoot(createMainMenu(app), true)
			}
		}).
		AddButton("Отмена", func() {
			app.SetRoot(createMainMenu(app), true)
		})

	return form
}

func registerUser(login, password string, platform pb.Platform) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	req := &pb.UserRegistrationRequest{
		Login:    login,
		Password: password,
		Platform: platform,
	}

	_, err = client.Registration(context.Background(), req)
	return err
}

func loginUser(login, password string) (string, error) {
	conn, err := getConn()
	if err != nil {
		return "", err
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	req := &pb.LoginRequest{
		Login:    login,
		Password: password,
	}

	resp, err := client.Login(context.Background(), req)
	if err != nil {
		return "", err
	}

	// Сохраняем токен
	if err := auth.SaveToken(resp.Token); err != nil {
		return "", err
	}

	return resp.Token, nil
}

func getConn() (*grpc.ClientConn, error) {
	//return grpc.Dial("172.20.255.21:3201", grpc.WithInsecure())
	return grpc.Dial("0.0.0.0:3201", grpc.WithInsecure())
}

func showModal(app *tview.Application, message string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			app.SetRoot(createMainMenu(app), true)
		})
	app.SetRoot(modal, true)
}

func logout(app *tview.Application) {
	if err := auth.DeleteToken(); err != nil {
		showModal(app, "Ошибка при выходе: "+err.Error())
	} else {
		showModal(app, "Выход выполнен успешно")
		app.SetRoot(createMainMenu(app), true)
	}
}
