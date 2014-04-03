package gobucket

type Link struct {
	Href string `json:"href"`
}

type NamedLink struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

type SelfLinks struct {
	Self   Link `json:"self"`
	Avatar Link `json:"avatar"`
}
