package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hellWorld)
	http.ListenAndServe(":8080", nil)
}

func hellWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from Go in docker 🐳🐳")
}