package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
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

func readTasksFromJson(filename string, todos *[]ToDo) {
	var tasks []ToDo
	var content = readJsonFile(filename)

	err := json.Unmarshal(content, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	*todos = append(*todos, tasks...)

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

// createTask adds an album from JSON received in the request body.
func createTask(c *gin.Context) {
	var newTask ToDo

	// Call BindJSON to bind the received JSON to
	// newTask.
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	newTask.Id = findLatestId()

	tasks = append(tasks, newTask)
	writeTasksToJsonFile("tasks.json", tasks...)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func updateTask(c *gin.Context) {
	var newValues ToDo
	// Call BindJSON to bind the received JSON to
	// newValues.
	if err := c.BindJSON(&newValues); err != nil {
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	_, i, _ := findTask(id)

	tasks[i].Description = newValues.Description
	tasks[i].Done = newValues.Done

	//tasks = append(tasks, newValues)
	writeTasksToJsonFile("tasks.json", tasks...)
	c.IndentedJSON(http.StatusCreated, tasks[i])
}

func findLatestId() int {
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].Id > tasks[j].Id
	})

	return tasks[0].Id + 1
}

func getTaskById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)

	for _, a := range tasks {
		if a.Id == i {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func main() {
	filename := "tasks.json"

	readTasksFromJson(filename, &tasks)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tasks.tmpl", getSortedTasks())
	})

	//router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskById)
	router.POST("/tasks", createTask)
	router.PUT("/tasks/:id", updateTask)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func getSortedTasks() []ToDo {
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].Id < tasks[j].Id
	})

	return tasks
}

func genTemplate() {
	funcMap := template.FuncMap{
		"dec":     func(i int) int { return i - 1 },
		"replace": strings.ReplaceAll,
	}
	var tmplFile = "tasks.tmpl"
	tmpl, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	var f *os.File
	f, err = os.Create("pets.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, tasks)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
