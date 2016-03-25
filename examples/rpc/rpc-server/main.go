package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gngeorgiev/gowamp"
)

var client *gowamp.Client

func main() {
	gowamp.Debug()
	s := gowamp.NewBasicWebsocketServer("gowamp.examples")
	server := &http.Server{
		Handler: s,
		Addr:    ":8000",
	}
	client, _ = s.GetLocalClient("gowamp.examples", nil)
	if err := client.BasicRegister("alarm.set", alarmSet); err != nil {
		panic(err)
	}
	log.Println("gowamp server starting on port 8000")
	log.Fatal(server.ListenAndServe())
}

// takes one argument, the (integer) number of seconds to set the alarm for
func alarmSet(args []interface{}, kwargs map[string]interface{}) (result *gowamp.CallResult) {
	duration, ok := args[0].(float64)
	if !ok {
		return &gowamp.CallResult{Err: gowamp.URI("rpc-example.invalid-argument")}
	}
	go func() {
		time.Sleep(time.Duration(duration) * time.Second)
		client.Publish("alarm.ring", nil, nil)
	}()
	return &gowamp.CallResult{Args: []interface{}{"hello"}}
}
