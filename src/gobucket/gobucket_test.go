package gobucket

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"net/url"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the GitHub client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// gobucket client configured to use test server
	client = NewClient(nil)
	url, _ := url.Parse(server.URL)
	client.request.BaseURL = url
}

// tearDown closes the test HTTP server.
func tearDown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	setup()
	defer tearDown()

	c := NewClient(nil)

	if c.request == nil {
		t.Error("Request shouldn't be empty")
	}
}
