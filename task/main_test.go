package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/atul-wankhade/Gophercises/task/cmd"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := cmd.RootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(out))
}

func Test_ExecuteSubcommand(t *testing.T) {
	cmd1 := cmd.RootCmd
	b := bytes.NewBufferString("")
	cmd1.SetOut(b)
	//cmd1.SetArgs([]string{"list"})
	cmd1.SetHelpCommand(cmd.ListCmd)
	cmd1.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(out))
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
