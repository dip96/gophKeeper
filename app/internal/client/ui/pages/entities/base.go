package entities

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/service"
	"gophKeeper/internal/client/ui/modals"
	"gophKeeper/internal/client/ui/navigation"
)

func ShowDataDetails[T service.DataItem, R service.DataRepository[T]](
	app *tview.Application, repository R, formCreator func(app *tview.Application, id string) tview.Primitive,
	getItemDetails func(item T) string) {

	idMap := make(map[int]string)

	list := tview.NewList().
		ShowSecondaryText(false).
		SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
			id := idMap[index]
			navigation.PushPage(app, formCreator(app, id), nil)
		})

	loadData := func() {
		list.Clear()
		idMap = make(map[int]string)
		items, err := repository.GetAllData()
		if err != nil {
			modals.ShowModal(app, "Ошибка при получении данных: "+err.Error())
			return
		}

		if len(items) == 0 {
			list.AddItem("Нет доступных записей", "", 0, nil)
		} else {
			for i, item := range items {
				list.AddItem(getItemDetails(item), "", 0, nil)
				idMap[i] = item.GetID()
			}
		}
	}

	page := tview.NewFlex().SetDirection(tview.FlexRow)

	createButton := tview.NewButton("Создать запись").SetSelectedFunc(func() {
		navigation.PushPage(app, formCreator(app, ""), nil)
	})

	backButton := tview.NewButton("Назад").SetSelectedFunc(func() {
		navigation.PopPage(app)
	})

	buttonsRow := tview.NewFlex().
		AddItem(createButton, 0, 1, false).
		AddItem(backButton, 0, 1, false)

	page.AddItem(list, 0, 1, true)
	page.AddItem(buttonsRow, 0, 1, false)

	page.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			switch app.GetFocus() {
			case list:
				app.SetFocus(createButton)
			case createButton:
				app.SetFocus(backButton)
			case backButton:
				app.SetFocus(list)
			}
			return nil
		case tcell.KeyBacktab:
			switch app.GetFocus() {
			case list:
				app.SetFocus(backButton)
			case createButton:
				app.SetFocus(list)
			case backButton:
				app.SetFocus(createButton)
			}
			return nil
		}
		return event
	})

	loadData()
	navigation.PushPage(app, page, loadData)
}
