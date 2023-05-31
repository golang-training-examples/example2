// Package server provides a simple HTTP server
package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// StatusResponse is a response for /status endpoint
type StatusResponse struct {
	// Hostname is a hostname of the server
	Hostname string `json:"hostname"`

	// StartedTimestamp is a timestamp when server was started
	StartedTimestamp int64 `json:"started_timestamp"`

	// Uptime is a number of seconds since server was started
	Uptime int64 `json:"uptime"`
}

// Server starts HTTP server
func Server() {
	HOSTNAME, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	STARTED := time.Now().Unix()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	http.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(StatusResponse{
			Hostname:         HOSTNAME,
			StartedTimestamp: STARTED,
			Uptime:           time.Now().Unix() - STARTED,
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	fmt.Println("Listen on 0.0.0.0:8000, see: http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
