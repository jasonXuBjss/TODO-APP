package main

import (
	"encoding/json"
	"net/http"
	
)

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}


// func addTodoHandler(w http.ResponseWriter, r *http.Request)  {
	
// }

// func updateTodoHandler(w http.ResponseWriter, r *http.Request) {

// }

// func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {

// }