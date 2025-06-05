package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/lcabrini/go-htmx-todo/internal/database"
)

type App struct {
	DB *database.Queries
}

//go:embed templates/*
var templates embed.FS
var t = template.Must(template.ParseFS(templates, "templates/*"))

func init() {
	godotenv.Load()
}

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		panic("DB_URL is not set")
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		panic("Unable to connect to the database")
	}
	defer conn.Close(ctx)

	queries := database.New(conn)
	app := App{
		DB: queries,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.Index)
	mux.HandleFunc("GET /tasks/", app.TaskListHandler)
	mux.HandleFunc("GET /tasks/new/", app.AddTaskForm)
	mux.HandleFunc("POST /tasks/add/", app.CreateTaskHandler)

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is not set")
	}

	addr := fmt.Sprintf(":%s", port)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	slog.Info("Starting server", "port", os.Getenv("PORT"))
	server.ListenAndServe()
}

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}
