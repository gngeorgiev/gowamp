package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/howeyc/gopass"
	"github.com/gngeorgiev/gowamp"
)

var password []byte

func exampleAuthFunc(hello map[string]interface{}, c map[string]interface{}) (string, map[string]interface{}, error) {
	challenge, ok := c["challenge"].(string)
	if !ok {
		log.Fatal("no challenge data recieved")
	}
	mac := hmac.New(sha256.New, password)
	mac.Write([]byte(challenge))
	signature := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature), nil, nil
}

func main() {
	gowamp.Debug()
	fmt.Println("Hint: the password is 'password'")
	fmt.Print("Password: ")
	password = gopass.GetPasswd()

	c, err := gowamp.NewWebsocketClient(gowamp.JSON, "ws://localhost:8000/ws")
	if err != nil {
		log.Fatal(err)
	}
	c.Auth = map[string]gowamp.AuthFunc{"example-auth": exampleAuthFunc}
	_, err = c.JoinRealm("gowamp.examples", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected with auth")
	time.Sleep(3 * time.Second)
	fmt.Println("Disconnecting")
	c.Close()
}
