package main

import (
	"log"
	"net/http"

	"github.com/gngeorgiev/gowamp"
)

func main() {
	gowamp.Debug()
	s := gowamp.NewBasicWebsocketServer("gowamp.examples")
	server := &http.Server{
		Handler: s,
		Addr:    ":8000",
	}
	log.Println("gowamp server starting on port 8000")
	log.Fatal(server.ListenAndServe())
}
