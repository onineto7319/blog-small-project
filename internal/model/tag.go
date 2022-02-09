package model

type Tag struct {
	*Common
	Title          string `json:"title"`
	Desc           string `json:"desc"`
	ConverImageUrl string `json:"conver_image_url"`
	Content        string `json:"content"`
	State          uint8  `json:"state"`
}
