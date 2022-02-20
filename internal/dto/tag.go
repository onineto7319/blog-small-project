package dto

type Tag struct {
	Common
	Name  string `json:"name"`
	State int8   `json:"state"`
}
