package gobucket

import (
	"testing"
	"io/ioutil"
)

func TestNewRequest(t *testing.T) {
	r := NewRequest(nil)

	if r.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, want %v", r.UserAgent, userAgent)
	}

	if r.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, want %v", r.BaseURL.String(), defaultBaseURL)
	}
}

func TestRequestNewRequest(t *testing.T) {
	r := NewRequest(nil)

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := &Link{Href: "l"}, `{"href":"l"}`+"\n"
	req, _ := r.NewHttpRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, want %v", inURL, req.URL, outURL)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewHttpRequest(%v) Body = %v, want %v", inBody, string(body), outBody)
	}

	// test that default user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if r.UserAgent != userAgent {
		t.Errorf("NewHttpRequest() User-Agent = %v, want %v", userAgent, r.UserAgent)
	}
}
