package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test CI Webserver 1")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":80", nil)
}
