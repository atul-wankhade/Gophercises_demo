package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Task",
	Short: "Task is a cli task manager",
}
