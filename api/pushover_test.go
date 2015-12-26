package pushover

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	goodResponse = `{"status":1,"request":"e460545a8b333d0da2f3602aff3133d6"}`
	badResponse  = `{"user":"invalid","errors":["user identifier is invalid"],"status":0,"request":"3c0b5952fba0ab27805159217077c177"}`
)

func TestGoodResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, goodResponse)
	}))
	defer ts.Close()

	m := newMessageTest("Token", "App", "Device", "Title", "Message", ts.URL)
	if err := m.Send(); err != nil {
		t.Error("Unable to send message: ", err)
	}
}

func TestBadResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, badResponse)
	}))
	defer ts.Close()

	m := newMessageTest("Token", "App", "Device", "Title", "Message", ts.URL)
	if err := m.Send(); err == nil {
		t.Error("Should have received an error but didn't")
	}
}
