package main

import (
	"fmt"
	"net/http"
	"os/exec"
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

	// We need to isolate the individual components of the path.
	components := strings.Split(path, "/")

	salt_node := fmt.Sprintf("G@node_type:%s", components[0])

	cmdName := "sudo"
	cmdArgs := []string{"salt", "-C", salt_node, "state.highstate"}

	out, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		fmt.Fprintf(w, "There was an error running command: %s %s %s", cmdName, cmdArgs, err)
		return
	}

	fmt.Fprintf(w, string(out))
	return
}
