package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var storage Storage

func main() {
	storeType := flag.String("store", "json", "Specify the storage type: 'json' or 'memory'")
	jsonFile := flag.String("file", "todos.json", "Specify the JSON file for storage (used with 'json' store)")
	flag.Parse()

	switch *storeType {
	case "json":
		storage = &JSONStorage{FileName: *jsonFile}
	case "memory":
		storage = &InMemoryStorage{}
	default:
	}

	if todos, err := storage.Load(); err == nil {
		fmt.Printf("Loaded %d todos from storage.\n", len(todos))
	} else {
		fmt.Println("No todos loaded or error:", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /todos", getTodosHandler)             
	mux.HandleFunc("POST /todos", addTodoHandler)               
	mux.HandleFunc("PUT /todos/{id}", updateTodoHandler)             
	mux.HandleFunc("DELETE /todos/{id}", deleteTodoHandler)
	mux.HandleFunc("PATCH /todos/{id}", toggleTodoHandler)    		

	
	fmt.Println("up n running on 8080")
	log.Fatal(http.ListenAndServe(":8080", Middleware(mux)))

}











