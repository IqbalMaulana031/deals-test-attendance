package resource

// UploadFile is a handler for upload file
type UploadFile struct {
	BaseURL string   `json:"base_url"`
	Path    []string `json:"file"`
}
