package router

import (
	_ "github.com/blog-small-project/docs"
	v1 "github.com/blog-small-project/internal/router/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	{
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	article := v1.GetArticleInstance()
	tag := v1.GetTagInstance()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/article/:id", article.GetArticle)
		apiv1.GET("/articles", article.GetAllArticle)
		apiv1.POST("/article", article.CreateArticle)
		apiv1.PUT("/article/:id", article.UpdateArticle)
		apiv1.DELETE("/article/:id", article.DeleteArticle)
	}

	{
		apiv1.GET("/tag/:id", tag.GetTag)
		apiv1.GET("/tags", tag.GetAllTag)
		apiv1.POST("/tag", tag.CreateTag)
		apiv1.PUT("/tag/:id", tag.UpdateTag)
		apiv1.DELETE("/tag/:id", tag.DeleteTag)
	}

	return r
}
