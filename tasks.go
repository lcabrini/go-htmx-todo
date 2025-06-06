package main

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/google/uuid"
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
		slog.Error("Error", "error", err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
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

	app.TaskListHandler(w, r)
}

func (app *App) UpdateTaskForm(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id, _ := uuid.Parse(r.PathValue("id"))
	task, err := app.DB.GetTask(ctx, id)
	if err != nil {
		slog.Error("Error", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, "updatetaskform", task); err != nil {
		slog.Error("Error", "error", err.Error())
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}

func (app *App) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.Parse(r.PathValue("id"))
	title := r.FormValue("title")
	description := r.FormValue("description")
	priority := strings.ToLower(r.FormValue("priority"))

	var completed bool
	switch strings.ToLower(r.FormValue("completed")) {
	case "yes", "on":
		completed = true
	case "no", "off":
		completed = false
	default:
		completed = false
	}

	ctx := context.Background()
	task := database.UpdateTaskParams{
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    database.Priorities(priority),
		Completed:   completed,
	}

	_, err := app.DB.UpdateTask(ctx, task)
	if err != nil {
		slog.Error("Error", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	app.TaskListHandler(w, r)
}

func (app *App) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

}
