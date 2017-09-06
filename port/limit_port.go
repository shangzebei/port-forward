package port

import (
	"math/big"
	"fmt"
	"port-info/util"
)

type Limit struct {
	maxBytes   big.Int
	useByte    big.Int
	speedBytes int
}

//50M  30M
func (kk *Limit) AddHasByte(b int64) {
	kk.maxBytes.Add(&kk.maxBytes, big.NewInt(int64(b)))
}
func (kk *Limit) ReSetSpeed() {

}
func (kk *Limit) StaticInfo(copyBytes int64, port *port) () {
	fmt.Println(util.GetBytePeerSecond(float64(copyBytes)))
	kk.useByte.Add(&kk.useByte, big.NewInt(copyBytes))
	if kk.maxBytes.Int64() > 0 &&
		kk.maxBytes.Sub(&kk.maxBytes, &kk.useByte).Int64() <= 0 {
		port.StopForward()
	}
}
