package gobucket

type Repository struct {
	Website     string `json:"website"`
	Fork        bool   `json:"fork"`
	Name        string `json:"name"`
	Scm         string `json:"scm"`
	Owner       string `json:"owner"`
	AbsoluteUrl string `json:"absolute_url"`
	Slug        string `json:"slug"`
	Private     bool   `json:"is_private"`
	Request     *Request
}
