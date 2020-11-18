package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/atul-wankhade/Gophercises/transform/processor"
)

func WelcomePage(w http.ResponseWriter, req *http.Request) {
	html := `
		<html><body>
		<form action="/upload"
			enctype="multipart/form-data" method="post">
		<p>
			Please specify a file:<br>
			<input type="file" name="file" size="40">
		</p>
		<div>
		<input type="submit" value="Send">
		</div>
		</form>
		</body></html>`
	fmt.Fprint(w, html)
}

// <label for="mode">Choose a mode:</label>
// <select id="mode" name="mode">
// <option value="1">Triangle</option>
// <option value="2">Rectangle</option>
// <option value="3">Ellipse</option>
// <option value="4">Circle</option>
// <option value="5">Rotated Rectangle</option>
// <option value="6">Baziers</option>
// <option value="7">Rotated Ellipse</option>
// <option value="8">Polygon</option>
// </select>
// <label for="shapes"><b>Shapes</b></label>
// <input type="number" placeholder="Enter number of shapes" name="shapes" id="shapes" min="10" max="500" required>

func Upload1(w http.ResponseWriter, req *http.Request) {

}

func Upload(w http.ResponseWriter, req *http.Request) {

	file, header, err := req.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	ext := filepath.Ext(header.Filename)

	onDiskFile, err := processor.MyTempFile("in_", ext, "./images/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer onDiskFile.Close()
	// mode, err := strconv.Atoi(req.FormValue("mode"))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// shapes, err := strconv.Atoi(req.FormValue("shapes"))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	_, err = io.Copy(onDiskFile, file)
	shapes := 50
	mode := 2

	image, err := createImage(file, ext, shapes, processor.Mode(mode))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	html := `<html><body>
		<a>
			<img style="width: 20%;" src="/images/{{.Name}}">
		</a>
	</body></html>`

	mytemplate := template.Must(template.New("").Parse(html))

	type placeholder struct {
		Name string
	}
	var placeholders placeholder
	placeholders.Name = filepath.Base(image)

	mytemplate.Execute(w, placeholders)
	//http.Redirect(w, req, "/images/"+filepath.Base(image), http.StatusFound)
}

//createImage transform incoming image into some out file and return its path
func createImage(r io.Reader, ext string, numShapes int, mode processor.Mode) (string, error) {
	out, err := processor.Transform(r, ext, numShapes, mode)
	if err == nil {
		outFile, err := processor.MyTempFile("out_", ext, "./images/")
		if err == nil {
			defer outFile.Close()
			io.Copy(outFile, out)
			return outFile.Name(), nil
		}
	}
	return "", err
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", WelcomePage).Methods("GET")
	router.HandleFunc("/upload", Upload).Methods("POST")
	router.HandleFunc("/upload1", Upload1).Methods("POST")
	fs := http.FileServer(http.Dir("./images/"))
	router.Handle("/images/", http.StripPrefix("/images/", fs))
	http.ListenAndServe(":3000", router)
}
