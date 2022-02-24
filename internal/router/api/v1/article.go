package v1

import (
	"strconv"
	"sync"

	"github.com/blog-small-project/global"
	"github.com/blog-small-project/internal/model"
	"github.com/blog-small-project/internal/service"
	"github.com/blog-small-project/pkg/app"
	"github.com/blog-small-project/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type article struct{}

var (
	articleInstance    *article
	getArticleFail     = errcode.NewError(20020001, "取得文章失敗")
	getArticleListFail = errcode.NewError(20020002, "取得文章列表失敗")
	createArticleFail  = errcode.NewError(20020003, "建立文章失敗")
	updateArticleFail  = errcode.NewError(20020004, "更新文章失敗")
	deleteArticleFail  = errcode.NewError(20020005, "刪除文章失敗")
)

func NewArticle() *article {
	var once sync.Once
	if articleInstance == nil {
		once.Do(func() { articleInstance = &article{} })
	}
	return articleInstance
}

func (a article) GetArticle(c *gin.Context) {
	request := app.NewResponse(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errReponse := errcode.InvalidParams.WithDetails(err.Error())
		request.ToErrorResponse(errReponse)
		return
	}

	srv := service.NewArticle(global.MysqlEngine)

	res, err := srv.Get(model.GetArticleRequest{ID: uint32(id)})
	if err != nil {
		request.ToErrorResponse(getArticleFail)
		return
	}

	request.ToResponse(res)
}
func (a article) GetAllArticle(c *gin.Context) {
	request := app.NewResponse(c)
	state := c.Query("state")
	model := model.GetArticleRequest{}
	switch state {
	case "":
		fallthrough
	case "-1":
		model.State = -1
	case "0":
		model.State = 0
	case "1":
		model.State = 1
	default:
		request.ToErrorResponse(errcode.InvalidParams)
		return
	}

	srv := service.NewArticle(global.MysqlEngine)
	res, err := srv.GetAll(model)
	if err != nil {
		request.ToErrorResponse(getArticleListFail)
		return
	}
	request.ToResponse(res)
}
func (a article) CreateArticle(c *gin.Context) {
	srv := service.NewArticle(global.MysqlEngine)
	create := model.CreateArticleRequest{}
	request := app.NewResponse(c)

	err := c.BindJSON(&create)
	if err != nil {
		errReponse := errcode.InvalidParams.WithDetails(err.Error())
		request.ToErrorResponse(errReponse)
		return
	}
	res, err := srv.Create(create)
	if err != nil {
		request.ToErrorResponse(createArticleFail)
		return
	}

	request.ToResponse(res)
	return
}
func (a article) UpdateArticle(c *gin.Context) {
	srv := service.NewArticle(global.MysqlEngine)
	update := model.UpdateArticleRequest{}
	request := app.NewResponse(c)

	err := c.BindJSON(&update)
	if err != nil {
		errResponse := errcode.InvalidParams.WithDetails(err.Error())
		request.ToErrorResponse(errResponse)
		return
	}
	res, err := srv.Update(update)

	if err != nil {
		request.ToErrorResponse(updateArticleFail)
	}

	request.ToResponse(res)
	return
}
func (a article) DeleteArticle(c *gin.Context) {
	srv := service.NewArticle(global.MysqlEngine)
	delete := model.DeleteArticleRequest{}
	err := c.BindJSON(&delete)
	request := app.NewResponse(c)
	if err != nil {
		errReponse := errcode.InvalidParams.WithDetails(err.Error())
		request.ToErrorResponse(errReponse)
		return
	}

	err = srv.Delete(delete)
	if err != nil {
		request.ToErrorResponse(deleteArticleFail)
	}
	request.ToResponse(model.DeleteArticleResponse{ID: delete.ID, IsDel: true})
	return
}
