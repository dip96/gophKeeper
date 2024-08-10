package list_entities

import (
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/auth"
	"gophKeeper/internal/client/ui/modals"
	"gophKeeper/internal/client/ui/navigation"
	"gophKeeper/internal/client/ui/pages/entities/credit_card_data"
	"gophKeeper/internal/client/ui/pages/entities/login_data"
	"gophKeeper/internal/client/ui/pages/entities/text_data"
)

func CreateDataTypeMenu(app *tview.Application) tview.Primitive {
	list := tview.NewList().
		AddItem("Логин/пароль", "Управление логинами и паролями", 'l', func() {
			login_data.ShowLoginPasswordDetails(app)
		}).
		AddItem("Текстовые данные", "Управление произвольными текстовыми данными", 't', func() {
			text_data.ShowTextDataDetails(app)
		}).
		AddItem("Банковские карты", "Управление данными банковских карт", 'c', func() {
			credit_card_data.ShowCreditCardDataDetails(app)
		}).
		AddItem("Выйти", "Выйти из аккаунта", 'q', func() {
			authService, err := auth.GetAuthService()
			if err != nil {
				modals.ShowModal(app, "Ошибка инициализации сервиса: "+err.Error())
				return
			}

			err = authService.Logout()
			if err == nil {
				navigation.PopPage(app)
			} else {
				modals.ShowModal(app, "Ошибка при выходе: "+err.Error())
			}
		})

	return list
}
