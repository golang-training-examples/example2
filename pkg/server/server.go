package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type StatusResponse struct {
	Hostname         string `json:"hostname"`
	StartedTimestamp int64  `json:"started_timestamp"`
	Uptime           int64  `json:"uptime"`
}

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
