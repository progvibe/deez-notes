package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/progvibe/deez-notes/controllers"
	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL := os.Getenv("DB_URL")
	// dbToken := os.Getenv("DB_TOKEN")
	// url := dbURL + "?authToken=" + dbToken
	url := dbURL
	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	noteHandler := controllers.NoteHandler{DB: db}
	e.GET("/notes", noteHandler.HandleGetNotes)
	e.GET("/note/:id", noteHandler.HandleNoteGet)
	e.POST("/note", noteHandler.HandleNoteSave)

	e.Logger.Fatal(e.Start(":1323"))
}
