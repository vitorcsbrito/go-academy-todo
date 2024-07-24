package main

import (
	"encoding/json"
	"fmt"
)

type ToDo struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func printTodosFormatted(todos ...ToDo) {
	for i, todo := range todos {
		fmt.Printf("[%d] %20s %t \n", i+1, todo.Description, todo.Done)
	}
}

func printTodosJson(todos ...ToDo) {
	for _, todo := range todos {
		jsonStr, _ := json.Marshal(todo)
		fmt.Printf("%+v\n", string(jsonStr))
	}
}

func main() {
	task := ToDo{"do dishes", false}
	task1 := ToDo{"do laundry", false}

	printTodosFormatted(task, task1)
	printTodosJson(task, task1)
}
