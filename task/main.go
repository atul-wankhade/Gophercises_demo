package main

//import "Task_Manager/cmd"

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/atul-wankhade/Gophercises/Task_Manager/cmd"
	"github.com/atul-wankhade/Gophercises/Task_Manager/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	var Home, _ = homedir.Dir()
	dbPath := filepath.Join(Home, "tasks.db")
	must(db.Init(dbPath))
	fmt.Println("Connected to boltDB!")
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
