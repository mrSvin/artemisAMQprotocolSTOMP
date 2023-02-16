package internal

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"log"
	"strconv"
)

func SendMessage() {
	for i := 1; i <= 5; i++ {
		message := "Hello ActiveMQ Artemis! - " + strconv.Itoa(i)
		sendMessageStomp(message)
	}

}

func sendMessageStomp(message string) {
	options := []func(*stomp.Conn) error{
		stomp.ConnOpt.Login("artemis", "simetraehcapa"),
	}

	conn, err := stomp.Dial("tcp", "localhost:61616", options...)
	if err != nil {
		log.Fatalf("Failed to connect to ActiveMQ Artemis: %s", err)
	}
	defer conn.Disconnect()

	err = conn.Send("/queue/test", "text/plain", []byte(message), nil)
	if err != nil {
		log.Fatalf("Failed to send me ssage: %s", err)
	}
	fmt.Println("Message sent successfully")
}
