package main

import (
	"fmt"
	"net/http"
	"os"
)

func declareHandlers() {
	http.HandleFunc("/", fetchFile)
}

func fetchFile(w http.ResponseWriter, r *http.Request) {
	p := "htdocs" + r.URL.Path
	if p == "" {
		p = "index.htm"
	}
	if _, err := os.Stat("./index.htm"); os.IsNotExist(err) {
		p = "index.html"
	}
	t, e := readFile(p)
	if e != nil || t == nil {
		http.Error(w, "SERVEHERE -- Page not found.", 404)
		return
	}
	fmt.Fprintf(w, "%s", *t)
}
