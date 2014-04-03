package gobucket

import (
	"testing"
)

func TestGetHookData(t *testing.T) {

	json := `{"repository": {"website": "", "fork": false, "name": "Recognition / Client Admin Application", "scm": "git", "owner": "workstars", "absolute_url": "/workstars/recognition-client-admin-application/", "slug": "recognition-client-admin-application", "is_private": true}, "truncated": false, "commits": [{"node": "82342c16fb54", "files": [{"type": "modified", "file": "README.md"}], "raw_author": "Iain Cambridge <iain.cambridge@workstars.com>", "utctimestamp": "2014-03-16 17:01:56+00:00", "author": "icambridge", "timestamp": "2014-03-16 18:01:56", "raw_node": "82342c16fb54cde537a0dc5b2a4773268bfd025b", "parents": ["f61e20d74f55"], "branch": "icambridge/readmemd-edited-online-with-bitbucket-1394989053935", "message": "README.md edited online with Bitbucket", "revision": null, "size": -1}], "canon_url": "https://bitbucket.org", "user": "icambridge"}`

	payload := []byte(json)

	h, err := GetHookData(payload)

	if err != nil {
		t.Errorf("No error expected, instead got %v", err)
	}

	if h.Repository.Website != "" {
		t.Errorf(".Repository.Website - Expected %v, got %v", "", h.Repository.Website)
	}
	if h.Repository.Name != "Recognition / Client Admin Application" {
		t.Errorf(".Repository.Name - Expected %v, got %v", "Recognition / Client Admin Application", h.Repository.Name)
	}
}
