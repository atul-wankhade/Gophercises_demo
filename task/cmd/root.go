package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd is a root command
var RootCmd = &cobra.Command{
	Use:   "mytask",
	Short: "Task is a cli task manager",
}
