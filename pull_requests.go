package gobucket

import (
	"fmt"
)

type PullRequestsService struct {
	client *Client
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
