package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("api/echo", echo)
	http.ListenAndServe(port(), nil)
}

// Set PORT via env
func port() string {
	port := os.Getenv("PORT")
	// fallback PORT
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, from Cloud Native Go!!")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content.Type", "text/plain")
	fmt.Fprintf(w, message)
}
