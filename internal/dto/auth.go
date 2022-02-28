package dto

type Auth struct {
	Common
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}
