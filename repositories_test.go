package gobucket

import (
	"net/http"
	"testing"
	"fmt"
	"reflect"
)

func TestRepositoriesService_Get(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/2.0/repositories/batman/cave-system", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			fmt.Fprint(w, `{"Name":"Cave System"}`)
		})

	req, _ := client.Repositories.Get("batman", "cave-system")

	expected := &Repository{Name: "Cave System"}

	if !reflect.DeepEqual(req, expected) {
		t.Errorf("Response body = %v, expected %v", req, expected)
	}

}


func TestRepositoriesService_GetAll(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/2.0/repositories/batman", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			if r.URL.Query().Get("page") == "1" {
			fmt.Fprint(w, `{
  "pagelen": 10,
  "next": "https://bitbucket.org/api/2.0/repositories/batman?page=2",
  "values": [
    {
      "name": "Recognition / GenericBundle"
    },
    {
      "name": "Recognition / ResellerBundle"
    },
    {
      "name": "Recognition / CatalogueBundle"
    },
    {
      "name": "Recognition / Recognition Service"
    },
    {
      "name": "Recognition / BatchBundle"
    },
    {
      "name": "Recognition / Image Service"
    },
    {
      "name": "Recognition / Super User"
    },
    {
      "name": "Utility / Phingistrano"
    },
    {
      "name": "Internal / GIT Hooks"
    },
    {
      "name": "Internal / Bare Project Template"
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
      "name": "Recognition / ClientBundle"
    }
  ],
  "page": 2,
  "size": 11
}`)
			}
		})

	resp, err := client.Repositories.GetAll("batman")

	if err != nil {
		t.Errorf("Expected no err, got %v", err)
	}

	respLen := len(resp);

	if respLen != 11 {
		t.Errorf("Response length = %v, expected %v", respLen, 11)
	}

	expected := []*Repository{
		&Repository{Name: "Recognition / GenericBundle"},
		&Repository{Name: "Recognition / ResellerBundle"},
		&Repository{Name: "Recognition / CatalogueBundle"},
		&Repository{Name: "Recognition / Recognition Service"},
		&Repository{Name: "Recognition / BatchBundle"},
		&Repository{Name: "Recognition / Image Service"},
		&Repository{Name: "Recognition / Super User"},
		&Repository{Name: "Utility / Phingistrano"},
		&Repository{Name: "Internal / GIT Hooks"},
		&Repository{Name: "Internal / Bare Project Template"},
		&Repository{Name: "Recognition / ClientBundle"},
	}

	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Response body = %v, expected %v", resp, expected)
	}

}


func TestRepositoriesService_GetBranches(t *testing.T) {
	setUp()
	defer tearDown()
	apiHit := false
	json := `{
 "bug/RECOG-1302":  {
    "node": "66b2d310a3de",
    "files":  [
       {
        "type": "modified",
        "file": "src/Workstars/Recognition/ClientUserBundle/Tests/Functional/Controller/Recognition/MakeFunctionalTest.php"
      }
    ],
    "raw_author": "Matthew Rowland <matthew.rowland@workstars.com>",
    "utctimestamp": "2014-03-03 15:49:49+00:00",
    "author": "roloking1806",
    "timestamp": "2014-03-03 16:49:49",
    "raw_node": "66b2d310a3de8e1f0add3b2c2aadaa3b3f130fed",
    "parents":  [
      "6f14ddaed98f"
    ],
    "branch": "bug/RECOG-1302",
    "message": "RECOG-1302 - Add functional tests to ensure cross country/org unit recognitions save correctly.\n",
    "revision": null,
    "size": -1
  }
	}`
	mux.HandleFunc("/1.0/repositories/batman/cave-system/branches", func(w http.ResponseWriter, r *http.Request) {

			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			fmt.Fprint(w, json)
			apiHit = true
		},
	)
	branches, err := client.Repositories.GetBranches("batman", "cave-system")


	if apiHit == false {
		t.Error("Api wasn't hit")
	}

	if err != nil {
		t.Errorf("Expected no err, got %v", err)
	}

	if branches == nil {
		t.Error("Nil return")
	}

	value, ok := branches["bug/RECOG-1302"]

	if ok != true {
		t.Error("Branch 'bug/RECOG-1302' didn't exist",)
	}

	expected := &Branch{
		Node:  "66b2d310a3de",
		Files: []File{
			File{Type: "modified", File: "src/Workstars/Recognition/ClientUserBundle/Tests/Functional/Controller/Recognition/MakeFunctionalTest.php"},
		},
		RawAuthor: "Matthew Rowland <matthew.rowland@workstars.com>",
		UtcTimestamp: "2014-03-03 15:49:49+00:00",
		Author: "roloking1806",
		Timestamp: "2014-03-03 16:49:49",
		RawNode: "66b2d310a3de8e1f0add3b2c2aadaa3b3f130fed",
		Parents: []string{"6f14ddaed98f"},
		Branch: "bug/RECOG-1302",
		Message: "RECOG-1302 - Add functional tests to ensure cross country/org unit recognitions save correctly.\n",
		Size: -1,
	}

	if !reflect.DeepEqual(value, expected) {
		t.Errorf("Branch = %v, expected %v", value, expected)
	}
}
