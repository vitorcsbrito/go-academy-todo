package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
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

func readTasksFromJson(filename string, todos []ToDo) {
	var tasks []ToDo
	var content = readJsonFile(filename)

	err := json.Unmarshal(content, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	todos = tasks

	printTodosFormatted(tasks)
}

func updateTaskDescription(id int, description string, ch chan []ToDo) {
	task, i, err := findTask(id)

	if err != nil {
		log.Fatal(err.Error())
	}

	time.Sleep(10 * time.Millisecond)

	task.Description = description
	tasks[i] = task

	ch <- []ToDo{task}
}

func findTask(id int) (ToDo, int, error) {

	for i, todo := range tasks {
		if todo.Id == id {
			return todo, i, nil
		}
	}
	return ToDo{-1, "", false}, -1, errors.New("math: square root of negative number")
}

var tasks []ToDo

func main() {
	tasks = append(tasks,
		ToDo{0, "do dishes", false},
		ToDo{1, "do laundry", false})

	filename := "tasks.json"

	ch := make(chan []ToDo)

	writeTasksToJsonFile(filename, tasks...)
	readTasksFromJson(filename, tasks)

	go updateTaskDescription(1, "one", ch)
	_ = <-ch

	go updateTaskDescription(1, "two", ch)
	_ = <-ch

	go updateTaskDescription(1, "three", ch)
	_ = <-ch

	go updateTaskDescription(1, "four", ch)
	_ = <-ch

	time.Sleep(1 * time.Second)

	updatedTask, _, _ := findTask(1)

	println(updatedTask.Description)
}
