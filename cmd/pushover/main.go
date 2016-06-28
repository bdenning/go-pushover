package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/bdenning/go-pushover"
)

func getEnvSettings() (token string, user string) {
	token = os.Getenv("PUSHOVER_TOKEN")
	user = os.Getenv("PUSHOVER_USER")

	if token == "" {
		log.Fatal("Missing PUSHOVER_TOKEN environment variable")
	}

	if user == "" {
		log.Fatal("Missing PUSHOVER_USER environment variable")
	}

	return token, user
}

func main() {
	// Get the required configuration information
	token, user := getEnvSettings()

	// Create a new pushover message object with our settings
	m := pushover.NewMessage(token, user)

	// Read the message from stdin and send via pushover.net
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Error getting input from stdin")
	}

	// Send the message
	_, err = m.Push(string(stdin))
	if err != nil {
		log.Fatal("Error while sending message")
	}
}
