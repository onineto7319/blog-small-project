package v1

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type article struct{}

var (
	articleInstance *article
)

func GetArticleInstance() *article {
	lock := &sync.Mutex{}
	if articleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		articleInstance = &article{}
	}
	return articleInstance
}

func (a article) GetArticle(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"status": "todo"})
}
func (a article) GetAllArticle(c *gin.Context) {}
func (a article) CreateArticle(c *gin.Context) {}
func (a article) UpdateArticle(c *gin.Context) {}
func (a article) DeleteArticle(c *gin.Context) {}
