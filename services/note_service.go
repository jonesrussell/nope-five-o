// services/note_service.go

package services

import (
	"database/sql"
	"log"

	"github.com/jonesrussell/nope-five-o/models"
	_ "github.com/xeodou/go-sqlcipher"
)

// NoteService manages note entities.
type NoteService struct {
	db *sql.DB
}

// NewNoteService initializes a new NoteService with an encrypted SQLite database.
func NewNoteService(dbPath string) (*NoteService, error) {
	db, err := sql.Open("sqlite3", dbPath+"?_key=yourpassword")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS notes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    body TEXT NOT NULL
);`)
	if err != nil {
		log.Fatal(err)
	}

	return &NoteService{db}, nil
}

// AddNote adds a new note to the database.
func (s *NoteService) AddNote(title, body string) error {
	_, err := s.db.Exec("INSERT INTO notes (title, body) VALUES (?,?)", title, body)
	return err
}

// GetAllNotes retrieves all notes from the database.
func (s *NoteService) GetAllNotes() ([]*models.Note, error) {
	rows, err := s.db.Query("SELECT id, title, body FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*models.Note
	for rows.Next() {
		note := &models.Note{}
		err := rows.Scan(&note.ID, &note.Title, &note.Body)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

// DeleteNote deletes a note by its ID.
func (s *NoteService) DeleteNote(id int64) error {
	_, err := s.db.Exec("DELETE FROM notes WHERE id =?", id)
	return err
}
