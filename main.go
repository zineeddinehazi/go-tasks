package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	IsDone  bool   `json:"isdone"`
}

var taskList = []Task{
	{ID: "1", Content: "Build CRUD app using Go", IsDone: false},
	{ID: "2", Content: "Prepare for my university exams", IsDone: false},
	{ID: "3", Content: "Drink 30 cups of coffee", IsDone: false},
	{ID: "4", Content: "Complete my lab reports", IsDone: false},
}

func getList(w http.ResponseWriter, r *http.Request) {
	// Tell the browser the content type
	w.Header().Set("Content-Type", "application/json")
	// Encode the Go slice of structs into JSON
	json.NewEncoder(w).Encode(taskList)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	// Handle URL parameters
	vars := mux.Vars(r)
	wantedID := vars["id"]
	for i, task := range taskList {
		if task.ID == wantedID {
			// Sending back updated data as JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(taskList[i])
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	// Handling URL parameters
	vars := mux.Vars(r)
	wantedID := vars["id"]
	for i, task := range taskList {
		if task.ID == wantedID {
			taskList[i].IsDone = !taskList[i].IsDone
			// Sending back updated data as JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(taskList[i])
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	// Handling URL parameters
	vars := mux.Vars(r)
	wantedID := vars["id"]
	for i, task := range taskList {
		if task.ID == wantedID {
			taskList = append(taskList[:i], taskList[i+1:]...)
		}
	}
	// Sending success JSON feedback
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
}

func addTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	// Decode the JSON object into Go struct and store it in the declared variable by referencego
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		// Handling not valid JSON
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	taskList = append(taskList, task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func main() {
	// Setting the router using gorilla library
	r := mux.NewRouter()
	port := ":8080"

	// Setting ifferent endpoints
	r.HandleFunc("/list", getList).Methods("GET")
	r.HandleFunc("/list/{id}", getTask).Methods("GET")
	r.HandleFunc("/list", addTask).Methods("POST")
	r.HandleFunc("/list/{id}", updateTask).Methods("PATCH")
	r.HandleFunc("/list/{id}", deleteTask).Methods("DELETE")

	// Start the server with handling any error
	fmt.Printf("Listening on port %s\n", port)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
