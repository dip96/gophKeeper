package login_data

import (
	"fmt"
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/service/login_data"
	uiLoginData "gophKeeper/internal/client/ui/forms/login_data"
	"gophKeeper/internal/client/ui/pages/entities"
)

func ShowLoginPasswordDetails(app *tview.Application) {
	repository := &login_data.LoginDataRepository{}

	entities.ShowDataDetails(
		app,
		repository,
		uiLoginData.ShowLoginPasswordForm,
		func(item login_data.LoginDataRequest) string {
			return fmt.Sprintf("Login: %s, ID: %s", item.Login, item.Id)
		},
	)
}
