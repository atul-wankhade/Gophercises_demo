package commands

import (
	"fmt"
	"path/filepath"

	"github.com/atul-wankhade/Gophercises/secret/locker"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var encodingkey string

//getCmd is a get command which gets secret from your secret stoarage
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets the secret from your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		l := locker.File(encodingkey, secretsPath())
		key := args[0]
		value, err := l.Get(key)
		if err != nil {
			fmt.Println("no value set")
			return
		}
		fmt.Printf("%s = %s \n", key, value)
	}}

func init() {
	RootCmd.AddCommand(getCmd)
	RootCmd.PersistentFlags().StringVarP(&encodingkey, "key", "k", "", "the key is required to encode and decode secrets")
}

func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secret")
}
