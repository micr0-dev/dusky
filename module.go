package main

import "fyne.io/fyne/v2"

// Result represents a single result item with a title and an action handler
type Result struct {
	Title  string
	Action func()
	Icon   fyne.Resource
}

// Module represents the interface that all modules must implement
type Module interface {
	// CanHandle checks if this module can handle the given query
	CanHandle(query string) bool

	// Handle executes the module's functionality and returns the results
	Handle(query string) []Result
}
