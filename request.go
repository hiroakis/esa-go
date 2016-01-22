package esa

type PostData struct {
	PostContent PostContent `json:"post"`
}

type PostContent struct {
	Name             string           `json:"name"`
	BodyMd           string           `json:"body_md"`
	Tags             []string         `json:"tags"`
	Category         string           `json:"category"`
	Wip              bool             `json:"wip"`
	Message          string           `json:"message"`
	OriginalRevision OriginalRevision `json:"original_revision"`
}

type OriginalRevision struct {
	BodyMd string `json:"body_md"`
	Number int    `json:"number"`
	User   string `json:"user"`
}

type CommentData struct {
	CommentContent CommentContent `json:"comment"`
}

type CommentContent struct {
	BodyMd string `json:"body_md"`
}
