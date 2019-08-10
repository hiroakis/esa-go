package request

type PostData struct {
	Post Post `json:"post"`
}

type Post struct {
	Name             string           `json:"name"`
	BodyMd           string           `json:"body_md"`
	Tags             []string         `json:"tags"`
	Category         string           `json:"category"`
	Wip              bool             `json:"wip"`
	Message          string           `json:"message"`
	OriginalRevision OriginalRevision `json:"original_revision,omitempty"`
	TemplatePostId   int              `json:"template_post_id"`
}

type OriginalRevision struct {
	BodyMd string `json:"body_md"`
	Number int    `json:"number"`
	User   string `json:"user"`
}

type CommentData struct {
	Comment Comment `json:"comment"`
}

type Comment struct {
	BodyMd string `json:"body_md"`
}
