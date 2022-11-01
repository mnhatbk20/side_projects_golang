package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"todoapi-mongo/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", controllers.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", controllers.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", controllers.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", controllers.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", controllers.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", controllers.DeleteAllTask).Methods("DELETE", "OPTIONS")

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", router))
}
