package router

import (
	"apigin/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", api.IndexApi)
	return router
}
