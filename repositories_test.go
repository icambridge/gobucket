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
