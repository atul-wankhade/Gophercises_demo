package cmd

import (
	"fmt"
	"strings"

	"github.com/atul-wankhade/Gophercises/Task_Manager/db"
	"github.com/spf13/cobra"
)

//addCmd is subcommand to add a new task in the database list
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
