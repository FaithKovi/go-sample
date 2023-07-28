package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a file server to serve static files
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	fmt.Println("Server is listening on port 3007...")
	http.ListenAndServe(":3007", nil)
}
