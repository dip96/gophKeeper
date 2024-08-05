package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/login_data"
)

func ShowLoginPasswordDetails(app *tview.Application) {
	page := tview.NewFlex().SetDirection(tview.FlexRow)

	idMap := make(map[int]string)

	list := tview.NewList().
		ShowSecondaryText(false).
		SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
			id := idMap[index]
			ShowLoginPasswordForm(app, id)
		})

	resp, err := login_data.GetAllLoginData()
	if err != nil {
		ShowModal(app, "Ошибка при получении данных: "+err.Error())
		return
	}

	if len(resp.Items) == 0 {
		list.AddItem("Нет доступных записей", "", 0, nil)
	} else {
		for i, item := range resp.Items {
			list.AddItem(fmt.Sprintf("Login: %s, ID: %s", item.Login, item.Id), "", 0, nil)
			idMap[i] = item.Id
		}
	}

	createButton := tview.NewButton("Создать запись").SetSelectedFunc(func() {
		ShowLoginPasswordForm(app, "")
	})

	page.AddItem(list, 0, 1, true)
	page.AddItem(createButton, 3, 1, true)

	app.SetRoot(page, true).SetFocus(list)

	app.SetFocus(list)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			if app.GetFocus() == list {
				app.SetFocus(createButton)
			} else {
				app.SetFocus(list)
			}
			return nil
		}
		return event
	})
}

func ShowLoginPasswordForm(app *tview.Application, id string) {
	form := tview.NewForm()

	var initialLogin, initialPassword string

	if id != "" {
		resp, err := login_data.GetLoginData(id)
		if err != nil {
			ShowModal(app, "Ошибка при загрузке данных: "+err.Error())
			return
		}
		initialLogin = resp.Login
		initialPassword = resp.Password
	}

	form.AddInputField("Login", initialLogin, 20, nil, nil)
	form.AddPasswordField("Password", initialPassword, 20, '*', nil)

	form.AddButton("Сохранить", func() {
		login := form.GetFormItemByLabel("Login").(*tview.InputField).GetText()
		password := form.GetFormItemByLabel("Password").(*tview.InputField).GetText()

		var err error
		if id == "" {
			_, err = login_data.SaveLoginData(login, password, "")
		} else {
			_, err = login_data.EditLoginData(login, password, id)
		}

		if err != nil {
			ShowModal(app, "Ошибка при сохранении данных: "+err.Error())
		} else {
			ShowModal(app, "Запись успешно сохранена")
			ShowLoginPasswordDetails(app)
		}
	})

	form.AddButton("Отмена", func() {
		ShowLoginPasswordDetails(app)
	})

	form.SetBorder(true).SetTitle("Ввод логина и пароля").SetTitleAlign(tview.AlignCenter)

	// Обработка фокуса и нажатий клавиш
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			app.SetFocus(form.GetFormItem(1 % form.GetFormItemCount()))
			return nil
		}
		if event.Key() == tcell.KeyBacktab {
			app.SetFocus(form.GetFormItem((1 + form.GetFormItemCount()) % form.GetFormItemCount()))
			return nil
		}
		return event
	})

	app.SetFocus(form.GetFormItem(0))

	app.SetRoot(form, true)
}
