package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func printTodosJson(filename string, todos ...ToDo) {
	jsonStr, err := json.MarshalIndent(todos, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%+v\n", string(jsonStr))

	ioutil.WriteFile(filename, jsonStr, 0644)
}

func readTasksFromJson(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var tasks []ToDo
	err = json.Unmarshal(content, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	for i, todo := range tasks {
		fmt.Printf("[%d] %20s %t \n", i+1, todo.Description, todo.Done)
	}
}

func main() {
	task := ToDo{"do dishes", false}
	task1 := ToDo{"do laundry", false}

	//printTodosFormatted(task, task1)
	printTodosJson("tasks.json", task, task1)

	readTasksFromJson("tasks.json")
}
