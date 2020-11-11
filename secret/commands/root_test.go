package commands

import (
	"testing"
)

func TestRootCmd(t *testing.T) {
	//file, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	RootCmd.Run(RootCmd, []string{"get", "atul"})
	//file.Seek(0, 0)
	// _, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	t.Error("expected nil but got ", err)
	// }
	// if string(content) != "value set successfully.." {
	// 	t.Error("test failed")
	// }
	//defer file.Close()
}
