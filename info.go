package gobucket


type PlaceInfo struct {
	Commit     CommitInfo `json:"commit"`
	Repository Repository `json:"repository"`
	Branch     Branch     `json:"branch"`
}

type CommitInfo struct {
	Hash       string     `json:"hash"`
	Links      SelfLinks  `json:"links"`
	Repository Repository `json:"repository"`
	Branch     Branch     `json:"branch"`
}

type Branch struct {
	Name string `json:"name"`
}
