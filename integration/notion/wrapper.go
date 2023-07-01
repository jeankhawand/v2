package notion

type Data struct {
	Children []Block `json:"children"`
}

type Block struct {
	Object   string   `json:"object"`
	Type     string   `json:"type"`
	Bookmark Bookmark `json:"bookmark"`
}

type Bookmark struct {
	Caption []interface{} `json:"caption"` // Assuming the "caption" field can have different types
	URL     string        `json:"url"`
}
