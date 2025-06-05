package main

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/lcabrini/go-htmx-todo/internal/database"
)

func (app *App) TaskListHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := app.DB.ListTasks(context.Background())
	if err != nil {
		slog.Error("Error", "err", err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, "tasklist", tasks); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}

func (app *App) AddTaskForm(w http.ResponseWriter, r *http.Request) {
	if err := t.ExecuteTemplate(w, "newtaskform", nil); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}

func (app *App) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	priority := r.FormValue("priority")
	priority = strings.ToLower(priority)

	ctx := context.Background()
	task := database.CreateTaskParams{
		Title:       title,
		Description: description,
		Priority:    database.Priorities(priority),
	}
	app.DB.CreateTask(ctx, task)

	tasks, err := app.DB.ListTasks(context.Background())
	if err != nil {
		slog.Error("Error", "err", err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, "tasklist", tasks); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}
