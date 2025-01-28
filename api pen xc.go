package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Message
	DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task Message
	json.NewDecoder(r.Body).Decode(&task)
	DB.Create(&task)
	json.NewEncoder(w).Encode(task)

}

func UpdateTaskByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["ID"]
	var task Message

	DB.First(&task, id)

	json.NewDecoder(r.Body).Decode(&task)

	DB.Save(&task)
	json.NewEncoder(w).Encode(task)

}
func DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["ID"]
	var task Message

	DB.First(&task, id)
	DB.Delete(&task)

	json.NewEncoder(w).Encode(task)
}
func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/message", GetAllTasks).Methods("GET")
	router.HandleFunc("/api/message", CreateNewTask).Methods("POST")
	router.HandleFunc("/api/message/{ID}", UpdateTaskByID).Methods("PATCH")
	router.HandleFunc("/api/message/{ID}", DeleteTaskByID).Methods("DELETE")
	http.ListenAndServe(":8080", router)

}
