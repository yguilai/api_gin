package router

import (
	"api_gin/api"
	"api_gin/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/auth", api.GetAuth)

	// 组路由 v1
	api_v1 := r.Group("/v1")
	// 组路由使用鉴权中间件
	// 改组路由下所有api请求的Header中, 均需携带token
	// Authentication:your token string
	api_v1.Use(jwt.JWTAuth())
	{
		api_v1.GET("/test", api.IndexApi)
	}

	return r
}
