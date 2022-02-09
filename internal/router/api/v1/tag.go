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

func (t tag) GetTag(c *gin.Context)    {}
func (t tag) GetAllTag(c *gin.Context) {}
func (t tag) CreateTag(c *gin.Context) {}
func (t tag) UpdateTag(c *gin.Context) {}
func (t tag) DeleteTag(c *gin.Context) {}
