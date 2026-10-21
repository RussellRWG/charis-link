package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("public"))
	mux.Handle("/", fileServer)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Starting server on port 8080")
	srv.ListenAndServe()
}
