package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var task Message
var nextID = 1

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Message
	DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&task)
	DB.Create(&task)
	json.NewEncoder(w).Encode(task)

}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var task Message

	DB.First(&task, id)

	json.NewDecoder(r.Body).Decode(&task)

	DB.Save(&task)
	json.NewEncoder(w).Encode(task)

}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var task Message
	
	DB.First(&task, id)
	DB.Delete(&task)

	json.NewEncoder(w).Encode(task)
}
func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/post", PostHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", DeleteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)

}
