package gobucket

import (
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
	"strings"
//	"reflect"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://bitbucket.org/api"
	userAgent      = "gobucket/" + libraryVersion
)

func NewClient(usernameStr string, passwordStr string) *Client {

	httpClient := http.DefaultClient


	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client: httpClient,
		UserAgent: userAgent,
		BaseURL: baseURL,
		username: usernameStr,
		password: passwordStr,
	}

	c.Repositories = &RepositoriesService{client: c}
	c.PullRequests = &PullRequestsService{client: c}
	return c
}

type Client struct {

	BaseURL  *url.URL

	client *http.Client

	UserAgent string

	username string
	password string

	Repositories *RepositoriesService
	PullRequests *PullRequestsService

}



func (c *Client) NewRequest(method string, urlString string, body interface{}) (*http.Request, error) {

	rel, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	u := strings.TrimRight(c.BaseURL.String(), "/") + rel.String()
	buf := new(bytes.Buffer)
	if body != nil {
		if str, ok := body.(string); ok {
			buf.WriteString(str)
		} else {
			err := json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}

	}


	req, err := http.NewRequest(method, u, buf)

	if err != nil {
		return nil, err
	}

	if c.username != "" {
		req.SetBasicAuth(c.username, c.password)
	}

	req.Header.Add("User-Agent", c.UserAgent)

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
