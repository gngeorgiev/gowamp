package main

import (
	"log"
	"net/http"

	"github.com/gngeorgiev/gowamp"
)

func main() {
	gowamp.Debug()
	s := gowamp.NewBasicWebsocketServer("gowamp.chat.realm")
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", s)
	log.Println("gowamp server starting on port 8000")
	log.Println("Hint: start clicking on the web page(s) you open to localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
