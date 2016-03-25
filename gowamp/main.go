package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gngeorgiev/gowamp"
)

var (
	realm string
	port  int
	debug bool
)

func init() {
	flag.StringVar(&realm, "realm", "realm1", "realm name")
	flag.IntVar(&port, "port", 8000, "port to run on")
	flag.BoolVar(&debug, "debug", false, "enable debug logging")
}

func main() {
	flag.Parse()
	if debug {
		gowamp.Debug()
	}
	s := gowamp.NewBasicWebsocketServer(realm)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	go func() {
		<-shutdown
		s.Close()
		log.Println("shutting down server...")
		time.Sleep(time.Second)
		os.Exit(1)
	}()

	server := &http.Server{
		Handler: s,
		Addr:    ":8000",
	}
	log.Printf("gowamp server starting on port %d...", port)
	log.Fatal(server.ListenAndServe())
}
