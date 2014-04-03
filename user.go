package gobucket

type User struct {
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Links       SelfLinks `json:"links"`
}
