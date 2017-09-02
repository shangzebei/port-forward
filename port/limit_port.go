package port

import (
	"fmt"
	"port-info/util"
	"math/big"
)

type Limit struct {
	maxBytes   big.Int
	hasByte    big.Int
	speedBytes int
}
//50M  30M
func (kk *Limit) ResetHasByte() {


}
func (kk *Limit) ReSetSpeed()  {



}

func (kk *Limit) StaticInfo(copyBytes int64, port *port) () {
	fmt.Println(util.GetSpeed(copyBytes))
	if port.TotalByte.Int64() > 381209328 {
		fmt.Println("======stop====")

		port.StopForward()

	}
}
