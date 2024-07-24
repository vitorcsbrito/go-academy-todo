package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	jsonStr, err := json.MarshalIndent(todos, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%+v\n", string(jsonStr))

	ioutil.WriteFile("config.json", jsonStr, 0644)
}

func main() {
	task := ToDo{"do dishes", false}
	task1 := ToDo{"do laundry", false}

	//printTodosFormatted(task, task1)
	printTodosJson(task, task1)
}
