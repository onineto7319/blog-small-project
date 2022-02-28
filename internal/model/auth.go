package model

type CheckAuthRequest struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

type CheckAuthResponse struct {
	Token string `json:"token"`
}
