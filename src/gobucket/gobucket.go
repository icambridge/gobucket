package gobucket

import (
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
	"strings"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://bitbucket.org/api"
	userAgent      = "gobucket/" + libraryVersion
)

func NewClient(httpClient *http.Client) *Client {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	c := Client{client: httpClient, UserAgent: userAgent, BaseURL: baseURL}

	return &c
}

type Client struct {

	BaseURL  *url.URL

	client *http.Client

	UserAgent string

}



func (r *Client) NewRequest(method string, urlString string, body interface{}) (*http.Request, error) {

	rel, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	u := strings.TrimRight(r.BaseURL.String(), "/") + rel.String()

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

func (c *Client) Do(req *http.Request, output interface {}) error {

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(output)

	return nil
}
