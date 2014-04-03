package gobucket

import (
	"encoding/json"
)

func GetHookData(payload []byte) (*Hook, error) {

	var h Hook

	err := json.Unmarshal(payload, &h)

	if err != nil {
		return nil, err
	}

	return &h, nil
}

type Hook struct {
	Repository HookRepository `json:"repository"`
	Truncated  bool           `json:"truncated"`
	Commits    []Commit       `json:"commits"`
	ConnonUrl  string         `json:"canon_url"`
	User       string         `json:"user"`
}

type Callback interface {
	Exec(h *Hook)
}

type HookObserver struct {
	callbacks []Callback
}

func (o *HookObserver) Add(c Callback) {
	o.callbacks = append(o.callbacks, c)
}

func (o *HookObserver) Process(h *Hook) {
	for _, c := range o.callbacks {
		c.Exec(h)
	}
}
