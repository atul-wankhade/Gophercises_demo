package encrypt

import (
	"os"
	"testing"
)

func TestEncryptorWriter(t *testing.T) {
	key := "your key"
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Error("failed to get io writer")
	}
	stream, err := EncryptorWriter(key, file)
	if err != nil || stream == nil {
		t.Error("failed to get encryptor writer")
	}
	defer file.Close()
}

func TestDecryptorReader(t *testing.T) {
	key := "decrypt key"
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Error("failed to get io reader")
	}
	reader, err := DecryptorReader(key, file)
	if err != nil || reader == nil {
		t.Error("failed to get decrypt reader")
	}
	file.Close()
}
