package cli

import (
	conn "bring-your-test/pkgs/connection"
	consts "bring-your-test/pkgs/consts"
	msg "bring-your-test/pkgs/models"
	"log"
	"sync"
	"time"
)

type Client struct {
	ReconnectChannel chan string
	MessagesReceived []msg.Message
	MessagesSent     []msg.Message
}

func Create() *Client {
	return &Client{
		MessagesReceived: []msg.Message{},
		MessagesSent:     []msg.Message{},
		ReconnectChannel: make(chan string),
	}
}

func (client *Client) Routine(serverPort string, message string) {
	// Connect to the TCP server
	handler, err := conn.Create(":" + serverPort)
	if err != nil {
		log.Println("--- failed to create ---")
		client.ReconnectChannel <- ""
		return
	}

	// Close connection after 60s
	timer := time.NewTimer(time.Second * consts.SessionTime)
	defer timer.Stop()

	// Buffered channel to store received data from the client
	messageChannel := make(chan msg.Message, 1)

	err = handler.Send(msg.Message{
		UUID:    consts.MockUUID,
		Prefix:  "message",
		Message: message,
	})
	if err != nil {
		handler.Close()
		client.ReconnectChannel <- ""
		return
	}

	go func() {
		for {
			receivedMessage, err := handler.Receive()
			if err != nil {
				// log.Println("-- failed to receive --", err)
				break
			}

			// Update message received history
			client.MessagesReceived = append(client.MessagesReceived, receivedMessage)

			if receivedMessage.Prefix == "close" {
				break
			}
			messageChannel <- receivedMessage
			if receivedMessage.Prefix == "ok" && receivedMessage.UUID == consts.MockUUID {
				break
			}
		}
		close(messageChannel)
		// log.Println("-- receive goroutine finished --")
	}()

	for {
		select {
		case receivedMessage := <-messageChannel:
			if receivedMessage.Prefix == "ok" && receivedMessage.UUID == consts.MockUUID {
				// log.Println("--- close by OX condition ---")
				handler.Close()
				client.ReconnectChannel <- ""
				return
			}
			if receivedMessage.Prefix == "message" {
				receivedMessage.Prefix = "ok"
			}
			handler.Send(receivedMessage)
			client.MessagesSent = append(client.MessagesSent, receivedMessage)

		case <-timer.C:
			// log.Println("--- close by timer ---")
			handler.Send(msg.Message{
				Prefix:  "close",
				UUID:    consts.MockUUID,
				Message: "",
			})
			handler.Close()
			client.ReconnectChannel <- message
			return
		}
	}
}

func (client *Client) Run(serverPort string, waitGroup *sync.WaitGroup) {
	go client.Routine(serverPort, consts.RandomMessage())
	for message := range client.ReconnectChannel {
		if message == "" {
			break
		}
		go client.Routine(serverPort, message)
	}
	waitGroup.Done()
}