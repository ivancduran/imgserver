package models

type Response struct {
	Response  bool   `json:"response"`
	Code      string `json:"code,omitempty"`
	Format    string `json:"format,omitempty"`
	Extension string `json:"extension,omitempty"`
}
