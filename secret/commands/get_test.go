package commands

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetCmd(t *testing.T) {
	file, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	getCmd.Run(getCmd, []string{"test-key"})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("expected nil but got ", err)
	}
	if string(content) != "value set successfully.." {
		t.Error("test failed")
	}
	defer file.Close()
}
