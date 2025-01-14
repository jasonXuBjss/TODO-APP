package main

import (
	"errors"
	"fmt"
)

type Todo struct {
	Title     string `json: "title"`
	Completed bool   `json: "completed"`
}

type Todos []Todo

var todos = Todos{
	{Title: "learn some go", Completed: true},
	{Title: "build an api", Completed: false},
	{Title: "testing the app", Completed: false},
}


func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:     title,
		Completed: false,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(i int) error {
	if i >= len(*todos) || i < 0 {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) Delete(i int) error {

	if err := todos.validateIndex(i); err != nil {
		return err
	} 

	*todos = append((*todos)[:i], (*todos)[i+1:]... )
		
	return nil
}

func (todos *Todos) Toggle(i int) error {

	if err := todos.validateIndex(i); err != nil {
		return err
	} 

	(*todos)[i].Completed =!(*todos)[i].Completed
		
	return nil
}

func (todos *Todos) Update(i int, title string) error {

	if err := todos.validateIndex(i); err != nil {
		return err
	} 

		(*todos)[i].Title = title
	
		
	return nil
}


func (todos *Todos) GetList(){
	for i, todo := range *todos {
		status := "Incomplete"
		if todo.Completed {
			status = "Complete"
		}
		fmt.Printf("ID: %d, Title: %s, Status: %s\n", i, todo.Title, status)
	}
}