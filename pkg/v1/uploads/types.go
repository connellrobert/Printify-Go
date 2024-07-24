package uploads

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

type ImageUpload struct {
	Filename string `json:"filename"`
	Contents []byte `json:"contents,omitempty"`
	Url      string `json:"url,omitempty"`
}
