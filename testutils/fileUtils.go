package testutils

import (
	"fmt"
	"os"
	"time"
)

func GetTestDbFileName() string {
	return fmt.Sprintf("test_%d.json", time.Now().Unix())
}

func Cleanup(filename string) {
	err := os.Remove(filename)

	if err != nil {
		panic(err)
	}
}
