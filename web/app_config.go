package web

import (
	"github.com/gin-gonic/gin"
)

func InitConfig() *gin.Engine {
	gin := gin.Default()
	gin.Static("static","./asset")
	bindRest(gin)
	return gin
}
func bindRest(gin *gin.Engine) {
	v1 := gin.Group("/v1")
	{
		v1.POST("startPort", startPortForward)
		v1.POST("stopPort", stopPort)
		v1.GET("listAll", listAllPort)
		v1.POST("setSpeed", setSpeed)
		v1.GET("getSystemInfo", getSystemInfo)
	}

}
