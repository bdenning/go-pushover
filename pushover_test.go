package pushover_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bdenning/pushover"
)

// TODO(@bowen) None of the following test cases are being used just yet.
var testCases = []struct {
	Title            string
	Message          string
	ExpectedResponse string
	ExpectedStatus   int
}{
	{"Testing", "Test message", `{"status":1,"request":"5e4a7a331ba4e45f3eb26cf447d61466"}`, 1},
	{"Testing", "Invalid user key", `{"user":"invalid","errors":["user identifier is not a valid user, group, or subscribed user key"],"status":0,"request":"024e029a6569c0224c8e3a5510657ee8"}`, 0},
	{"Testing", "Invalid device token", `{"token":"invalid","errors":["application token is invalid"],"status":0,"request":"2eb28a69b6d9d67e5a937829954a8273"}`, 0},
}

func mockPushoverServer(mockResponse string) *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockResponse)
	}))

	return s
}

func TestPush(t *testing.T) {
	for _, test := range testCases {
		s := mockPushoverServer(test.ExpectedResponse)
		defer s.Close()

		m := pushover.NewMessage("", "", "", test.Title)
		m.URL = s.URL

		resp, err := m.Push(test.Message)
		if err != nil {
			t.Error(err)
		}

		if resp != test.ExpectedResponse {
			t.Errorf("The response of %v is not equal to the expected response of %v", resp, test.ExpectedResponse)
		}
	}
}

func ExampleMessage_Push() {
	// Set API keys. You'll need to configure these by logging in to https://pushover.net.
	token := "KzGDORePKggMaC0QOYAMyEEuZJnyUi"
	user := "e9e1495ec75826de5983cd1abc8031"
	device := "test_device"
	title := "Example message"

	// Send a new message using the Push method.
	m := pushover.NewMessage(token, user, device, title)
	m.Push("Example message contents")
}

func TestWrite(t *testing.T) {
	for _, test := range testCases {
		s := mockPushoverServer(test.ExpectedResponse)
		defer s.Close()

		m := pushover.NewMessage("", "", "", test.Title)
		m.URL = s.URL

		byteCount, err := fmt.Fprintln(m, test.Message)
		if err != nil {
			t.Error(err)
		}

		if byteCount < 0 {
			t.Error("Response was smaller than expected")
		}
	}
}

func ExampleMessage_Write() {
	// Set API keys. You'll need to configure these by logging in to https://pushover.net.
	token := "KzGDORePKggMaC0QOYAMyEEuZJnyUi"
	user := "e9e1495ec75826de5983cd1abc8031"
	device := "test_device"
	title := "Example message"

	// Send a new message using the Write method.
	m := pushover.NewMessage(token, user, device, title)
	fmt.Fprintf(m, "Example message contents")

	// Here's another example using an io.Reader object and io.Copy
	io.Copy(m, strings.NewReader("Another example"))
}
