package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"fwdlio/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// Static files
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))

	// Routes
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/submit", handlers.SubmitHandler)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Server starting on port %s", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
