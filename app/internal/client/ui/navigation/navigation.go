package navigation

import (
	"container/list"
	"github.com/rivo/tview"
)

type Page struct {
	Primitive   tview.Primitive
	OnFocusFunc func()
}

var (
	pageStack   = list.New()
	currentPage *Page
)

func PushPage(app *tview.Application, newPage tview.Primitive, onFocusFunc func()) {
	if currentPage != nil {
		pageStack.PushBack(currentPage)
	}
	currentPage = &Page{
		Primitive:   newPage,
		OnFocusFunc: onFocusFunc,
	}
	app.SetRoot(newPage, true)
}

func PopPage(app *tview.Application) {
	if pageStack.Len() > 0 {
		prevPage := pageStack.Back()
		pageStack.Remove(prevPage)
		currentPage = prevPage.Value.(*Page)
		app.SetRoot(currentPage.Primitive, true)

		if currentPage.OnFocusFunc != nil {
			currentPage.OnFocusFunc()
		}
	}
}
