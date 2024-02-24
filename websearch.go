package main

import (
	"log"

	"fyne.io/fyne/v2"
)

// WebSearchModule searches the web
type WebSearchModule struct{}

// NewWebSearchModule creates a new instance of WebSearchModule
func NewWebSearchModule() *WebSearchModule {
	return &WebSearchModule{}
}

func (w *WebSearchModule) Icon() fyne.Resource {
	icon, err := fyne.LoadResourceFromPath("./icons/search.svg")
	if err != nil {
		log.Println("Failed to load icon:", err)
		return nil
	}
	return icon
}

func (w *WebSearchModule) CanHandle(query string) bool {
	// Assume web search module can handle any query
	return true
}

func (w *WebSearchModule) Handle(query string) []Result {
	// Placeholder for web search functionality
	// In a real implementation, you'd query a search API here
	searchResult := Result{
		Title: "Search the web for: " + query,
		Action: func() {
			// open web browser with search results
		},
		Icon: w.Icon(),
	}
	return []Result{searchResult}
}
