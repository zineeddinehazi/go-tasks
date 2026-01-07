package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	// Tell the browser the content type
	w.Header().Set("Content-Type", "application/json")
	// Encode the Go slice of structs into JSON
	json.NewEncoder(w).Encode(TaskList)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	// Handle URL parameters
	vars := mux.Vars(r)
	wantedID := vars["id"]
	for i, task := range TaskList {
		if task.ID == wantedID {
			// Sending back updated data as JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(TaskList[i])
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Handling URL parameters
	vars := mux.Vars(r)
	wantedID := vars["id"]
	for i, task := range TaskList {
		if task.ID == wantedID {
			TaskList[i].IsDone = !TaskList[i].IsDone
			// Sending back updated data as JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(TaskList[i])
		}
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Handling URL parameters
	vars := mux.Vars(r)
	wantedID := vars["id"]
	for i, task := range TaskList {
		if task.ID == wantedID {
			TaskList = append(TaskList[:i], TaskList[i+1:]...)
		}
	}
	// Sending success JSON feedback
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	// Decode the JSON object into Go struct and store it in the declared variable by referencego
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		// Handling not valid JSON
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	TaskList = append(TaskList, task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
