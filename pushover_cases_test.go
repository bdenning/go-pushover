package pushover_test

import "github.com/bdenning/go-pushover"

var testCases = []struct {
	Token            string
	User             string
	Message          string
	URL              string
	ExpectedResponse string
	ExpectedStatus   int
}{
	{"$token$",
		"$user$",
		"Ok Test",
		pushover.PushoverURL,
		`{"status":1,"request":"5e4a7a331ba4e45f3eb26cf447d61466"}`,
		1},
	{"[invalid device token]",
		"$user$",
		"Invalid Device Token",
		pushover.PushoverURL,
		`{"token":"invalid","errors":["application token is invalid"],"status":0,"request":"2eb28a69b6d9d67e5a937829954a8273"}`,
		0},
	{"$token$",
		"[invalid user token]",
		"Invalid User Token",
		pushover.PushoverURL,
		`{"user":"invalid","errors":["user identifier is not a valid user, group, or subscribed user key"],"status":0,"request":"024e029a6569c0224c8e3a5510657ee8"}`,
		0},
	{"$token$",
		"[invalid user token]",
		"Error Status Without Error Response",
		pushover.PushoverURL,
		`{"user":"invalid","status":0,"request":"024e029a6569c0224c8e3a5510657ee8"}`,
		0},
	{"$token$",
		"$user$",
		"Connection Failure Test",
		"http://localhost:5555",
		``,
		0},
	{"$token$",
		"$user$",
		"External Connection Failure Test",
		"http://example.net:5555",
		``,
		0},
	{"$token$",
		"$user$",
		"Invalid HTTP protocol",
		"ftp://example.net",
		``,
		0},
}
