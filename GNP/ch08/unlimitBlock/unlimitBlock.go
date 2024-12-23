package main

import (
	"fmt"
	"net/http"
)

func blockIndefinitely(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi\n")
	select {}
}

var PORT = ":1234"

func main() {
	mux := http.NewServeMux()

	s := &http.Server{
		Addr:    PORT,
		Handler: mux,
		// IdleTimeout:  10 * time.Second,
		// ReadTimeout:  time.Second,
		// WriteTimeout: time.Second,
	}

	mux.Handle("/block", http.HandlerFunc(blockIndefinitely))

	fmt.Println("ready to serve at", PORT)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
