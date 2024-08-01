package files

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)
import . "go-todo-app/model"

func WriteTasksToJsonFile(filename string, todos ...Task) {
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

func ReadTasksFromJson(filename string, todos *[]Task) {
	var tasks []Task
	var content = readJsonFile(filename)

	err := json.Unmarshal(content, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	*todos = append(*todos, tasks...)

	//printTodosFormatted(tasks)
}

func checkFileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func createFile(filename string) {
	file, err := os.Create(filename)

	_closeFile(file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file created")
}

func readJsonFile(filename string) []byte {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done reading")
	return f
}

func writeFile(fileName string, val string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	_, err1 := f.WriteString(val)
	if err1 != nil {
		log.Fatal(err)
	}

	_closeFile(f)

	fmt.Println("done writing")
}

func _closeFile(f *os.File) {
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("error closing file", err)
		}
	}(f)
}
