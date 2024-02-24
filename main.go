package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var error_icon fyne.Resource

func loadIcons() {
	var err error
	error_icon, err = fyne.LoadResourceFromPath("./icons/error.svg")
	if err != nil {
		log.Println("Failed to load icon:", err)
		// Consider loading a built-in fallback icon from fyne in case of error
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("dusky")

	loadIcons()

	modules := []Module{
		NewCalculatorModule(),
		NewWebSearchModule(), // Add the new modules here
	}

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Enter query...")

	var results []Result // Ensure this is declared correctly

	resultsList := widget.NewList(
		func() int { return len(results) },
		func() fyne.CanvasObject {
			icon := canvas.NewImageFromResource(nil) // Initially, no resource
			label := widget.NewLabel("")
			icon.SetMinSize(fyne.NewSize(24, 24)) // Ensure visibility
			hbox := container.NewHBox(icon, label)
			hbox.Layout = layout.NewHBoxLayout() // Ensure horizontal layout
			return hbox
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			hbox := co.(*fyne.Container)
			icon := hbox.Objects[0].(*canvas.Image)
			label := hbox.Objects[1].(*widget.Label)

			result := results[id]

			label.SetText(result.Title)
			if result.Icon != nil {
				icon.Resource = result.Icon
			} else {
				icon.Resource = error_icon
			}
			icon.Refresh() // Ensure the icon is refreshed after update
		},
	)

	searchEntry.OnChanged = func(query string) {
		results = []Result{} // Clear existing results for each new query
		for _, module := range modules {
			if module.CanHandle(query) {
				results = append(results, module.Handle(query)...)
			}
		}

		// Refresh the list to update with new results
		resultsList.Refresh()
	}

	content := container.NewBorder(searchEntry, nil, nil, nil, resultsList)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(500, 300))
	myWindow.ShowAndRun()
}
