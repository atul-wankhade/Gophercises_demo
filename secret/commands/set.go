package commands

import (
	"fmt"

	"github.com/atul-wankhade/Gophercises/secret/locker"
	"github.com/spf13/cobra"
)

//setCmd is a set command which sets your key value in secret file
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "sets the secret key",
	Run: func(cmd *cobra.Command, args []string) {
		l := locker.File(encodingkey, secretsPath())
		//Need minimum two arguments
		if len(args) < 2 {
			fmt.Println("Not enough arguments")
			return
		}
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
