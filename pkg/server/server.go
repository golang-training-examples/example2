package server

import (
	"fmt"
	"net/http"
)

func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	http.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	fmt.Println("Listen on 0.0.0.0:8000, see: http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
