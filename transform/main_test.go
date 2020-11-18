package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/atul-wankhade/Gophercises/transform/processor"
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
		"file": f,
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

func TestCreateImage(t *testing.T) {
	input, err := os.Open("monalisa.png")
	if err != nil {
		t.Errorf("Failed to open image file")
	}
	out, err := createImage(input, "png", 100, processor.ModeCombo)
	if err != nil || out == "" {
		t.Errorf("Failed to transform image")
	}

}

// func TestCreateImages(t *testing.T) {
// 	input, err := os.Open("monalisa.png")
// 	if err != nil {
// 		t.Errorf("failed to open image file")
// 	}
// 	out, err := createImages(input, "png", Combinations{shapes: 10, mode: processor.ModeCombo})
// 	if err != nil || len(out) < 1 {
// 		t.Errorf("Failed to transform images")
// 	}
// }
