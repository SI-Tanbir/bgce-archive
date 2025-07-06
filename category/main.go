package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World with ServeMux!")
}

// class note route functions handaler
func classnoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "class-note is here - ")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/class-note", classnoteHandler) // class note route

	port := ":8080"
	fmt.Println("Server is running on http://localhost" + port)

	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
