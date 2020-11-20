package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestGetCmd(t *testing.T) {
	file, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	stdOut := os.Stdout
	os.Stdout = file
	getCmd.Run(getCmd, []string{"test-key"})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("expected nil but got ", err)
	}
	if true != strings.Contains(string(content), "tesy-key = test-value") {
		t.Error("test failed")
	}
	os.Stdout = stdOut
	defer os.Remove("test.txt")
}

func TestGetCmdNoValueSet(t *testing.T) {
	file, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	stdOut := os.Stdout
	os.Stdout = file
	getCmd.Run(getCmd, []string{"Test"})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("Expected nil but got", err)
	}
	val := strings.Contains(string(content), "no value set")
	if val != true {
		fmt.Println("Unexpected output.")
	}
	fmt.Println(string(content))
	os.Stdout = stdOut
	defer os.Remove("test.txt")
}
