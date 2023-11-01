package main

import (
	"log"
	"net/http"
)

var logger *log.Logger = log.Default()

func startServer(handler http.Handler, done chan bool) {
	logger.Println("Starting webserver")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("Error occurred starting webserver")
	}
	done <- true
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	done := make(chan bool)
	go startServer(fileServer, done)
	<-done
}
