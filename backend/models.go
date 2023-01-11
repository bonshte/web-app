package main

import (
	"database/sql"
)

// Create an exported global variable to hold the database connection pool.
var DB *sql.DB

// AllBooks returns a slice of all books in the books table.
func AllNotes() ([]Note, error) {
	// Note that we are calling Query() on the global variable.
	rows, err := DB.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note

	for rows.Next() {
		var note Note

		err := rows.Scan(&note.UserId, &note.Title, &note.Description, &note.Date, &note.PriorityId)
		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func AllUsers() ([]User, error) {
	// Note that we are calling Query() on the global variable.
	rows, err := DB.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
