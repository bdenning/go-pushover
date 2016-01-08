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
	{"$token$",
		"[invalid user token]",
		"$device$",
		"Error Status Without Error Response", "This test should fail",
		`{"user":"invalid","status":0,"request":"024e029a6569c0224c8e3a5510657ee8"}`,
		0},
}

// TestPushConnectionFailure tests what will happen if the pushover.net API cannot be contacted due to network connectivity problems.
func TestPushConnectionFailure(t *testing.T) {
	m := pushover.NewMessage("", "", "")
	m.URL = "http://example.net:1234" // Set a bogus URL
	_, err := m.Push("Test Title", "Test message contents")

	// Make sure the attempt to send always results in an error being returned.
	if err == nil {
		t.Fail()
	}

	// TODO(@bdenning) We should test here for specific error responses.
	// These will be either ErrHTTPStatus (if behind a proxy) or a connection failure if directly connected to the internet.
}

// TestPush runs through a number of test cases (testCases) and ensures that API responses are as expected.
func TestPush(t *testing.T) {
	for _, test := range testCases {
		// Create a fresh new message object for each test case
		m := pushover.NewMessage(test.Token, test.User, test.Device)

		// Replace the token, user and device values. Some of these replace statement are intended to fail.
		m.Token = strings.Replace(m.Token, "$token$", os.Getenv("PUSHOVER_TOKEN"), 1)
		m.User = strings.Replace(m.User, "$user$", os.Getenv("PUSHOVER_USER"), 1)
		m.Device = strings.Replace(m.Device, "$device$", os.Getenv("PUSHOVER_DEVICE"), 1)

		// If the PUSHOVER_USE_REAL_API environment variable isn't set, then use a mock http service running locally.
		if os.Getenv("PUSHOVER_USE_REAL_API") != "true" {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, test.ExpectedResponse)
			}))
			defer s.Close()

			m.URL = s.URL
		}

		// Send a message and check for errors
		resp, err := m.Push(test.Title, test.Message)
		if err != nil {
			if test.ExpectedStatus != 0 {
				t.Errorf("A test that should have failed \"%s\" has passed: %v", test.Title, err)
			}
		}

		// Just to double check, we make sure that the status code returned by the API is what we were expecting.
		if resp.Status != test.ExpectedStatus {
			t.Errorf("The \"%s\" test returned an unexpected status code: %v", test.Title, resp.Status)
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
