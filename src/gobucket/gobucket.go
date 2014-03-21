package gobucket

import (
	"net/url"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://bitbucket.org/api/"
	userAgent      = "gobucket/" + libraryVersion
)

func NewClient(r *Request) *Client {

	if r == nil {
		r = NewRequest(nil)
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	return &Client{request: r, BaseURL: baseURL}

}

type Client struct {
	BaseURL  *url.URL

	request   *Request

}
