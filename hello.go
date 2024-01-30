package main

import (
	"fmt"
	"net/http"
)

var version string

func main() {

	version := "1.0.0"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<H1>Hi folks, welcome to DevOpsLP!</h1><br> Version: %s", version)
	})

	http.ListenAndServe(":8080", nil)
}
