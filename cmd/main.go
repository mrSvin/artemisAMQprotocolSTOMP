package main

import (
	"artemisAMQ/internal"
	"time"
)

func main() {

	go internal.Listener()
	time.Sleep(1 * time.Second)
	internal.SendMessage()
	time.Sleep(3 * time.Second)
}
