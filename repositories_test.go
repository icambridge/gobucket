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
