//Package main consists of recovery middleware which panics and recover with links to the source code file which has panicked.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/", sourceCodeHandler)
	mux.HandleFunc("/panic/", panicDemo)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), returnHandlerFunc(mux)))
}

//sourceCodeHandler renders the source code with highlighted lines which has panicked.
func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	lineStr := r.FormValue("line")
	var fileRead []byte
	line, err := strconv.Atoi(lineStr)
	if err != nil {
		line = -1
	}
	file, err := os.Open(path) //for read access
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}

	fileRead, err = ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, string(fileRead))
	style := styles.Get("dracula")
	formatter := html.New(html.TabWidth(2), html.WithLineNumbers(true), html.HighlightLines([][2]int{{line, line}}))

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>")
	formatter.Format(w, style, iterator)

}

func returnHandlerFunc(app http.Handler) http.HandlerFunc {
	return func(respWriter http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				respWriter.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(respWriter, "<h1>panic: %v</h1><pre>%s</pre>", err, createLinks(string(stack)))
			}
		}()
		app.ServeHTTP(respWriter, req)
	}
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	panic("Oh no!")
}

//createLinks function accepts stack trace in string format and returns string with links to the file which has panicked with line number
func createLinks(stack string) string {
	lines := strings.Split(stack, "\n")
	for li, line := range lines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}
		file := ""
		c := 0
		for i, ch := range line {

			if ch == ':' {
				c = c + 1
			}
			if ch == ':' && c > 1 {
				file = line[1:i]
				break
			}
		}

		var lineStr strings.Builder
		for i := len(file) + 2; i < len(line); i++ {
			if line[i] < '0' || line[i] > '9' {
				break
			}
			lineStr.WriteByte(line[i])
		}
		v := url.Values{}
		v.Set("path", file)
		v.Set("line", lineStr.String())
		lines[li] = "\t<a href=\"/debug/?" + v.Encode() + "\">" + file + ":" + lineStr.String() + "</a>" + line[len(file)+2+len(lineStr.String()):]
	}
	return strings.Join(lines, "\n")
}
