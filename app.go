package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func printTodosFormatted(todos []ToDo) {
	for i, todo := range todos {
		fmt.Printf("[%d] %20s %t \n", i+1, todo.Description, todo.Done)
	}
}

func writeTasksToJsonFile(filename string, todos ...ToDo) {
	jsonStr, err := json.MarshalIndent(todos, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	fileExists := checkFileExists(filename)
	if !fileExists {
		createFile(filename)
	}

	writeFile(filename, string(jsonStr))
}

func readTasksFromJson(filename string) {
	var tasks []ToDo
	var content = readJsonFile(filename)

	err := json.Unmarshal(content, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	printTodosFormatted(tasks)
}

func main() {
	task := ToDo{"do dishes", false}
	task1 := ToDo{"do laundry", false}

	filename := "tasks.json"

	writeTasksToJsonFile(filename, task, task1)
	readTasksFromJson(filename)
}
