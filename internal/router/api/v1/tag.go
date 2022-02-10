package v1

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type tag struct{}

var (
	tagInstance *tag
)

func GetTagInstance() *tag {
	lock := &sync.Mutex{}
	if tagInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		tagInstance = &tag{}
	}
	return tagInstance
}

// @Summary 取得單個標籤
// @Produce json
// @Param name query string false "標籤名稱" maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags [get]
func (t tag) GetTag(c *gin.Context)    {}
func (t tag) GetAllTag(c *gin.Context) {}
func (t tag) CreateTag(c *gin.Context) {}
func (t tag) UpdateTag(c *gin.Context) {}
func (t tag) DeleteTag(c *gin.Context) {}
