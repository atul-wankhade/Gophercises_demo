package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDevMw(t *testing.T) {
	mux := http.NewServeMux()
	response := devMw(mux)
	//var x http.HandlerFunc
	if response == nil {
		fmt.Println("test failed..")
	}
}

func TestFuncThatPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// The following is the code under test
	funcThatPanics()
}

func TestPanicAfterDemo(t *testing.T) {
	var w http.ResponseWriter
	var r *http.Request
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// The following is the code under test
	panicAfterDemo(w, r)
}

func TestHello(t *testing.T) {
	var w http.ResponseWriter
	var r *http.Request
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// The following is the code under test
	hello(w, r)
}

func TestMakeLinks(t *testing.T) {
	test := "	C:/Users/GSC-30564/go/src/github.com/atul-wankhade/Gophercises/recover/main.go:70 +0xaa"

	//response := " 	Users/GSC-30564/go/src/github.com/atul-wankhade/Gophercises/recover/main.go:70:70 +0xaa"

	response := makeLinks(test)

	if response != " 	Users/GSC-30564/go/src/github.com/atul-wankhade/Gophercises/recover/main.go:70:70 +0xaa" {
		fmt.Println("test failed...")
	}

}

func TestSourceCodeHandler(t *testing.T) {
	//var w http.ResponseWriter
	var req http.Request
	req.FormValue("path")

}
