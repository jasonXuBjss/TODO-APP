# TODO App

A Go project implementing a simple REST API for managing todos. (WIP)

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

git clone https://github.com/jasonXuBjss/TODO-APP.git

cd todo-app

go mod tidy


### Run the App

#### Start the server with JSON storage:

go run . -store=json -file=todos.json

#### Start the server with  in-memory storage:

go run . -store=memory