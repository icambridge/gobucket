package gobucket

import (
	"fmt"
)

type RepositoriesService struct {
	client *Client
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

type Repository struct {
	Scm         string          `json:"scm"`
	HasWiki     bool            `json:"has_wiki"`
	Description string          `json:"description"`
	Links       RepositoryLinks `json:"links"`
	ForkPolicy  string          `json:"fork_policy"`
	Language    string          `json:"language"`
	CreatedOn   string          `json:"created_on"`
	FullName    string          `json:"full_name"`
	HasIssues   bool            `json:"has_isses"`
	Owner       Owner           `json:"owner"`
	UpdatedOn   string          `json:"updated_on"`
	Size        int             `json:"size"`
	IsPrivate   bool            `json:"is_private"`
	Name        string          `json:"name"`
}

type RepositoryLinks struct {
	Watchers Link `json:"watchers"`
	Commits Link `json:"commits"`
	Self Link `json:"self"`
	Html Link `json:"html"`
	Avatar Link `json:"avatar"`
	Forks Link `json:"fork"`
	Clone []NamedLink `json:"clone"`
	PullRequests Link `json:"pullrequests"`
}
