package gobucket

import (
	"fmt"
	"strings"
)

type PullRequestsService struct {
	client *Client
}

func (s *PullRequestsService) Approve(owner string, repo string, id int) error {
	url := fmt.Sprintf("/2.0/repositories/%s/%s/pullrequests/%d/approve", strings.ToLower(owner), strings.ToLower(repo), id)

	req, err := s.client.NewRequest("POST", url, nil)

	if err != nil {
		return err
	}

	err = s.client.Do(req, nil)

	if err != nil {
		return nil
	}

	return nil
}

func (s *PullRequestsService) Unapprove(owner string, repo string, id int) error {
	url := fmt.Sprintf("/2.0/repositories/%s/%s/pullrequests/%d/approve", strings.ToLower(owner), strings.ToLower(repo), id)

	req, err := s.client.NewRequest("DELETE", url, nil)

	if err != nil {
		return err
	}

	err = s.client.Do(req, nil)

	if err != nil {
		return nil
	}

	return nil
}

func (s *PullRequestsService) GetBranch(owner string, repo string, branch string) (*PullRequest, error) {

	pullRequests, err := s.GetAll(owner, repo)

	if err != nil {
		return nil, err
	}

	for _, pr := range pullRequests {
		if pr.Source.Branch.Name == branch {
			return pr, nil
		}
	}

	// Not found. Not really an error?
	return nil, nil
}

func (s *PullRequestsService) GetById(owner string, repo string, id int) (*PullRequest, error) {
	url := fmt.Sprintf("/2.0/repositories/%s/%s/pullrequests/%d", owner, repo, id)
	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	var pullRequest PullRequest
	err = s.client.Do(req, &pullRequest)

	if err != nil {
		return nil, err
	}

	return &pullRequest, nil
}

func (s *PullRequestsService) GetAll(owner string, repo string) ([]*PullRequest, error) {
	output := []*PullRequest{}
	count := 0
	page := 1

	for {
		url := fmt.Sprintf("/2.0/repositories/%s/%s/pullrequests/?page=%d", owner, repo, page)

		req, err := s.client.NewRequest("GET", url, nil)

		if err != nil {
			return nil, err
		}

		var pullRequests PullRequestList
		err = s.client.Do(req, &pullRequests)

		if err != nil {
			return nil, err
		}

		repoValues := len(pullRequests.Values)

		if repoValues <= 0 {
			break;
		}

		output = append(output, pullRequests.Values...)
		count += repoValues

		if count >= pullRequests.Size {
			break;
		}

		page++
	}
	return output, nil
}

func (s *PullRequestsService) Merge(owner string, repo string, id int) error {

	url := fmt.Sprintf("/2.0/repositories/%s/%s/pullrequests/%d/merge", strings.ToLower(owner), strings.ToLower(repo), id)
	req, err := s.client.NewRequest("POST", url, nil)

	if err != nil {
		return err
	}

	err = s.client.Do(req, nil)

	if err != nil {
		return err
	}

	return nil
}

type PullRequestList struct {
	PageLen int            `json:"pagelen"`
	Values  []*PullRequest `json:"values"`
	Page    int            `json:"page"`
	Size    int            `json:"size"`
}

type PullRequest struct {
	Description       string           `json:"description"`
	Links             PullRequestLinks `json:"links"`
	Author            User             `json:"author"`
	CloseSourceBranch bool             `json:"close_source_branch"`
	Destination       PlaceInfo        `json:"destination"`
	Source            PlaceInfo        `json:"source"`
	Title             string           `json:"title"`
	State             string           `json:"state"`
	CreatedOn         string           `json:"created_on"`
	UpdatedOn         string           `json:"updated_on"`
	Reviewers         []User           `json:"reviewers"`
	Participants	  []Participant    `json:"participants"`
	Id                int              `json:"id"`
}


type PullRequestLinks struct {
	Decline  Link `json:"decline"`
	Commits  Link `json:"commits"`
	Self     Link `json:"self"`
	Comments Link `json:"comments"`
	Merge    Link `json:"merge"`
	Html     Link `json:"html"`
	Activity Link `json:"activity"`
	Diff     Link `json:"diff"`
	Approve  Link `json:"approve"`
}

type Participant struct {
	Role string `json:"role"`
	User User `json:"user"`
	Approved bool `json:"approved"`
}

func (pr *PullRequest) GetApprovals() []User {
	approvals := []User{}
	for _, p := range pr.Participants {
		if p.Approved == true {
			approvals = append(approvals, p.User)
		}
	}
	return approvals
}

func (pr *PullRequest) GetOwner() string {

	parts := strings.Split(pr.Destination.Repository.FullName ,"/")

	if len(parts) < 1 || parts[0] == "" {
		return "unknown owner"
	}

	return parts[0]
}

func (pr *PullRequest) GetRepoName() string {

	parts := strings.Split(pr.Destination.Repository.FullName ,"/")

	if len(parts) < 2 {
		return "unknown repo"
	}

	return parts[1]
}

type PullRequestMerge struct {
	Title string `json:"title"`
	Description string `json:"title"`
	Source PlaceInfo `json:"source"`
	Destination PlaceInfo `json:"destination"`
}
