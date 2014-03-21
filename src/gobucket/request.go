package gobucket

import (
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
	"strings"
)

func NewRequest(httpClient *http.Client) *Request {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	r := Request{client: httpClient, UserAgent: userAgent, BaseURL: baseURL}

	return &r

}

type Request struct {

	BaseURL  *url.URL

	client *http.Client

	UserAgent string
}

func (r *Request) NewHttpRequest(method string, urlString string, body interface{}) (*http.Request, error) {

	rel, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	u := strings.TrimRight(r.BaseURL.String(), "/")

	u += rel.String()

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}


	req, err := http.NewRequest(method, u, buf)

	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", r.UserAgent)

	return req, nil
}
