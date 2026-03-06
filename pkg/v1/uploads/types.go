package uploads

// Image represents an uploaded asset returned by uploads endpoints.
type Image struct {
	Id         string `json:"id"`
	FileName   string `json:"file_name"`
	Height     int    `json:"height"`
	Width      int    `json:"width"`
	Size       int    `json:"size"`
	MimeType   string `json:"mime_type"`
	PreviewUrl string `json:"preview_url"`
	UploadTime string `json:"upload_time"`
}

// ImageUpload represents an upload request body for /v1/uploads/images.json.
type ImageUpload struct {
	Filename string `json:"file_name"`
	Contents []byte `json:"contents,omitempty"`
	Url      string `json:"url,omitempty"`
}
