package commands

import "github.com/spf13/cobra"

//RootCmd is a root command
var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "secret is a cli to save the secret key-value",
}
