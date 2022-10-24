package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type FileReader interface {
	ReadFile() ([]byte, error)
	Stat() (os.FileInfo, error)
}

type file struct {
	path string
}

func (f file) ReadFile() ([]byte, error)  { return os.ReadFile(f.path) }
func (f file) Stat() (os.FileInfo, error) { return os.Stat(f.path) }

func pingHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp := "OK"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

// imgHandler is an extended handler function that takes an additional
// argument that implements FileReader in order to abstract the implementation
// of the FileReader away from the function and to allow the caller to handle
// the implementation
func imgHandler(w http.ResponseWriter, req *http.Request, f FileReader) {
	if _, err := f.Stat(); err == nil {
		fb, _ := f.ReadFile()
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fb)
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "503 service unavailable")
	}
}

func main() {
	port := "8880"
	log.Printf("The server is running on localhost port %s", port)

	// ping
	http.HandleFunc("/", pingHandler)

	// img
	http.HandleFunc("/img", func(w http.ResponseWriter, r *http.Request) {
		imgHandler(w, r, &file{"/tmp/path"})
	})

	// listen port
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
