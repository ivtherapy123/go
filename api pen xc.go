package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var task Message

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Message
	DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&task)
	DB.Create(&task)
	json.NewEncoder(w).Encode(task)``
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/post", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)

}
