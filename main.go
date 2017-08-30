package main

import (
	"net"
	"time"
	//log "github.com/Sirupsen/logrus"
	"log"
	"io"
	"strconv"
)

func main() {
	StartPortForward("0.0.0.0:8700", "0.0.0.0:8088")
}

// sourcePort 源地址和端口，0.0.0.0:8700
// targetPort 数据转发给哪个端口192.168.1.100:3306
func StartPortForward(sourcePort string, targetPort string) {
	go begin()
	localListener, err := net.Listen("tcp", sourcePort)
	if err != nil {
		log.Print("port bind")
		return
	}
	for {
		log.Print("aaaaaa")
		sourceConn, err := localListener.Accept()

		if err != nil {
			break
		}
		//id := sourceConn.RemoteAddr().String()
		//targetPort := "172.16.128.83:22"
		targetConn, err := net.DialTimeout("tcp", targetPort, 30*time.Second)

		go func() {
			log.Print("-->")
			_, err = Copy(targetConn, sourceConn)
			if err != nil {
				log.Println("error",err)
			}
		}()

		go func() {
			log.Print("<---")
			_, err = Copy(sourceConn, targetConn)
			if err != nil {
				log.Println("error",err)
			}
		}()

	}

}

var sum int64

func begin() {
	for true {
		time.Sleep(time.Second)
		log.Println(getSpeed(sum))
		sum = 0
	}
}
func getSpeed(_sum int64) string {
	if _sum > 1048576 {
		return strconv.FormatInt(_sum/1048576, 10) + " MB/s"
	}
	if _sum > 1024 {
		return strconv.FormatInt(_sum/1024, 10) + " KB/s"
	}
	return strconv.FormatInt(_sum, 10) + " B/s"
}
func Copy(src net.Conn, dst net.Conn) (written int64, err error) {
	buf := make([]byte, 32*1024)
	log.Println("local:" + src.LocalAddr().String() + " ==== " + "remote" + dst.RemoteAddr().String())
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			sum += int64(nw)
			//log.Println(nw,ew)
			if nw > 0 {
				written += int64(nw)
			} else {
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				//err = ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	log.Println("end")
	return written, err
}
