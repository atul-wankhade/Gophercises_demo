package db

import (
	"fmt"
	"testing"
)

func TestCreateTask(t *testing.T) {
	task := "testing create task"
	response, err := CreateTask(task)
	if err != nil {
		fmt.Printf("Create task has returned error")
	} else {
		fmt.Println(response)
	}
}

func TestAllTask(t *testing.T) {
	var tasks []Task
	tasks, err := AllTasks()
	if err != nil {
		fmt.Println("alltask has returned error")
	}
	if len(tasks) == 0 {
		fmt.Println("You have no task to complete! Why not take a vacation?")
		return
	}
	fmt.Println("You have the following tasks:")
	for i, task := range tasks {
		fmt.Println(i+1, task.Value)
	}
}
