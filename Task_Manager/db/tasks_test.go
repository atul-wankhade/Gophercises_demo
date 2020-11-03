package db

import (
	"fmt"
	"testing"
)

func TestCreateTask(t *testing.T) {
	task := "testing create task"
	response, err := CreateTask(task)
	if err != nil {
		fmt.Printf("Create task test has failed")
	}
	fmt.Println(response)
}
