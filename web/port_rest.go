package web

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"port-info/port"
	"port-info/util"
	"strconv"
	"port-info/system"
	"encoding/json"
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
	//a.AddStatics(&port.Limit{})
}

func stopPort(c *gin.Context) {
	src := c.PostForm("port")
	if !util.CheckParam(src) {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	port.ForwardPoll[src].StopForward()
}

func listAllPort(c *gin.Context) {
	arry := make([]interface{}, len(port.ForwardPoll))
	for index, value := range port.ForwardPoll {
		if value == nil {
			continue
		}
		info := make(map[string]interface{})
		info["src"] = index
		info["dst"] = value.TargetPort
		info["run"] = value.B_stop
		arry = append(arry, info)
	}
	c.JSON(http.StatusOK, arry)

}

func setSpeed(c *gin.Context)  {
	port_t:=c.PostForm("port")
	speed:=c.PostForm("speed")
	sp,_:=strconv.ParseInt(speed,10,64)
	port.ForwardPoll[port_t].SetSpeed(sp)
}

func getSystemInfo(c* gin.Context) {
	info:=system.GetHostInfo()
	by,_:=json.Marshal(&info)
	fmt.Println(string(by))
	c.JSON(http.StatusOK,info)
}