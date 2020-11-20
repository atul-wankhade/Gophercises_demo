package processor

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// func TestTransform(t *testing.T) {
// 	file, _ := os.Open("../monalisa.png")
// 	reader, err := Transform(file, ".png", 10, ModeCircle)
// 	if err != nil || reader == nil {
// 		t.Error("TestTransform: " + err.Error())
// 		return
// 	}
// }

func TestPrimitive(t *testing.T) {
	file, _ := os.Open("../monalisa.png")
	out, err := ioutil.TempFile("images", "result_*.png")
	string, err := Primitive(file.Name(), out.Name(), 50, ModeBeziers)
	if err != nil {
		fmt.Println("failed to transform image")
	}
	fmt.Println(string)
}
