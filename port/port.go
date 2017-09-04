package port

import (
	"math/big"
	"io"
	"net"
	"time"
	"github.com/juju/ratelimit"
	"log"
)

type port struct {
	bucket        *ratelimit.Bucket
	localListener *net.Listener
	LocalPort     string
	TotalByte     big.Int
	SpeedPeerByte int64
	LimitSpeed    int64
	B_stop        bool
	B_pause       bool

	statics []Statics
}

var ForwardPoll = make(map[string]*port)

func StartPortForward(sourcePort string, targetPort string) *port {
	p := port{}

	go p.processPort(sourcePort, targetPort)
	p.statics = make([]Statics, 5)
	return &p
}

func (p *port) StopForward() {
	p.B_stop = true
	(*p.localListener).Close()
	delete(ForwardPoll, p.LocalPort)
}
func (p *port) SetSpeed(bytes int64) {
	p.LimitSpeed = bytes
}
func (p *port) Pause() {
	p.B_pause = true
}

func (p *port) UnPause() {
	p.B_pause = false
}

func (p *port) AddStatics(statics Statics) {
	p.statics = append(p.statics, statics)
}

///////////////////////////////////////////////
func (p *port) processPort(sourcePort string, targetPort string) {

	p.LocalPort = sourcePort

	go p.staticsPort()

	ForwardPoll[p.LocalPort] = p

	if p.LimitSpeed != 0 {
		p.bucket = ratelimit.NewBucketWithRate(float64(p.LimitSpeed), int64(p.LimitSpeed))
	}

	localListener, err := net.Listen("tcp", sourcePort)
	p.localListener = &localListener
	if err != nil {
		log.Print("port bind")
		return
	}

	for !p.B_stop {
		log.Print(p.B_pause)
		sourceConn, err := localListener.Accept()

		if err != nil {
			break
		}
		//id := sourceConn.RemoteAddr().String()
		//targetPort := "172.16.128.83:22"
		targetConn, err := net.DialTimeout("tcp", targetPort, 30*time.Second)

		go func() {
			_, err = p.copy(targetConn, sourceConn)
			if err != nil {
				log.Println("error", err)
			}
		}()

		go func() {
			_, err = p.copy(sourceConn, targetConn)
			if err != nil {
				log.Println("error", err)
			}
		}()

	}

	delete(ForwardPoll, sourcePort)
}

func (p *port) staticsPort() {
	for {
		time.Sleep(time.Second * 3)
		//fmt.Println(key,value)
		for _, value := range p.statics {
			if value != nil {
				value.StaticInfo(p.SpeedPeerByte/3, p)
			}
		}
		p.SpeedPeerByte = 0
	}
}

func (p *port) copy(src net.Conn, dst net.Conn) (written int64, err error) {
	defer src.Close()
	defer dst.Close()
	buf := make([]byte, 1024*16) //1M

	for {
		var nr int
		var er error
		if p.B_stop || p.B_pause {
			break
		}
		if p.LimitSpeed == 0 {
			nr, er = src.Read(buf)
		} else {
			nr, er = ratelimit.Reader(src, p.bucket).Read(buf)
		}
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			p.SpeedPeerByte += int64(nw)
			p.TotalByte.Add(big.NewInt(p.SpeedPeerByte), &p.TotalByte)
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
