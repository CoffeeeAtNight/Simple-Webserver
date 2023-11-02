package main

import (
	"log"
	"net/http"
)

var logger *log.Logger = log.Default()

func main() {
	logger.Println("Starting webserver")
	go func() {
		if err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./static"))); err != nil {
			log.Fatal("Error occurred starting webserver")
		}
	}()

	select {}
}
