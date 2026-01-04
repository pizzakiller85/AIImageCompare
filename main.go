package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "good job you clicked")
	})

	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		text := r.FormValue("userInput")
		fmt.Println("Received message:", text)
		fmt.Fprintf(w, "You Sent: %s", text)
	})
	fmt.Println("started server")
	http.ListenAndServe(":8080", nil)
}
