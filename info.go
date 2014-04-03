package gobucket

type PlaceInfo struct {
	Commit     CommitInfo `json:"commit"`
	Repository Repository `json:"repository"`
	Branch     BranchName `json:"branch"`
}

type CommitInfo struct {
	Hash       string     `json:"hash"`
	Links      SelfLinks  `json:"links"`
	Repository Repository `json:"repository"`
	Branch     BranchName `json:"branch"`
}
