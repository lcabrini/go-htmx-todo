package main

import (
	"context"
	"log/slog"
	"net/http"
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
