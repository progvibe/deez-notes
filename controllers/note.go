package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/progvibe/deez-notes/repositories"
)

type NoteHandler struct {
	DB *sql.DB
}

type User struct {
	ID   int
	Name string
}

type Note struct {
	ID      int
	Title   string
	Content string
}

func (h NoteHandler) HandleGetNotes(c echo.Context) error {
	noteRepository := repositories.NoteRepository{DB: h.DB}
	notes, err := noteRepository.AllNotesForUser(1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to retrieve notes for user: %v\n", err)
		os.Exit(1)
	}
	return c.JSON(http.StatusFound, notes)
}

func (h NoteHandler) HandleNoteSave(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	fmt.Fprintf(os.Stdout, "title: %v\n", title)
	fmt.Fprintf(os.Stdout, "content: %v\n", content)
	id := 0
	err := h.DB.QueryRow("insert into notes (title, content, user_id) values (?, ?, ?) returning id;", title, content, 1).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	return c.JSON(http.StatusCreated, id)
}

func (h NoteHandler) HandleNoteGet(c echo.Context) error {
	id := c.Param("id")
	rows, err := h.DB.Query("select * from notes where id=?", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content); err != nil {
			fmt.Println("Error scanning row:", err)
			return err
		}
		return c.JSON(http.StatusFound, note)
	}
	return c.NoContent(http.StatusNotFound)
}
