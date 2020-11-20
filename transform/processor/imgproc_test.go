package processor

import (
	"os"
	"testing"
)

func TestTransform(t *testing.T) {
	file, _ := os.Open("../monalisa.png")
	reader, err := Transform(file, ".png", 10, ModeCircle)
	if err != nil || reader == nil {
		t.Error("TestTransform: " + err.Error())
		return
	}
}
