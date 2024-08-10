package text_data

import (
	"fmt"
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/service/text_data"
	uiTextData "gophKeeper/internal/client/ui/forms/text_data"
	"gophKeeper/internal/client/ui/pages/entities"
)

func ShowTextDataDetails(app *tview.Application) {
	repository := &text_data.TextDataRepository{}

	entities.ShowDataDetails(
		app,
		repository,
		uiTextData.ShowTextDataForm,
		func(item text_data.TextDataRequest) string {
			return fmt.Sprintf("Text: %s, ID: %s", item.Text, item.Id)
		},
	)
}
