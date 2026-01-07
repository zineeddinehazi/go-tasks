package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetRouter() (*mux.Router, *sql.DB) {

	// Open connection to database ---------------------------------------------------------------------------
	dbConnection, err := sql.Open("sqlite3", "./tasks.db")
	handleFatal(err)
	taskRepository := &TaskRepository{DB: dbConnection}

	// Create table
	err = taskRepository.CreateTable()
	handleFatal(err)

	// Router -----------------------------------------------------------------------------------------------------------------------
	var r = mux.NewRouter()

	// Set different endpoints ---------------------------------------------------------------------------------------------------

	// Get list -------------------------------------------------------------------------------------------------------------------------------------------
	r.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		list, err := taskRepository.QuerryList()
		handleFatal(err)
		w.Header().Set("Content-Type", "application/json")
		if len(list) == 0 {
			json.NewEncoder(w).Encode(map[string]string{"message": "List empty"})
			return
		}
		json.NewEncoder(w).Encode(list)
	}).Methods("GET")

	// Get task -------------------------------------------------------------------------------------------------------------------------------------------
	r.HandleFunc("/list/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		wantedID := vars["id"]
		task, err := taskRepository.QuerryTask(wantedID)
		handleFatal(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}).Methods("GET")

	// Add task -------------------------------------------------------------------------------------------------------------------------------------------
	r.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		err := taskRepository.InsertTask(task)
		handleFatal(err)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{"content": task.Content, "isdone": task.IsDone})
	}).Methods("POST")

	// Update task -------------------------------------------------------------------------------------------------------------------------------------------
	r.HandleFunc("/list/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Handling URL parameters
		vars := mux.Vars(r)
		wantedID := vars["id"]
		err := taskRepository.UpdateTask(wantedID)
		handleFatal(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Task with id %s updated", wantedID)})

	}).Methods("PATCH")

	// Delete task -------------------------------------------------------------------------------------------------------------------------------------------
	r.HandleFunc("/list/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		wantedID := vars["id"]
		err := taskRepository.DeleteTask(wantedID)
		handleFatal(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Task with id %s deleted", wantedID)})
	}).Methods("DELETE")

	return r, dbConnection
}
