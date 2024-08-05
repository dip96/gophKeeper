package main

import (
	"github.com/rivo/tview"
	"gophKeeper/internal/client/ui"
)

func main() {
	app := tview.NewApplication()
	mainMenu := ui.CreateMainMenu(app)
	if err := app.SetRoot(mainMenu, true).Run(); err != nil {
		panic(err)
	}
}
