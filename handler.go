package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)


func getTodosHandler(w http.ResponseWriter, r *http.Request) {

	uuid, ok := uuidFromContext(r.Context())
	if !ok {
		http.Error(w, "no uuid in context", http.StatusInternalServerError)
		return
	} 

	slog.Info("GET /todos",
	slog.String("uuid", uuid),
	slog.String("method", r.Method),
)
	
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func addTodoHandler(w http.ResponseWriter, r *http.Request)  {

	uuid, ok := uuidFromContext(r.Context())
	if !ok {
		http.Error(w, "no uuid in context", http.StatusInternalServerError)
		return
	} 

	slog.Info("POST /todos",
	slog.String("uuid", uuid),
	slog.String("method", r.Method),
)
	
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return	
	}
	if newTodo.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	newTodo.ID = newId
	newId++

	todos = append(todos, newTodo)

}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	uuid, ok := uuidFromContext(r.Context())
	if !ok {
		http.Error(w, "no uuid in context", http.StatusInternalServerError)
		return
	} 

	slog.Info("DELETE /todos/{id}",
	slog.String("uuid", uuid),
	slog.String("method", r.Method),
)


	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var indexToDelete = -1
	for i, todo := range todos {
		if todo.ID == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		http.Error(w, fmt.Sprintf("ID %d not found", id), http.StatusNotFound)
		return
	}

	todos = append(todos[:indexToDelete], todos[indexToDelete+1:]...)

	w.WriteHeader(http.StatusNoContent)

}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	uuid, ok := uuidFromContext(r.Context())
	if !ok {
		http.Error(w, "no uuid in context", http.StatusInternalServerError)
		return
	} 

	slog.Info("PUT /todos/{id}",
	slog.String("uuid", uuid),
	slog.String("method", r.Method),
	)

	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTodo Todo
	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if updatedTodo.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Title = updatedTodo.Title
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, fmt.Sprintf("ID %d not found", id), http.StatusNotFound)
}


func toggleTodoHandler(w http.ResponseWriter, r *http.Request) {
	uuid, ok := uuidFromContext(r.Context())
	if !ok {
		http.Error(w, "no uuid in context", http.StatusInternalServerError)
		return
	} 

	slog.Info("PATCH /todos/{id}",
	slog.String("uuid", uuid),
	slog.String("method", r.Method),
	)

	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = !todos[i].Completed
			w.WriteHeader(http.StatusNoContent) 
			return
		}
	}

	http.Error(w, fmt.Sprintf("ID %d not found", id), http.StatusNotFound)
}
