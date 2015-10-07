package main

import (
	"net/http"
	"strings"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSpace(r.URL.Path)

	uriSegments := strings.Split(path, "/")

	// if naked url, ......
	if len(uriSegments[1]) < 1 {
		panic("todo....")
		return
	}

}
