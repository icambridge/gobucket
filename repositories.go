package gobucket

import (
	"fmt"
)

type RepositoriesService struct {
	client *Client
}

func (s *RepositoriesService) Get(user string, repoName string) (*Repository, error) {

	endPoint := fmt.Sprintf("/2.0/repositories/%s/%s", user, repoName)

	var repo Repository

	req, err := s.client.NewRequest("GET", endPoint, nil)

	if err != nil {
		return nil, err
	}

	s.client.Do(req, &repo)

	return &repo, nil

}

func (s *RepositoriesService) GetAll(user string) ([]*Repository, error) {
	output := []*Repository{}
	count := 0
	page := 1

	for {
		url := fmt.Sprintf("/2.0/repositories/%s?page=%d", user, page)

		req, err := s.client.NewRequest("GET", url, nil)

		if err != nil {
			return nil, err
		}

		var repositories RepositoriesList

		err = s.client.Do(req, &repositories)

		if err != nil {
			return nil, err
		}

		repoValues := len(repositories.Values)

		if repoValues <= 0 {
			break;
		}

		output = append(output, repositories.Values...)
		count += repoValues

		if count >= repositories.Size {
			break;
		}

		page++
	}

	return output, nil
}

// What is returned by the API.
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

// What is sent by the web hooks.
type HookRepository struct {
	Website     string `json:"website"`
	Fork        bool   `json:"fork"`
	Name        string `json:"name"`
	Scm         string `json:"scm"`
	Owner       string `json:"owner"`
	AbsoluteUrl string `json:"absolute_url"`
	Slug        string `json:"slug"`
	Private     bool   `json:"is_private"`
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

type RepositoriesList struct {
	PageLen int          `json:"pagelen"`
	Values []*Repository  `json:"values"`
	Page    int          `json:"page"`
	Size    int          `json:"size"`
}
