package main

import (
	"log"
	"net/http"
	"strings"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSpace(r.URL.Path)

	//Cut off the leading and trailing forward slashes, if they exist.
	//This cuts off the leading forward slash.
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	//This cuts off the trailing forward slash.
	if strings.HasSuffix(path, "/") {
		cut_off_last_char_len := len(path) - 1
		path = path[:cut_off_last_char_len]
	}

	if len(path) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//We need to isolate the individual components of the path.
	components := strings.Split(path, "/")

	log.Println(components[0])
}
