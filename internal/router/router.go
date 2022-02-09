package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"test": "ok"}) })
	}

	return r
}
