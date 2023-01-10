package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var err error

	// Initalize the sql.DB connection pool and assign it to the models.DB
	// global variable.
	DB, err = sql.Open("postgres", "postgres://postgres:123456@localhost/todo_db")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/notes", notesIndex)
	http.ListenAndServe(":3000", nil)
}

// notesIndex sends a HTTP response listing all notes.
func notesIndex(w http.ResponseWriter, r *http.Request) {
	notes, err := AllNotes()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, note := range notes {
		fmt.Fprintf(w, "%d, %s, %s, %s, %s\n", note.Id, note.Title, note.Description, note.Date, note.Priority)
	}
}
