package main

import (
	"calculator/handlers"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("starting server")
	//1.initialize a custom request router(multiplexer)
	mux := http.NewServeMux()

	//2. register route paths and their corresponding handler functions
	mux.HandleFunc("/", handlers.Home)

	//3. configure explicit server parameters for safety and timeouts
	server := &http.Server{
		//first the port we are gonna listen to wich is 8080
		Addr: ":8080",
		//inject our router custom configuration
		Handler:      mux,
		ReadTimeout:  10 * time.Second, //// Max time reading request headers/body
		WriteTimeout: 10 * time.Second, //Max time writing response
		IdleTimeout:  30 * time.Second, // Max time to keep connection open idle
	}
	//4. Fire up the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}

}
