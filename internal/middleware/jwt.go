package middleware

import (
	"github.com/blog-small-project/pkg/app"
	"github.com/blog-small-project/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ecode = errcode.Success
		if s := c.GetHeader("token"); len(s) > 0 {
			_, err := app.ParseToken(s)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		} else {
			ecode = errcode.UnauthorizedNotToken
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
