package main

import (

	"fmt"
)

type Todo struct {
	ID			int		`json:"id"`
	Title		string	`json:"title"`
	Completed	bool	`json:"completed"`
}

type Todos []Todo

var todos = Todos{
	{ID: 1, Title: "learn some go", Completed: true},
	{ID: 2,Title: "build an api", Completed: false},
	{ID: 3,Title: "testing the app", Completed: false},
}
var newId = len(todos) + 1

func (todos *Todos) Add(title string) {
	todo := Todo{
		ID: newId,
		Title:     title,
		Completed: false,
	}
	newId++
	*todos = append(*todos, todo)
	fmt.Println("new to-do added.")
}

func (todos *Todos) Delete(id int) error {
	for i := range *todos {
		if (*todos)[i].ID == id {
			*todos = append((*todos)[:i], (*todos)[i+1:]...)
			fmt.Println("to-do deleted.")
			return nil
		}
	}
	return fmt.Errorf("id %d not found", id)
}

func (todos *Todos) Update(id int, title string) error {
	for i := range *todos {
		if (*todos)[i].ID == id {
			(*todos)[i].Title = title
			fmt.Println("to-do updated.")
			return nil
		}
	}
	return fmt.Errorf("id %d not found", id)
}

func (todos *Todos) Toggle(id int) error {
	for i := range *todos {
		if (*todos)[i].ID == id {
			(*todos)[i].Completed = !(*todos)[i].Completed
			fmt.Println("to-do taggled.")
			return nil
		}
	}
	return fmt.Errorf("id %d not found", id)
}

func (todos *Todos) GetList(){
	for _, todo := range *todos {
		status := "Incomplete"
		if todo.Completed {
			status = "Complete"
		}
		fmt.Printf("ID: %d, Title: %s, Status: %s\n", todo.ID, todo.Title, status)
	}
}