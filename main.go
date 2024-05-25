package main

import (
	"log"

	"github.com/jonesrussell/nope-five-o/services"
	"github.com/jonesrussell/nope-five-o/ui"
)

func main() {
	dbPath := "./nope.db"
	service, err := services.NewNoteService(dbPath)
	if err != nil {
		log.Fatalf("Failed to create note service: %v", err)
	}
	ui.StartUI(service)
}
