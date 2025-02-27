package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve static files (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}


// Start the server