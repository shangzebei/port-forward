package web

import (
	"github.com/gin-gonic/gin"
)

func InitConfig() *gin.Engine {
	route := gin.Default()
	route.StaticFS("static", assets)
	bindRest(route)
	return route
}
func bindRest(route *gin.Engine) {
	v1 := route.Group("/v1")
	{
		v1.POST("startPort", startPortForward)
		v1.POST("stopPort", stopPort)
		v1.GET("listAll", listAllPort)
		v1.POST("setSpeed", setSpeed)
		v1.GET("getSystemInfo", getSystemInfo)
		v1.GET("info", func(c *gin.Context) {
			echo(c.Writer, c.Request)
		})
	}

}

