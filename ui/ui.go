package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jonesrussell/nope-five-o/services"
)

func StartUI(service *services.NoteService) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n1. Add Note")
		fmt.Println("2. View Notes")
		fmt.Println("3. Delete Note")
		fmt.Println("4. Exit")
		fmt.Print("Enter choice: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		switch text {
		case "1":
			AddNote(service)
		case "2":
			ViewNotes(service)
		case "3":
			DeleteNote(service)
		case "4":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

func AddNote(service *services.NoteService) {
	fmt.Print("Enter note title: ")
	title, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter note body: ")
	body, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	body = strings.TrimSpace(body)

	err := service.AddNote(title, body)
	if err != nil {
		fmt.Printf("Error adding note: %v\n", err)
	} else {
		fmt.Println("Note added successfully.")
	}
}

func ViewNotes(service *services.NoteService) {
	notes, err := service.GetAllNotes()
	if err != nil {
		fmt.Printf("Error retrieving notes: %v\n", err)
		return
	}

	for _, note := range notes {
		fmt.Printf("Title: %s\nBody: %s\n\n", note.Title, note.Body)
	}
}

func DeleteNote(service *services.NoteService) {
	fmt.Print("Enter note ID to delete: ")
	idStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	// Convert the ID from string to int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Invalid ID format: %v\n", err)
		return
	}

	err = service.DeleteNote(id)
	if err != nil {
		fmt.Printf("Error deleting note: %v\n", err)
	} else {
		fmt.Println("Note deleted successfully.")
	}
}
