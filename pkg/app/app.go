package app

import (
	"github.com/blog-small-project/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type response struct {
	ctx *gin.Context
}

func NewResponse(c *gin.Context) *response {
	return &response{ctx: c}
}

func (r *response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "mes": err.Msg()}
	details := err.Details()

	if len(details) > 0 {
		response["details"] = details
	}

	r.ctx.JSON(err.StatusCode(), response)
}
