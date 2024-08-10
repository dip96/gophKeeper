package credit_card_data

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/service/credit_card_data"
	"gophKeeper/internal/client/ui/modals"
	"gophKeeper/internal/client/ui/navigation"
)

type creditCardFormData struct {
	CardholderName string
	CardNumber     string
	ExpirationDate string
	CVV            string
	Notes          string
}

func createInputField(label, initialValue string, fieldWidth int, isPassword bool) *tview.InputField {
	field := tview.NewInputField().
		SetLabel(label + ": ").
		SetText(initialValue).
		SetFieldWidth(fieldWidth)

	if isPassword {
		field.SetMaskCharacter('*')
	}

	return field
}

func createForm(data creditCardFormData) *tview.Form {
	return tview.NewForm().
		AddFormItem(createInputField("Cardholder Name", data.CardholderName, 40, false)).
		AddFormItem(createInputField("Card Number", data.CardNumber, 19, false)).
		AddFormItem(createInputField("Expiration Date (MM/YY)", data.ExpirationDate, 5, false)).
		AddFormItem(createInputField("CVV", data.CVV, 3, true)).
		AddFormItem(createInputField("Notes", data.Notes, 100, false))
}

func getFormData(form *tview.Form) creditCardFormData {
	return creditCardFormData{
		CardholderName: form.GetFormItemByLabel("Cardholder Name: ").(*tview.InputField).GetText(),
		CardNumber:     form.GetFormItemByLabel("Card Number: ").(*tview.InputField).GetText(),
		ExpirationDate: form.GetFormItemByLabel("Expiration Date (MM/YY): ").(*tview.InputField).GetText(),
		CVV:            form.GetFormItemByLabel("CVV: ").(*tview.InputField).GetText(),
		Notes:          form.GetFormItemByLabel("Notes: ").(*tview.InputField).GetText(),
	}
}

func handleSave(app *tview.Application, form *tview.Form, id string, repository *credit_card_data.CreditCardDataRepository) {
	formData := getFormData(form)

	dataRequest := credit_card_data.CreditCardDataRequest{
		CardholderName: formData.CardholderName,
		CardNumber:     formData.CardNumber,
		ExpirationDate: formData.ExpirationDate,
		CVV:            formData.CVV,
		Notes:          formData.Notes,
		Id:             id,
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

func setupFormNavigation(app *tview.Application, form *tview.Form) {
	formItems := []tview.Primitive{
		form.GetFormItemByLabel("Cardholder Name: "),
		form.GetFormItemByLabel("Card Number: "),
		form.GetFormItemByLabel("Expiration Date (MM/YY): "),
		form.GetFormItemByLabel("CVV: "),
		form.GetFormItemByLabel("Notes: "),
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
}

func ShowCreditCardDataForm(app *tview.Application, id string) tview.Primitive {
	repository := credit_card_data.CreditCardDataRepository{}
	var formData creditCardFormData

	if id != "" {
		resp, err := repository.GetData(id)
		if err != nil {
			modals.ShowModal(app, "Ошибка при загрузке данных: "+err.Error())
			return nil
		}
		formData = creditCardFormData{
			CardholderName: resp.CardholderName,
			CardNumber:     resp.CardNumber,
			ExpirationDate: resp.ExpirationDate,
			CVV:            resp.CVV,
			Notes:          resp.Notes,
		}
	}

	form := createForm(formData)

	form.AddButton("Сохранить", func() {
		handleSave(app, form, id, &repository)
	})

	form.AddButton("Отмена", func() {
		navigation.PopPage(app)
	})

	form.SetBorder(true).SetTitle("Ввод данных банковской карты").SetTitleAlign(tview.AlignCenter)

	setupFormNavigation(app, form)

	app.SetFocus(form.GetFormItem(0))

	return form
}
