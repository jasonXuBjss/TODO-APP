package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", getTodosHandler)              // GET /todos 
	// mux.HandleFunc("/todos", addTodoHandler)                // POST /todos 
	// mux.HandleFunc("/todos/", updateTodoHandler)            // PUT /todos/{id} 
	// mux.HandleFunc("/todos/", deleteTodoHandler)    		//DEL /todos/{id}


	fmt.Println("up n running on 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}











