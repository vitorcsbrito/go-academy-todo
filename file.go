package main

import (
	"fmt"
	"log"
	"os"
)

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
