package v1

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/blog-small-project/internal/model"

	"github.com/blog-small-project/pkg/errcode"

	"github.com/blog-small-project/global"
	"github.com/blog-small-project/internal/service"
	"github.com/blog-small-project/pkg/app"

	"github.com/gin-gonic/gin"
)

type tag struct{}

var (
	tagInstance    *tag
	getTagFail     = errcode.NewError(20010001, "取得標籤失敗")
	getTagListFail = errcode.NewError(20010002, "取得標籤列表失敗")
	createTagFail  = errcode.NewError(20010003, "建立標籤失敗")
	updateTagFail  = errcode.NewError(20010004, "更新標籤失敗")
	deleteTagFail  = errcode.NewError(20010005, "刪除標籤失敗")
)

func NewTag() *tag {
	var once sync.Once
	if tagInstance == nil {
		once.Do(func() { tagInstance = &tag{} })
	}
	return tagInstance
}

// @Summary 取得單個標籤
// @Produce json
// @Param id query int false "標籤編號"
// @Success 200 {object} model.GetTagResponse "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tag/:id [get]
func (t tag) GetTag(c *gin.Context) {
	request := app.NewResponse(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errReponse := errcode.InvalidParams.WithDetails(err.Error())
		request.ToErrorResponse(errReponse)
		return
	}

	srv := service.NewTag(global.MysqlEngine)

	res, err := srv.Get(model.GetTagRequest{ID: uint32(id)})
	if err != nil {
		request.ToErrorResponse(getTagFail)
		return
	}

	request.ToResponse(res)
}

// @Summary 取得標籤列表
// @Produce json
// @Param state query int false "狀態"
// @Success 200 {object} model.GetTagResponse "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags [get]
func (t tag) GetAllTag(c *gin.Context) {
	request := app.NewResponse(c)
	state := c.Query("state")
	model := model.GetTagRequest{}
	fmt.Println("state:", state)
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

	srv := service.NewTag(global.MysqlEngine)
	res, err := srv.GetAll(model)
	if err != nil {
		request.ToErrorResponse(getTagListFail)
		return
	}
	request.ToResponse(res)
}

// @Summary 新增標籤
// @Produce json
// @Param id query int false "標籤編號"
// @Success 200 {object} model.CreateTagResponse "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tag [post]
func (t tag) CreateTag(c *gin.Context) {
	srv := service.NewTag(global.MysqlEngine)
	create := model.CreateTagRequest{}
	request := app.NewResponse(c)

	err := c.BindJSON(&create)
	if err != nil {
		errReponse := errcode.InvalidParams.WithDetails(err.Error())
		request.ToErrorResponse(errReponse)
		return
	}
	res, err := srv.Create(create)
	if err != nil {
		request.ToErrorResponse(createTagFail)
		return
	}

	request.ToResponse(res)
	return
}

// @Summary 更新標籤
// @Produce json
// @Param id query int false "標籤編號"
// @Success 200 {object} model.UpdateTagResponse "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tag [put]
func (t tag) UpdateTag(c *gin.Context) {
	srv := service.NewTag(global.MysqlEngine)
	update := model.UpdateTagRequest{}
	request := app.NewResponse(c)

	err := c.BindJSON(&update)
	if err != nil {
		errResponse := errcode.InvalidParams.WithDetails(err.Error())
		request.ToErrorResponse(errResponse)
		return
	}
	res, err := srv.Update(update)

	if err != nil {
		request.ToErrorResponse(updateTagFail)
	}

	request.ToResponse(res)
	return
}

// @Summary 刪除標籤
// @Produce json
// @Param id query int false "標籤編號"
// @Success 200 {object} model.DeleteTagResponse "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tag [delete]
func (t tag) DeleteTag(c *gin.Context) {
	srv := service.NewTag(global.MysqlEngine)
	delete := model.DeleteTagRequest{}
	err := c.BindJSON(&delete)
	request := app.NewResponse(c)
	if err != nil {
		errReponse := errcode.InvalidParams.WithDetails(err.Error())
		request.ToErrorResponse(errReponse)
		return
	}

	res, err := srv.Delete(delete)
	if err != nil {
		request.ToErrorResponse(deleteTagFail)
	}
	request.ToResponse(model.DeleteTagResponse{ID: res, IsDel: true})
	return
}
