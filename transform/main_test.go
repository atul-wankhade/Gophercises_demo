package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestRoot(t *testing.T) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	go main()
	time.Sleep(10)
	req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
	if err != nil {
		t.Errorf("TestUpload: Failed to get upload form")
	}
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		t.Errorf("TestUpload: Failed to get upload form with response code = %d", resp.StatusCode)
	}
}

func TestUpload1(t *testing.T) {
	go main()
	time.Sleep(10)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	file, err := os.Open("monalisa.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	//buffer to store our request body as buffer
	var requestBody bytes.Buffer

	//create a multipart writer
	multipartWriter := multipart.NewWriter(&requestBody)

	//initialize the file feild
	fileWriter, err := multipartWriter.CreateFormFile("myFile", file.Name())
	if err != nil {
		log.Fatalln(err)
	}

	//copy the actual content of the file to the file feilds writer
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		log.Fatalln(err)
	}

	feildWriter, err := multipartWriter.CreateFormField("mode")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = feildWriter.Write([]byte("2"))
	if err != nil {
		log.Fatalln(err)
	}

	feildWriter1, err := multipartWriter.CreateFormField("shapes")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = feildWriter1.Write([]byte("50"))
	if err != nil {
		log.Fatalln(err)
	}

	//we completed adding file and feilds to multipartWriter so now we close it so that it can writes the ending boundry
	multipartWriter.Close()

	//now we use our original populated request body with our custom request
	req, err := http.NewRequest("POST", "http://localhost:3000/upload", &requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	//we need to set the content type from the writer , it includes necessary boundary as well
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	//do the request
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if response.StatusCode != 200 {
		t.Errorf("TestUpload: Failed to upload image with status code = %d", response.StatusCode)
		return
	}
}

func TestUploadBadRequest(t *testing.T) {
	go main()
	time.Sleep(10)
	resp, err := followURL("POST", "http://localhost:3000/upload")
	if err != nil {
		t.Errorf("TestUploadBadRequest: Failed to get upload form")
	}
	if resp.StatusCode != 400 {
		t.Errorf("TestUploadBadRequest: Failed to get upload form with response code = %d", resp.StatusCode)
	}
}

func followURL(method, path string) (*http.Response, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	var req *http.Request
	var resp *http.Response
	var err error
	req, _ = http.NewRequest(method, path, nil)
	resp, err = client.Do(req)
	fmt.Printf("Request is %v\n", req)
	fmt.Printf("Response is %v\n", resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
