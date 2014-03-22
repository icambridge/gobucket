package gobucket

type Owner struct {
	Username string `json:"usernane"`
	DisplayName string `json:"display_name"`
	Links       OwnerLinks `json:"links"`
}

type OwnerLinks struct {
	Self Link `json:"self"`
	Avatar Link `json:"avatar"`
}
