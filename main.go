package main

import (
	"log"
	"net/http"
)

var logger *log.Logger = log.Default()

func startWebServer(started chan bool) {
	if err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./static"))); err != nil {
		log.Fatal("Error occurred starting webserver")
	}
	started <- true
}

func main() {
	logger.Println("Starting webserver")
	started := make(chan bool)
	go startWebServer(started)
	<-started
}
