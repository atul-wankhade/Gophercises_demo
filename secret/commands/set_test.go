package commands

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSetCmd(t *testing.T) {
	file, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	setCmd.Run(setCmd, []string{"test-key", "test-value"})
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
