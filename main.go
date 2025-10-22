package main

import (
	"net/http"
	"time"
)

func main() {
	startedAt := time.Now()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uptime := time.Since(startedAt)
		if uptime.Seconds() < 10 {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		// w.Write([]byte("ok"))
	})
	http.ListenAndServe(":80", nil)
}
