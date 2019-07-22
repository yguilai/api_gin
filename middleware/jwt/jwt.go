package jwt

import (
	e "api_gin/pkg/e"
	"api_gin/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// jwt token 鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			code int
			data interface{}
			key  = "Authentication"
		)

		code = e.SUCCESS
		// 从请求头中获取token
		token := c.Request.Header.Get(key)
		if token == "" {
			code = e.ERROR_AUTH
		}

		if token != "" {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Set("")
		c.Next()
	}
}
