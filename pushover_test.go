package pushover_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bdenning/pushover"
)

var testCases = []struct {
	Token            string
	User             string
	Device           string
	Title            string
	Message          string
	ExpectedResponse string
	ExpectedStatus   int
}{
	{"$token$",
		"$user$",
		"$device$",
		"Ok Test", "This test should succeed",
		`{"status":1,"request":"5e4a7a331ba4e45f3eb26cf447d61466"}`,
		1},
	{"[invalid device token]",
		"$user$",
		"$device$",
		"Invalid Device Token", "This test should fail",
		`{"token":"invalid","errors":["application token is invalid"],"status":0,"request":"2eb28a69b6d9d67e5a937829954a8273"}`,
		0},
	{"$token$",
		"[invalid user token]",
		"$device$",
		"Invalid User Token", "This test should fail",
		`{"user":"invalid","errors":["user identifier is not a valid user, group, or subscribed user key"],"status":0,"request":"024e029a6569c0224c8e3a5510657ee8"}`,
		0},
}

func TestPush(t *testing.T) {
	for _, test := range testCases {
		m := pushover.NewMessage(test.Token, test.User, test.Device)

		m.Token = strings.Replace(m.Token, "$token$", os.Getenv("PUSHOVER_TOKEN"), 1)
		m.User = strings.Replace(m.User, "$user$", os.Getenv("PUSHOVER_USER"), 1)
		m.Device = strings.Replace(m.Device, "$device$", os.Getenv("PUSHOVER_DEVICE"), 1)

		if os.Getenv("PUSHOVER_USE_REAL_API") != "true" {
			// Run tests against a mock HTTP service
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, test.ExpectedResponse)
			}))
			defer s.Close()

			m.URL = s.URL
		}

		resp, err := m.Push(test.Title, test.Message)
		if err != nil {
			t.Error(err)
		}

		if resp.Status != test.ExpectedStatus {
			t.Errorf("The test message \"%s\" returned an unexpected status code of %d: %v\n", test.Title, resp.Status, resp.Errors[0])
		}
	}
}

func ExampleMessage_Push() {
	// You'll need to configure these by logging in to https://pushover.net.
	token := "KzGDORePKggMaC0QOYAMyEEuZJnyUi"
	user := "e9e1495ec75826de5983cd1abc8031"
	device := "test_device"

	// Send a new message using the Push method.
	m := pushover.NewMessage(token, user, device)
	m.Push("Test Title", "Test message contents")
}
