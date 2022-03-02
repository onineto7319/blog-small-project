package router

import (
	_ "github.com/blog-small-project/docs"
	"github.com/blog-small-project/internal/middleware"
	v1 "github.com/blog-small-project/internal/router/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	auth := v1.NewAuth()

	{
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.POST("/auth", auth.GetAuth)
	}

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("/articles/:id", article.GetArticle)
		apiv1.GET("/articles", article.GetAllArticle)
		apiv1.POST("/article", article.CreateArticle)
		apiv1.PUT("/article", article.UpdateArticle)
		apiv1.DELETE("/article", article.DeleteArticle)
	}

	{
		apiv1.GET("/tags/:id", tag.GetTag)
		apiv1.GET("/tags", tag.GetAllTag)
		apiv1.POST("/tag", tag.CreateTag)
		apiv1.PUT("/tag", tag.UpdateTag)
		apiv1.DELETE("/tag", tag.DeleteTag)
	}

	return r
}
