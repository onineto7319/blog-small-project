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

func NewArticle() *article {
	var once sync.Once
	if articleInstance == nil {
		once.Do(func() { articleInstance = &article{} })
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
