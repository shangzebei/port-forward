package web

import (
	"github.com/gin-gonic/gin"
)

func InitConfig() *gin.Engine {
	gin := gin.Default()
	bindRest(gin)
	return gin
}
func bindRest(gin *gin.Engine) {
	gin.POST("v1/startPort", startPortForward)
	gin.POST("v1/stopPort", stopPort)
	gin.GET("v1/listAll", listAllPort)
	gin.POST("v1/setSpeed", setSpeed)
	gin.GET("v1/getSystemInfo",getSystemInfo)
}
