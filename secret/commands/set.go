package commands

import (
	"fmt"

	"github.com/atul-wankhade/Gophercises/secret/locker"
	"github.com/spf13/cobra"
)

//RootCmd is a root command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "sets the secret key",
	Run: func(cmd *cobra.Command, args []string) {
		l := locker.File(encodingkey, secretsPath())
		var key, value string
		key, value = args[0], args[1]
		err := l.Set(key, value)
		if err != nil {
			fmt.Println("Operation failed....")
			return
		}
		fmt.Println("value set successfully..")
	}}

func init() {
	RootCmd.AddCommand(setCmd)
}
