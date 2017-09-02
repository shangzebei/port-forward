package system

import (
	"runtime"
	"os"
)

type InfoSystem struct {
	cpuNum   int
	system   string
	arch     string
	pid      int
	hostname string
}

func GetHostInfo() InfoSystem {
	info := new(InfoSystem)
	info.cpuNum = runtime.NumCPU()
	info.system = runtime.GOOS
	info.arch = runtime.GOARCH
	info.pid = os.Getgid()
	info.hostname,_ = os.Hostname()
	return *info
}
