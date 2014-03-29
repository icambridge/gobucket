package gobucket

import (
	"net/http"
	"testing"
	"fmt"
	"reflect"
)

func TestPullRequestsService_GetAll(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/2.0/repositories/batman/batcave/pullrequests/", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			if r.URL.Query().Get("page") == "1" {
				fmt.Fprint(w, `{
  "pagelen": 10,
  "next": "https://bitbucket.org/api/2.0/repositories/batman/?page=2",
  "values": [
    {
      "title": "Recognition / GenericBundle"
    },
    {
      "title": "Recognition / ResellerBundle"
    },
    {
      "title": "Recognition / CatalogueBundle"
    },
    {
      "title": "Recognition / Recognition Service"
    },
    {
      "title": "Recognition / BatchBundle"
    },
    {
      "title": "Recognition / Image Service"
    },
    {
      "title": "Recognition / Super User"
    },
    {
      "title": "Utility / Phingistrano"
    },
    {
      "title": "Internal / GIT Hooks"
    },
    {
      "title": "Internal / Bare Project Template"
    }
  ],
  "page": 1,
  "size": 11
}`)
			} else {
				fmt.Fprint(w, `{
  "pagelen": 1,
  "values": [
    {
      "title": "Recognition / ClientBundle"
    }
  ],
  "page": 2,
  "size": 11
}`)
			}
		})

	resp, err := client.PullRequests.GetAll("batman", "batcave")

	if err != nil {
		t.Errorf("Expected no err, got %v", err)
	}

	respLen := len(resp);

	if respLen != 11 {
		t.Errorf("Response length = %v, expected %v", respLen, 11)
	}

	expected := []*PullRequest{
		&PullRequest{Title: "Recognition / GenericBundle"},
		&PullRequest{Title: "Recognition / ResellerBundle"},
		&PullRequest{Title: "Recognition / CatalogueBundle"},
		&PullRequest{Title: "Recognition / Recognition Service"},
		&PullRequest{Title: "Recognition / BatchBundle"},
		&PullRequest{Title: "Recognition / Image Service"},
		&PullRequest{Title: "Recognition / Super User"},
		&PullRequest{Title: "Utility / Phingistrano"},
		&PullRequest{Title: "Internal / GIT Hooks"},
		&PullRequest{Title: "Internal / Bare Project Template"},
		&PullRequest{Title: "Recognition / ClientBundle"},
	}

	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Response body = %v, expected %v", resp, expected)
	}
}


func TestPullRequestsService_GetBranch(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/2.0/repositories/batman/batcave/pullrequests/", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			if r.URL.Query().Get("page") == "1" {
				fmt.Fprint(w, `{
  "pagelen": 10,
  "next": "https://bitbucket.org/api/2.0/repositories/batman/?page=2",
  "values": [
    {
      "title": "Recognition / GenericBundle",
      "branch": {"Name":"test"}
    },
    {
      "title": "Recognition / ResellerBundle",
      "branch": {"Name":"test"}
    },
    {
      "title": "Recognition / CatalogueBundle",
      "branch": {"Name":"test"}
    },
    {
      "title": "Recognition / Recognition Service",
      "branch": {"Name":"test"}
    },
    {
      "title": "Recognition / BatchBundle",
      "Source": {"branch": {"name":"thisone"}}
    },
    {
      "title": "Recognition / Image Service",
      "branch": {"name":"test"}
    },
    {
      "title": "Recognition / Super User",
      "branch": {"name":"test"}
    },
    {
      "title": "Utility / Phingistrano",
      "source": {"branch": {"name":"test"}}
    },
    {
      "title": "Internal / GIT Hooks",
      "source": {"branch": {"name":"test"}}
    },
    {
      "title": "Internal / Bare Project Template",
      "source": {"branch": {"name":"test"}}
    }
  ],
  "page": 1,
  "size": 11
}`)
			} else {
				fmt.Fprint(w, `{
  "pagelen": 1,
  "values": [
    {
      "title": "Recognition / ClientBundle",
      "source": {"branch": {"name":"test"}}
    }
  ],
  "page": 2,
  "size": 11
}`)
			}
		})

	resp, err := client.PullRequests.GetBranch("batman", "batcave", "thisone")

	if err != nil {
		t.Errorf("Expected no err, got %v", err)
	}

	expected := 		&PullRequest{
		Title: "Recognition / BatchBundle",
		Source: PlaceInfo{
			Branch: Branch{
				Name:"thisone",
			},
		},
	}
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Response body = %v, expected %v", resp, expected)
	}
}

func TestPullRequestsService_Approve(t *testing.T) {
	setUp()
	defer tearDown()

	apiHit := false

	mux.HandleFunc("/2.0/repositories/batman/batcave/pullrequests/123/approve", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			apiHit = true
			fmt.Fprint(w, `{"status": "success"}`)
		})

	err := client.PullRequests.Approve("batman", "batcave", 123)

	if err != nil {
		t.Errorf("Expected no err, got %v", err)
	}

	if apiHit != true {
		t.Errorf("Expected to hit api but didn't")
	}
}


func TestPullRequestsService_Approve_LowerCase(t *testing.T) {
	setUp()
	defer tearDown()

	apiHit := false

	mux.HandleFunc("/2.0/repositories/batman/batcave/pullrequests/123/approve", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			apiHit = true
			fmt.Fprint(w, `{"status": "success"}`)
		})

	err := client.PullRequests.Approve("Batman", "batCave", 123)

	if err != nil {
		t.Errorf("Expected no err, got %v", err)
	}

	if apiHit != true {
		t.Errorf("Expected to hit api but didn't")
	}
}


func TestPullRequestsService_Unapprove(t *testing.T) {
	setUp()
	defer tearDown()

	apiHit := false

	mux.HandleFunc("/2.0/repositories/batman/batcave/pullrequests/123/approve", func(w http.ResponseWriter, r *http.Request) {
			if m := "DELETE"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			apiHit = true
			fmt.Fprint(w, `{"status": "success"}`)
		})

	err := client.PullRequests.Unapprove("batman", "batcave", 123)

	if err != nil {
		t.Errorf("Expected no err, got %v", err)
	}

	if apiHit != true {
		t.Errorf("Expected to hit api but didn't")
	}
}


func TestPullRequest_GetApprovals(t *testing.T) {
	setUp()
	defer tearDown()

	expectedUser :=  User{DisplayName: "Iain Cambridge"}


	pr := PullRequest{}
	pr.Participants = []Participant{
		Participant{Role: "REVIEWER", User:expectedUser, Approved: true},
		Participant{Role: "REVIEWER", User:User{DisplayName: "Johnny"}, Approved: false},
	}

	found := pr.GetApprovals()
	expected := []User{expectedUser}
	if !reflect.DeepEqual(found, expected) {
		t.Errorf("Approvals = %v, expected %v", found, expected)
	}
}
