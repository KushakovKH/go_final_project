package main

import (
	"log"
	"net/http"

	"go_final_project/database"
	"go_final_project/handlers"

	"github.com/go-chi/chi/v5"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := database.InitializeDB("scheduler.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	repository := database.NewRepository(db)

	startServer(repository)
}

func startServer(repository *database.Repository) {
	cNR := chi.NewRouter()

	directoryPath := "web"
	fileServer := http.FileServer(http.Dir(directoryPath))
	cNR.Handle("/*", fileServer)

	cNR.Get("/api/nextdate", handlers.NextDateHandler)
	cNR.Post("/api/task", handlers.TaskHandler)
	cNR.Get("/api/tasks", handlers.HandleTaskGet(repository))
	cNR.Put("/api/task", handlers.HandleTaskPut(repository))
	cNR.Delete("/api/task", handlers.HandleTaskDelete(repository))
	cNR.Get("/api/task", handlers.HandleTaskID(repository))
	cNR.Post("/api/task/done", handlers.HandleTaskDone(repository))

	port := ":7540"
	err := http.ListenAndServe(port, cNR)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
