package credit_card_data

import (
	"fmt"
	"github.com/rivo/tview"
	"gophKeeper/internal/client/grpc/service/credit_card_data"
	uiCreditCardData "gophKeeper/internal/client/ui/forms/credit_card_data"
	"gophKeeper/internal/client/ui/pages/entities"
)

func ShowCreditCardDataDetails(app *tview.Application) {
	repository := &credit_card_data.CreditCardDataRepository{}

	entities.ShowDataDetails(
		app,
		repository,
		uiCreditCardData.ShowCreditCardDataForm,
		func(item credit_card_data.CreditCardDataRequest) string {
			return fmt.Sprintf("Cardholder: %s, Card Number: %s, Expiration: %s, ID: %s",
				item.CardholderName, maskCardNumber(item.CardNumber), item.ExpirationDate, item.Id)
		},
	)
}

func maskCardNumber(cardNumber string) string {
	if len(cardNumber) < 4 {
		return cardNumber
	}
	return fmt.Sprintf("**** **** **** %s", cardNumber[len(cardNumber)-4:])
}
