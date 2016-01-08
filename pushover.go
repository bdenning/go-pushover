// Package pushover provides methods for sending messages using the http://pushover.net API.
package pushover

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// PushoverURL is the API endpoint that will be used for sending all messages.
	PushoverURL = "https://api.pushover.net/1/messages.json"
	// StatusSuccess is the expected status code when a message has been succesfully sent.
	StatusSuccess = 1
)

// Message contains all the required settings for sending messages via the pushover.net API
type Message struct {
	Token  string
	User   string
	Device string
	URL    string
}

// Response contains the JSON response returned by the pushover.net API
type Response struct {
	Request string   `json:"request"`
	Status  int      `json:"status"`
	Errors  []string `json:"errors"`
}

// NewMessage returns a new Message with API token values and a recipient device configured.
func NewMessage(token string, user string, device string) *Message {
	return &Message{token, user, device, PushoverURL}
}

// Push sends a message via the pushover.net API and returns the json response
func (m *Message) Push(title string, message string) (r *Response, err error) {
	msg := url.Values{}
	msg.Set("token", m.Token)
	msg.Set("user", m.User)
	msg.Set("device", m.Device)
	msg.Set("title", title)
	msg.Set("message", message)

	// Initalise an empty response
	r = &Response{}

	// Send the message the the pushover.net API
	resp, err := http.PostForm(m.URL, msg)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	// Read the JSON response in to a []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}

	// Copy the response from pushover.net in to a pushover.Response struct
	if err := json.Unmarshal(body, r); err != nil {
		return r, err
	}

	// Check to see if pushover.net set the status to indicate an error without providing and explanation
	if r.Status != StatusSuccess {
		if len(r.Errors) < 1 {
			return r, ErrUnknown
		}

		// TODO(@bdenning) Looks like the API can actualy return an array. We should support this.
		return r, errors.New(r.Errors[0])
	}

	return r, nil
}
