package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ping(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp := "OK"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func img(w http.ResponseWriter, req *http.Request) {
	if _, err := os.Stat("/tmp/ok"); err == nil {
		fb, _ := ioutil.ReadFile("img.png")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fb)
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintln(w, "503 service unavailable")
	}
}

func main() {
	port := "8880"
	log.Printf("The server is running on localhost port %s", port)

	// ping
	http.HandleFunc("/", ping)

	// img
	http.HandleFunc("/img", img)

	// listen port
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
