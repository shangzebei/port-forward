package port

import (
	"math/big"
	"io"
	"log"
	"net"
	"time"
)

type port struct {
	TotalByte    big.Int
	SpeedSumByte int64
	statics      *Statics
	b_stop       bool
	b_pause      bool
}

func StartPortForward(sourcePort string, targetPort string) *port {
	p := port{}
	go p.portForward(sourcePort, targetPort)
	return &p
}

func (p *port) StopForward() {
	p.b_stop = true
}

func (p *port) Pause() {
	p.b_pause = true
}

func (p *port) UnPause() {
	p.b_pause = false
}

func (p *port) SetStatics(statics *Statics) {
	p.statics = statics
}

///////////////////////////////////////////////
func (p *port) portForward(sourcePort string, targetPort string) {

	go p.staticsPort()

	localListener, err := net.Listen("tcp", sourcePort)
	if err != nil {
		log.Print("port bind")
		return
	}

	for !p.b_stop {

		sourceConn, err := localListener.Accept()

		if p.b_pause {
			break
		}

		if err != nil {
			break
		}
		//id := sourceConn.RemoteAddr().String()
		//targetPort := "172.16.128.83:22"
		targetConn, err := net.DialTimeout("tcp", targetPort, 30*time.Second)

		go func() {
			log.Print("-->")
			_, err = p.copy(targetConn, sourceConn)
			if err != nil {
				log.Println("error", err)
			}
		}()

		go func() {
			log.Print("<---")
			_, err = p.copy(sourceConn, targetConn)
			if err != nil {
				log.Println("error", err)
			}
		}()

	}

}

func (p *port) staticsPort() {
	for true {
		time.Sleep(time.Second)
		if p.statics != nil {
			(*p.statics).StaticInfo(p.SpeedSumByte, p)
		}
		p.SpeedSumByte = 0
	}
}

func (p *port) copy(src net.Conn, dst net.Conn) (written int64, err error) {
	defer src.Close()
	defer dst.Close()
	buf := make([]byte, 1048576) //1M
	log.Println("local:" + src.LocalAddr().String() + " ==== " + "remote" + dst.RemoteAddr().String())
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			p.SpeedSumByte += int64(nw)
			p.TotalByte.Add(big.NewInt(p.SpeedSumByte), &p.TotalByte)
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
	return written, err
}
