package gobucket

import (
	"net/http"
)

func NewRequest(httpClient *http.Client) *Request {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	r := Request{client: httpClient, UserAgent: userAgent}

	return &r

}

type Request struct {
	client *http.Client

	UserAgent string
}
