package repositories

import (
	"database/sql"
	"fmt"
	"os"
)

type NoteRepository struct {
	DB *sql.DB
}

type Note struct {
	ID      int
	UserID  int
	Title   string
	Content string
}

func (r NoteRepository) AllNotesForUser(userID int) ([]Note, error) {
	rows, err := r.DB.Query("select * from notes where user_id = ?;", userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var notes []Note
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Content); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}
