package pushover_test

import "github.com/bdenning/pushover"

var testCases = []struct {
	Token            string
	User             string
	Device           string
	Title            string
	Message          string
	URL              string
	ExpectedResponse string
	ExpectedStatus   int
}{
	{"$token$",
		"$user$",
		"$device$",
		"Ok Test",
		"This test should succeed",
		pushover.PushoverURL,
		`{"status":1,"request":"5e4a7a331ba4e45f3eb26cf447d61466"}`,
		1},
	{"[invalid device token]",
		"$user$",
		"$device$",
		"Invalid Device Token",
		"This test should fail",
		pushover.PushoverURL,
		`{"token":"invalid","errors":["application token is invalid"],"status":0,"request":"2eb28a69b6d9d67e5a937829954a8273"}`,
		0},
	{"$token$",
		"[invalid user token]",
		"$device$",
		"Invalid User Token",
		"This test should fail",
		pushover.PushoverURL,
		`{"user":"invalid","errors":["user identifier is not a valid user, group, or subscribed user key"],"status":0,"request":"024e029a6569c0224c8e3a5510657ee8"}`,
		0},
	{"$token$",
		"[invalid user token]",
		"$device$",
		"Error Status Without Error Response",
		"This test should fail",
		pushover.PushoverURL,
		`{"user":"invalid","status":0,"request":"024e029a6569c0224c8e3a5510657ee8"}`,
		0},
	{"$token$",
		"$user$",
		"$device$",
		"Connection Failure Test",
		"This test should fail",
		"http://localhost:5555",
		``,
		0},
	{"$token$",
		"$user$",
		"$device$",
		"External Connection Failure Test",
		"This test should fail",
		"http://example.net:5555",
		``,
		0},
}
