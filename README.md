# TODO App

A Go project implementing a simple REST API for managing todos.  *(Work in Progress)*

This app supports both in-memory storage and JSON-based file storage, with the ability to choose the storage type at runtime.



---

## Features

- REST API for managing todos (`GET`, `POST`, `PUT`, `DELETE`, `PATCH`).
- Supports JSON file storage for persistent data.
- Option to use in-memory storage.
- UUID middleware for request tracing.
- Thread-safe operations with `sync.Mutex`.

---

## Installation

### Clone the Repository
```bash
git clone https://github.com/jasonXuBjss/TODO-APP.git

cd todo-app

go mod tidy
```

### Run the App

#### Start the server with JSON storage:
```bash
go run . -store=json -file=todos.json
```
#### Start the server with  in-memory storage:
```bash
go run . -store=memory
```
---
### Project Stucture
```bash
todo-app/
├── context.go       # Middleware for UUID management
├── handler.go       # HTTP handler functions
├── inMemStorage.go  # In-memory storage implementation
├── jsonStorage.go   # JSON storage implementation with sync.Mutex
├── main.go          # Entry point of the application
├── todo.go          # Todo-related helper functions
├── todos.json       # Persistent JSON data
├── types.go         # Todo struct definition
├── go.mod           # Go module file
└── go.sum           # Dependency checksums
```
---
### API Endpoints

| Method | Endpoint        | Description                     |
|--------|------------------|---------------------------------|
| GET    | `/todos`         | Retrieve all todos.            |
| POST   | `/todos`         | Add a new todo.                |
| PUT    | `/todos/{id}`    | Update an existing todo.       |
| DELETE | `/todos/{id}`    | Delete a todo.                 |
| PATCH  | `/todos/{id}`    | Toggle the completion status.  |
