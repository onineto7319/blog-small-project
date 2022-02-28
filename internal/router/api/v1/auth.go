package v1

import (
	"sync"

	"github.com/blog-small-project/pkg/errcode"

	"github.com/blog-small-project/internal/model"

	"github.com/blog-small-project/global"
	"github.com/blog-small-project/internal/service"
	"github.com/blog-small-project/pkg/app"
	"github.com/gin-gonic/gin"
)

type auth struct{}

var (
	authInstance *auth
)

func NewAuth() *auth {
	var once sync.Once
	if authInstance == nil {
		once.Do(func() { authInstance = &auth{} })
	}
	return authInstance
}

func (a auth) GetAuth(c *gin.Context) {
	request := app.NewResponse(c)

	checkAuth := model.CheckAuthRequest{}
	err := c.BindJSON(&checkAuth)

	if err != nil {
		errResponse := errcode.InvalidParams.WithDetails(err.Error())
		request.ToErrorResponse(errResponse)
		return
	}

	srv := service.NewAuth(global.MysqlEngine)

	err = srv.CheckAuth(checkAuth)
	if err != nil {
		request.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToekn(checkAuth.AppKey, checkAuth.AppSecret)

	if err != nil {
		request.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	request.ToResponse(model.CheckAuthResponse{Token: token})
}
