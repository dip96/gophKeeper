package text_data

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/service/text_data"
	"gophKeeper/internal/client/ui/modals"
	"gophKeeper/internal/client/ui/navigation"
)

//TODO вынести похожую логику в формах в другой файл

func createTextField(initialText string) *tview.InputField {
	return tview.NewInputField().
		SetLabel("Text: ").
		SetText(initialText)
}

func handleSave(app *tview.Application, textField *tview.InputField, id string, repository *text_data.TextDataRepository) {
	text := textField.GetText()

	dataRequest := text_data.TextDataRequest{
		Text: text,
		Id:   id,
	}

	var err error
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
}

func setupFormNavigation(app *tview.Application, form *tview.Form, formItems []tview.Primitive) {
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
}

func ShowTextDataForm(app *tview.Application, id string) tview.Primitive {
	repository := text_data.TextDataRepository{}
	var initialText string

	if id != "" {
		resp, err := repository.GetData(id)
		if err != nil {
			modals.ShowModal(app, "Ошибка при загрузке данных: "+err.Error())
			return nil
		}
		initialText = resp.Text
	}

	textField := createTextField(initialText)

	form := tview.NewForm().
		AddFormItem(textField).
		AddButton("Сохранить", func() {
			handleSave(app, textField, id, &repository)
		}).
		AddButton("Отмена", func() {
			navigation.PopPage(app)
		})

	form.SetBorder(true).SetTitle("Ввод текста").SetTitleAlign(tview.AlignCenter)

	formItems := []tview.Primitive{
		textField,
		form.GetButton(0),
		form.GetButton(1),
	}

	setupFormNavigation(app, form, formItems)

	app.SetFocus(textField)

	return form
}
