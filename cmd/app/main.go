package main

import (
	"awesomeProject1/internal/database"
	"awesomeProject1/internal/handlers"
	"awesomeProject1/internal/taskService"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHendler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/message", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/message", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/message/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/message/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
