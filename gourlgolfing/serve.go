package main

import (
	"fmt"
	"net/http"
)

func serve(folder string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		payload, err := bundle(folder)
		if err != nil {
			http.Error(w, "Bundle error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(payload)
	})

	fmt.Printf("Serving dev testing files from '%s' at localhost:8080\n", folder)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return fmt.Errorf("Dev server stopped: %v", err)
	}

	return nil
}
