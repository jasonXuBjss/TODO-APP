package main

import (
	"fmt"
	"log"
	"net/http"
	
	
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /todos", getTodosHandler)             
	mux.HandleFunc("POST /todos", addTodoHandler)               
	mux.HandleFunc("PUT /todos/{id}", updateTodoHandler)             
	mux.HandleFunc("DELETE /todos/{id}", deleteTodoHandler)
	mux.HandleFunc("PATCH /todos/{id}", toggleTodoHandler)    		

	
	fmt.Println("up n running on 8080")
	log.Fatal(http.ListenAndServe(":8080", Middleware(mux)))

}











