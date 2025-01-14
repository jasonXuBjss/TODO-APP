### Phase 2

* Wrap the Data Store with the V1 REST API.

GET /todos: list all todos.
POST /todos: create a new todo.
PUT /todos/{id}: update a todo.
DELETE /todos/{id}: delete a todo.

http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))




* Use the [context] package to add a TraceID and [slog] to enable traceability of calls through the solution.


* At the ToDo level, use CSP to support concurrent reads and concurrent safe write.



* Use Parallel tests to validate that the solutin is concurrent safe.


* Update the CLI App to use the REST API.


* Add an JSON Data Store and use a startup value to tell the REST API which data store to use.
