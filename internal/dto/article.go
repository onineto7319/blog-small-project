package dto

type Article struct {
	*Common
	Name  string `json:"name"`
	State int8   `json:"state"`
}
