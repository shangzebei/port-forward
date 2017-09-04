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
	gin.POST("startPort", startPortForward)
	gin.POST("stopPort", stopPort)
	gin.GET("listAll", listAllPort)
}
