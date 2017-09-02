package port

import (
	"fmt"
	"port-info/util"
)

type Limit struct {
}

func (kk *Limit) StaticInfo(copyBytes int64, port *port) () {
	fmt.Println(util.GetSpeed(copyBytes))
	if port.TotalByte.Int64() > 381209328 {
		fmt.Println("======stop====")

			port.StopForward()

	}
	fmt.Println(port.TotalByte.Int64())
}
