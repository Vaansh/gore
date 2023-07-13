package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// Legacy: move to cloud storage
func main() {
	port := "8080"
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %s\n", err)
		os.Exit(1)
	}

	dir := filepath.Join(workingDir, "data")
	fileServer := http.FileServer(http.Dir(dir))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if filepath.Ext(r.URL.Path) == ".mp4" {
			fileServer.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	fmt.Printf("Serving MP4 files from directory '%s' on port %s\n", dir, port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
		os.Exit(1)
	}
}
