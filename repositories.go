package gobucket

import (
	"fmt"
)

type RepositoriesService struct {
	client *Client
}

type Repository struct {
	Website     string `json:"website"`
	Fork        bool   `json:"fork"`
	Name        string `json:"name"`
	Scm         string `json:"scm"`
	Owner       string `json:"owner"`
	AbsoluteUrl string `json:"absolute_url"`
	Slug        string `json:"slug"`
	Private     bool   `json:"is_private"`
}

func (rs *RepositoriesService) Get(user string, repoName string) (*Repository, error) {

	endPoint := fmt.Sprintf("/2.0/repositories/%s/%s", user, repoName)

	var repo Repository

	req, err := rs.client.NewRequest("GET", endPoint, nil)

	if err != nil {
		return nil, err
	}

	rs.client.Do(req, &repo)

	return &repo, nil

}
