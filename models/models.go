package models

type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Error ErrorInfo `json:"error"`
}

type Data struct {
	Gzip    int    `json:"gzip"`
	Size    int    `json:"size"`
	Version string `json:"version"`
	Name    string `json:"name"`
	Error   *Error `json:"error"`
}
