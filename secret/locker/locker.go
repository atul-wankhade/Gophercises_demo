package locker

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/atul-wankhade/Gophercises/secret/encrypt"
)

//Locker structure to store secret file
type Locker struct {
	encodingKey string
	filepath    string
	keyValues   map[string]string
}

//File initializes locker
func File(encodingKey string, filepath string) Locker {
	return Locker{
		encodingKey: encodingKey,
		filepath:    filepath,
		keyValues:   make(map[string]string),
	}
}

func (l *Locker) save() error {
	f, err := os.OpenFile(l.filepath, os.O_RDWR|os.O_CREATE, 0775)
	if err != nil {
		return err
	}

	w, err := encrypt.EncryptorWriter(l.encodingKey, f)
	if err != nil {
		return err
	}
	return l.writeKeyValues(w)
}

func (l *Locker) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(l.keyValues)
}

func (l *Locker) load() error {
	f, err := os.Open(l.filepath)
	if err != nil {
		l.keyValues = make(map[string]string)
		return nil
	}

	defer f.Close()

	r, err := encrypt.DecryptorReader(l.encodingKey, f)
	if err != nil {
		return err
	}
	return l.readKeyValues(r)
}

func (l *Locker) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&l.keyValues)
}

//Set value for key
func (l *Locker) Set(key, value string) error {
	err := l.load()
	if err != nil {
		return err
	}
	l.keyValues[key] = value
	return l.save()
}

//Get the value set inside locker
func (l *Locker) Get(key string) (string, error) {
	err := l.load()
	if err != nil {
		return "", err
	}
	value, ok := l.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for that key")
	}
	return value, nil
}
