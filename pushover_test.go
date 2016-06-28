package pushover_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bdenning/go-pushover"
)

// TestPush runs through a number of test cases (testCases) and ensures that API responses are as expected.
func TestPush(t *testing.T) {
	for _, test := range testCases {
		// Run tests that are intended to test network connectivity and other non-API failures against the real API
		if os.Getenv("PUSHOVER_USE_REAL_API") == "true" && test.URL != pushover.PushoverURL {
			t.Skip()
		}

		// Create a fresh new message object for each test case
		m := pushover.NewMessage(test.Token, test.User)

		// Replace the token, user and device values. Some of these replace statement are intended to fail.
		m.Token = strings.Replace(m.Token, "$token$", os.Getenv("PUSHOVER_TOKEN"), 1)
		m.User = strings.Replace(m.User, "$user$", os.Getenv("PUSHOVER_USER"), 1)

		// If the PUSHOVER_USE_REAL_API environment variable isn't set, then use a mock http service running locally.
		if os.Getenv("PUSHOVER_USE_REAL_API") != "true" {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, test.ExpectedResponse)
			}))
			defer s.Close()

			m.URL = s.URL
		}

		// Send a message and check for errors
		resp, err := m.Push(test.Message)

		// Check for failures that did not result in Push() returning an error
		if err == nil && resp.Status != pushover.StatusSuccess {
			t.Errorf("A test that should have failed \"%s\" has passed: %v", test.Message, err)
		}

		// Check that the the status returned by the API is what we were expecting.
		if resp.Status != test.ExpectedStatus {
			t.Errorf("The \"%s\" test returned an unexpected status code: %v", test.Message, resp.Status)
		}
	}
}

func ExampleMessage_Push() {
	// You'll need to configure these by logging in to https://pushover.net.
	token := "KzGDORePKggMaC0QOYAMyEEuZJnyUi"
	user := "e9e1495ec75826de5983cd1abc8031"

	// Send a new message using the Push method.
	m := pushover.NewMessage(token, user)
	m.Push("Test message contents")
}
