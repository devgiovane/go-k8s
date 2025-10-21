package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := os.Getenv("NAME")
		fmt.Fprintf(w, "Hello, I'm %s.", name)
	})
	http.ListenAndServe(":80", nil)
}
