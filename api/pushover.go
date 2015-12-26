package pushover

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiURL = "https://api.pushover.net/1/messages.json"

type Message struct {
	token   string
	user    string
	device  string
	title   string
	message string
	url     string
}

type Response struct {
	Status  int    `json:"status"`
	Request string `json:"request"`
}

func NewMessage(token string, user string, device string, title string, message string) *Message {
	return &Message{token, user, device, title, message, apiURL}
}

func newMessageTest(token string, user string, device string, title string, message string, url string) *Message {
	return &Message{token, user, device, title, message, url}
}

func (m *Message) Send() (err error) {
	resp, err := http.PostForm(m.url, m.encode())
	if err != nil {
		return errors.New("Unable to sent HTTP POST to the Pushover API")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Unable to read HTTP response from " + m.url)
	}

	r := new(Response)
	if err = json.Unmarshal(body, &r); err != nil {
		return errors.New("Unable to interpret the response from " + m.url)
	}

	if r.Status != 1 {
		return errors.New("Recieved response code indicating an error")
	}

	return nil
}

func (m *Message) encode() url.Values {
	v := url.Values{}
	v.Set("token", m.token)
	v.Set("user", m.user)
	v.Set("device", m.device)
	v.Set("title", m.title)
	v.Set("message", m.message)

	return v
}
