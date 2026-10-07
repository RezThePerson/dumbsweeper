package dev

import (
	"log"
	"net/http"
	"os"
)

func Serve(folder string) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		payload, err := bundle(folder)
		if err != nil {
			http.Error(w, "Bundle error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(payload)
	})

	log.Printf("Serving dev testing files from '%s' at http://localhost:%s\n", folder, port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Dev server stopped: %v\n", err)
	}
}
