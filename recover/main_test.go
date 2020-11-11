package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDevMw(t *testing.T) {
	mux := http.NewServeMux()
	response := returnHandlerFunc(mux)
	//var x http.HandlerFunc
	if response == nil {
		fmt.Println("test failed..")
	}
}

func TestMakeLinks(t *testing.T) {
	test := "	C:/Users/GSC-30564/go/src/github.com/atul-wankhade/Gophercises/recover/main.go:70 +0xaa"

	//response := " 	Users/GSC-30564/go/src/github.com/atul-wankhade/Gophercises/recover/main.go:70:70 +0xaa"

	response := createLinks(test)

	if response != " 	C:/Users/GSC-30564/go/src/github.com/atul-wankhade/Gophercises/recover/main.go:70:70 +0xaa" {
		fmt.Println("test failed...")
	}

}

func TestSourceCodeHandler(t *testing.T) {
	handler := http.HandlerFunc(sourceCodeHandler)
	response, _ := executeRequest("Get", "/debug/?line=95&path=C%3A%2FUsers%2FGSC-30564%2Fgo%2Fsrc%2Fgithub.com%2Fatul-wankhade%2FGophercises%2Frecover%2Fmain.go", returnHandlerFunc(handler))
	//assert.Equal(t, response.Code, 200)
	if http.StatusOK != response.Code {
		fmt.Println("handler responded with code :", response.Code)
	}
}

func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	recorder := httptest.NewRecorder()
	recorder.Result()
	handler.ServeHTTP(recorder, req)
	return recorder, err
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
