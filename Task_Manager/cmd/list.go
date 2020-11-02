package cmd

import (
	"fmt"
	"os"

	"github.com/atul-wankhade/Gophercises/Task_Manager/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all your task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no task to complete! Why not take a vacation?")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Println(i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
