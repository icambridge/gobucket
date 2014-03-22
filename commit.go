package gobucket

type Commit struct {
	Node string `json:"node"`
	Files []File `json:"files"`
	RawAuthor string `json:"raw_author"`
	UtcTimestamp string `json:"utctimestamp"`
	Author string `json:"author"`
	Timestamp string `json:"timestamp"`
	RawNode string `json:"raw_node"`
	Parents []string `json:"parents"`
	Branch string `json:"branch"`
	Message string `json:"message"`
	Size int `json:"size"`
}

type File struct {
	Type string `json:"type"`
	File string `json:"file"`
}
