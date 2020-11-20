package main

import (
	"bytes"
	"fmt"
	"io"
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

func TestUpload(t *testing.T) {
	go main()
	time.Sleep(10)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	f, _ := os.Open("monalisa.png")
	values := map[string]io.Reader{
		"myFile": f,
	}
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		var err error
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return
		}
		if err = w.WriteField("mode", "3"); err != nil {
			return
		}
		if err = w.WriteField("shapes", "50"); err != nil {
			return
		}
		w.Close()
		req, err := http.NewRequest("POST", "http://localhost:3000/upload", &b)
		if err != nil {
			t.Errorf("TestUpload: " + err.Error())
		}
		req.Header.Set("Content-Type", w.FormDataContentType())

		// Submit the request
		res, err := client.Do(req)
		if err != nil {
			t.Errorf("TestUpload: Failed to upload image")
			return
		}
		if res.StatusCode != 200 {
			t.Errorf("TestUpload: Failed to upload image with status code = %d", res.StatusCode)
			return
		}
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
