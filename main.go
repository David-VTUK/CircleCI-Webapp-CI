package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "End to end")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":80", nil)
}
