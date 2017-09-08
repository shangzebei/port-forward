package web

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"port-info/port"
	"port-info/util"
	"port-info/system"
)

func startPortForward(c *gin.Context) {
	src := c.PostForm("src")
	dst := c.PostForm("dst")
	if !util.CheckParam(src, dst) {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	fmt.Println("src== "+src, "dst= "+dst)
	if port.ForwardPoll[util.GetPort(src)] != nil {
		c.JSON(http.StatusOK, gin.H{"state": "hasBind"})
		return
	}
	a := port.StartPortForward(src, dst)
	a.LimitSpeed = 0
	c.JSON(http.StatusOK, gin.H{"state": "ok"})
	a.AddStatics(&port.Limit{})
}

func stopPort(c *gin.Context) {
	src := c.PostForm("port")
	if !util.CheckParam(src) {
		c.JSON(http.StatusNotFound, gin.H{"state": "no param"})
		return
	}
	find_port := port.ForwardPoll[src]
	if find_port == nil {
		c.JSON(http.StatusOK, gin.H{"state": "not find port"})
		return
	}
	port.ForwardPoll[src].StopForward()
	c.JSON(http.StatusOK, gin.H{"state": "ok"})
}

func listAllPort(c *gin.Context) {
	type Info struct {
		Src        string
		Dst        string
		B_run      bool
		UseBytes   string
		LimitSpeed string
	}
	var arrar []interface{}
	for _, value := range port.ForwardPoll {
		info := Info{
			value.LocalPort,
			value.TargetPort,
			value.B_stop,
			util.GetBytes(float64(value.TotalByte.Uint64())),
			util.GetBytes(float64(value.LimitSpeed)) + "/s",
		}
		arrar=append(arrar, info)
	}
	fmt.Println(arrar)
	c.JSON(http.StatusOK, arrar)

}

func setSpeed(c *gin.Context) {
	port_t := c.PostForm("port")
	speed := c.PostForm("speed")
	value := util.GetByteFromString(speed)
	if !util.CheckParam(port_t, speed) {
		c.JSON(http.StatusOK, gin.H{"state": "port or speed param not find"})
		return
	}
	v_port := port.ForwardPoll[port_t]
	v_port.SetSpeed(int64(value))
	c.JSON(http.StatusOK, gin.H{"state": "ok"})
}

func getSystemInfo(c *gin.Context) {
	info := system.GetHostInfo()
	c.JSON(http.StatusOK, info)
}
