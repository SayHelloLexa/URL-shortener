package model

type Url struct {
	Id       int    `json:"id,omitempty"`
	LongUrl  string `json:"long_url,omitempty"`
	ShortUrl string `json:"short_url,omitempty"`
}
