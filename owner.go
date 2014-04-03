package gobucket

type Owner struct {
	Username    string    `json:"usernane"`
	DisplayName string    `json:"display_name"`
	Links       SelfLinks `json:"links"`
}
