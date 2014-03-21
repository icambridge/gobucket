package gobucket

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://bitbucket.org/api/"
	userAgent      = "gobucket/" + libraryVersion
)

func NewClient(r *Request) *Client {

	if r == nil {
		r = NewRequest(nil)
	}


	return &Client{request: r}

}

type Client struct {

	request   *Request
}


