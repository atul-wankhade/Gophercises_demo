package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/atul-wankhade/Gophercises/transform/processor"
)

//WelcomePage takes input file , no of shapes and mode and send it to upload handler
func WelcomePage(w http.ResponseWriter, req *http.Request) {
	html := `
		<!DOCTYPE html>
		<html><body>
		<form action="/upload" method="post" enctype="multipart/form-data">
		<h3>Select image to upload:	<input type="file" name="myFile"></h3>
		<div>
		<label for="mode">Choose a mode:</label>
		<select id="mode" name="mode">
		<option value="1">Triangle</option>
		<option value="2">Rectangle</option>
		<option value="3">Ellipse</option>
		<option value="4">Circle</option>
		<option value="5">Rotated Rectangle</option>
		<option value="6">Baziers</option>
		<option value="7">Rotated Ellipse</option>
		<option value="8">Polygon</option>
		</select>
		<label for="shapes"><b>Shapes</b></label>
		<input type="number" placeholder="Enter number of shapes" name="shapes" id="shapes" min="10" max="500" required>				</div>
		<br/><br/>
		<h3><input type="submit" value="Upload Image" name="submit"></h3>
		</form>
		</body></html>`
	fmt.Fprint(w, html)
}

//Upload takes the image and transform it and shows the transformed image to user
func Upload(w http.ResponseWriter, r *http.Request) {
	// Just make sure that name you specify here should match with the name in image tag from HTML
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	mode, err := strconv.Atoi(r.FormValue("mode"))
	if err != nil {
		fmt.Println("failed to get mode")
	}

	shapes, err := strconv.Atoi(r.FormValue("shapes"))
	if err != nil {
		fmt.Println("failed to get shapes")
	}

	// logging some details of the file and inputs
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	fmt.Printf("Mode provided: %+v\n", mode)
	fmt.Printf("No of shapes provided: %+v\n", shapes)

	// Create a temporary file within our images directory that follows
	// a particular naming pattern ------- make sure you create one folder named images
	tempFile, err := ioutil.TempFile("images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	outputFile, err := ioutil.TempFile("images", "out_*.png")
	if err != nil {
		fmt.Println(err)
	}

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	in := tempFile.Name()
	out := outputFile.Name()
	mode1 := processor.Mode(mode)

	result, err := processor.Primitive(in, out, shapes, mode1)
	if err != nil {
		fmt.Println("failed to transform image")
	}
	_ = result

	fileBytes2, err := ioutil.ReadAll(outputFile)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	//serveFile.Write(fileBytes2)

	// Render the image on to the browser
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(fileBytes2)))
	if _, err := w.Write(fileBytes2); err != nil {
		log.Println("unable to write image.")
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", WelcomePage).Methods("GET")
	router.HandleFunc("/upload", Upload).Methods("POST")
	http.ListenAndServe(":3000", router)
}
