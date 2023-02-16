package internal

import (
	"fmt"
	"github.com/go-stomp/stomp"
)

func Listener() error {
	options := []func(*stomp.Conn) error{
		stomp.ConnOpt.Login("artemis", "simetraehcapa"),
	}
	conn, err := stomp.Dial("tcp", "localhost:61616", options...)
	if err != nil {
		fmt.Println("Failed to connect to ActiveMQ Artemis:", err)
		return err
	}

	sub, err := conn.Subscribe("/queue/test", stomp.AckAuto)
	if err != nil {
		fmt.Println("Failed to subscribe to queue:", err)
		return err
	}

	for {
		msg, err := sub.Read()
		if err != nil {
			fmt.Println("Failed to read message:", err)
			return err
		}
		messageID := msg.Header.Get("message-id")
		fmt.Println("Received message:", string(msg.Body), " Received message ID:", messageID)
	}

	defer conn.Disconnect()

	return nil
}
