// Package pushover provides methods for sending messages using the http://pushover.net API.
package pushover

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const pushoverURL = "https://api.pushover.net/1/messages.json"

// Message implements an io.Writer that will send messages to pushover.net
type Message struct {
	Token  string
	User   string
	Device string
	Title  string
	URL    string
}

// NewMessage returns a new Message with API token values and a recipient device configured.
func NewMessage(token string, user string, device string, title string) *Message {
	return &Message{token, user, device, title, pushoverURL}
}

// Push sends a message via the pushover.net API and returns the json response
func (m *Message) Push(message string) (r string, err error) {
	msg := url.Values{}
	msg.Set("token", m.Token)
	msg.Set("user", m.User)
	msg.Set("device", m.Device)
	msg.Set("title", m.Title)
	msg.Set("message", message)

	resp, err := http.PostForm(m.URL, msg)
	if err != nil {
		return "", errors.New("Unable to POST request to " + m.URL)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Unable to read HTTP response from " + m.URL)
	}

	return strings.Trim(string(body), "\n"), nil
}

// Write implements the io.Writer interface for convienience
func (m *Message) Write(p []byte) (n int, err error) {
	_, err = m.Push(string(p))
	if err != nil {
		return len(p), err
	}

	return len(p), nil
}
