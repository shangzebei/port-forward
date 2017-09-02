package port

import (
	"fmt"
	"port-info/util"
	"math/big"
)

type Limit struct {
	hasByte    big.Int
	speedBytes int
}

func (kk *Limit) StaticInfo(copyBytes int64, port *port) () {
	fmt.Println(util.GetSpeed(copyBytes))
	if port.TotalByte.Int64() > 381209328 {
		fmt.Println("======stop====")

		port.StopForward()

	}
}
