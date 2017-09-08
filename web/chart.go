package web

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"port-info/port"
	"fmt"
	"encoding/json"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()
	mt, ms, err := c.ReadMessage()
	for {
		//fmt.Println("hello")
		time.Sleep(time.Second*3)
		err = c.WriteMessage(mt, GetInfo(string(ms)))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func GetInfo(ms string) []byte {
	fmt.Println(ms)
	by := port.ForwardPoll[ms]
	res := map[string]int64{"speed": 0}
	if by != nil {
		res["speed"]=by.SpeedPeerByte
	}
	bt, _ := json.Marshal(res)
	return bt
}
