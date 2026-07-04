package main

import (
	"fmt"
	sr "go-web-server/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	//1.initialize a custom request router(multiplexer)
	mux := http.NewServeMux()

	//2. register route paths and their corresponding handler functions
	mux.HandleFunc("/", sr.HomeHandler)
	mux.HandleFunc("/static/", sr.CssHandler)
	mux.HandleFunc("/api/data", sr.ApiHandler)
	mux.HandleFunc("/img/", sr.ServeImages)
	//3. configure explicit server parameters for safety and timeouts
	server := &http.Server{
		Addr:         ":8080",          // listen on port 8080
		Handler:      mux,              //inject our router custom configuration
		ReadTimeout:  10 * time.Second, //// Max time reading request headers/body
		WriteTimeout: 10 * time.Second, //Max time writing response
		IdleTimeout:  30 * time.Second, // Max time to keep connection open idle
	}
	//4. Fire up the server
	fmt.Println("Server successfully setarted on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
