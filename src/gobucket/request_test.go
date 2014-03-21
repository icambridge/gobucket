package gobucket

import (
	"testing"
)

func TestNewRequest(t *testing.T) {
	r := NewRequest(nil)

	if r.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, want %v", r.UserAgent, userAgent)
	}
}
