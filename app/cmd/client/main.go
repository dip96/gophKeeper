package main

import (
	"github.com/rivo/tview"
	"gophKeeper/internal/client/ui/navigation"
	"gophKeeper/internal/client/ui/pages/index"
)

func main() {
	app := tview.NewApplication()
	mainMenu := index.CreateMainMenu(app)
	navigation.PushPage(app, mainMenu, nil)
	if err := app.SetRoot(mainMenu, true).Run(); err != nil {
		panic(err)
	}
}
