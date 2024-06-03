package main

import (
	cli "bringyour-test/client"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	// Get the values of the environment variables
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Println("Failed to load enviroment.")
		return
	}
	clientCount, err := strconv.Atoi(os.Getenv("CLIENT_COUNT"))
	if err != nil {
		log.Println("Failed to load enviroment.")
		return
	}
	sessionTime, err := strconv.Atoi(os.Getenv("SESSION_TIME"))
	if err != nil {
		log.Println("Failed to load enviroment.")
		return
	}

	// Create a WaitGroup
	var waitGroup sync.WaitGroup
	waitGroup.Add(clientCount)

	// Initialize
	clients := []*cli.Client{}

	// Create clients
	for i := 0; i < clientCount; i++ {
		clients = append(clients, cli.Create())
	}

	// Run clients
	randCount := rand.Intn(100)
	for index, client := range clients {
		go client.Run(serverPort, &waitGroup, sessionTime)
		if index == randCount {
			// randTime := rand.Intn(10) + 10
			time.Sleep(time.Second * 10)
		}
	}

	// Wait for all clients finished
	waitGroup.Wait()
	log.Println("All clients finished.")
}
