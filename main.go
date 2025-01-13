package main

import "fmt"


func main() {
	todos := Todos{}

	todos.Add("buy stuff")
	todos.Add("buy more stuff")
	todos.Add("learn go with test")


	fmt.Printf("%+v\n\n", todos)

	todos.Delete(2)
	fmt.Printf("%+v\n\n", todos)

	todos.Update(1, "learn more go")
	fmt.Printf("%+v\n\n", todos)

	todos.Toggle(1)
	fmt.Printf("%+v\n\n", todos)
}









