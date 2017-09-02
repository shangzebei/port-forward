package port

import (
	"math/big"
)

type Limit struct {
	maxBytes   big.Int
	useByte    big.Int
	speedBytes int
}
//50M  30M
func (kk *Limit) AddHasByte(b int64) {
  kk.maxBytes.Add(&kk.maxBytes,big.NewInt(int64(b)))
}
func (kk *Limit) ReSetSpeed()  {

}

func (kk *Limit) StaticInfo(copyBytes int64, port *port) () {
	kk.useByte.Add(&kk.useByte,big.NewInt(copyBytes))
	if kk.maxBytes.Sub(&kk.maxBytes,&kk.useByte).Int64()<= 0 {
		port.StopForward()
	}
}
