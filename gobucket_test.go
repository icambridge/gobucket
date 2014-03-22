package gobucket

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"io/ioutil"
	"testing"
	"fmt"
	"reflect"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the GitHub client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setUp() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// src.gobucket client configured to use test server
	client = NewClient("", "")
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

// tearDown closes the test HTTP server.
func tearDown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	setUp()
	defer tearDown()

	c := NewClient("batman", "alfredletmein")

	if c.username != "batman" {
		t.Errorf("NewClient username = %v, expected %v", c.username, "batman")
	}

	if c.password != "alfredletmein" {
		t.Errorf("NewClient password = %v, expected %v", c.password, "alfredletmein")
	}

	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, expected %v", c.UserAgent, userAgent)
	}

	if c.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, expected %v", c.BaseURL.String(), defaultBaseURL)
	}
}

func TestClientNewRequest(t *testing.T) {
	c := NewClient("", "")

	inURL, outURL := "/foo", defaultBaseURL+"/foo"
	inBody, outBody := &Link{Href: "l"}, `{"href":"l"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v) Body = %v, expected %v", inBody, string(body), outBody)
	}

	// test that default user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", userAgent, c.UserAgent)
	}
}

func TestDo_GET(t *testing.T) {
	setUp()
	defer tearDown()

	type Foo struct {
		Bar string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			fmt.Fprint(w, `{"Bar":"drink"}`)
		})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(Foo)
	client.Do(req, body)

	expected := &Foo{"drink"}

	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Response body = %v, expected %v", body, expected)
	}

}

func TestDo_POST(t *testing.T) {
	setUp()
	defer tearDown()

	type Foo struct {
		Bar string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			fmt.Fprint(w, `{"Bar":"drink"}`)
		})

	req, _ := client.NewRequest("POST", "/", nil)
	body := new(Foo)
	client.Do(req, body)

	expected := &Foo{"drink"}

	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Response body = %v, expected %v", body, expected)
	}

}
