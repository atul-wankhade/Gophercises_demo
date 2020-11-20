package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestSetCmd(t *testing.T) {
	file, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	stdOut := os.Stdout
	os.Stdout = file
	setCmd.Run(setCmd, []string{"test-key", "test-value"})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("expected nil but got ", err)
	}
	if true != strings.Contains(string(content), "value set successfully..") {
		t.Error("test failed")
	}
	os.Stdout = stdOut
	defer os.Remove("test.txt")
}

func TestSetCmdNegativeTwo(t *testing.T) {
	file, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	stdOut := os.Stdout
	os.Stdout = file
	setCmd.Run(setCmd, []string{""})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("Expected nil but got", err)
	}
	val := strings.Contains(string(content), "Not enough arguments")
	if val != true {
		fmt.Println("Unexpected output.")
	}
	os.Stdout = stdOut
	defer os.Remove("test.txt")
}
