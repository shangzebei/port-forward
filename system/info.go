package system

import (
	"runtime"
	"os"
	"github.com/shirou/gopsutil/mem"
	"port-info/util"
	"github.com/shirou/gopsutil/process"
	"fmt"
)

type InfoSystem struct {
	CpuNum   int
	System   string
	Arch     string
	Pid      int
	Hostname string
	M_Total  string
	M_Available  string
}

func GetHostInfo() *InfoSystem {
	m, _ :=mem.VirtualMemory()
	info := new(InfoSystem)
	info.Pid = os.Getgid()
	ret, _ := process.NewProcess(int32(info.Pid))
	fmt.Println(ret.MemoryInfo())
	info.M_Total=util.GetBytes(float64(m.Total))
	info.M_Available=util.GetBytes(float64(m.Available))
	info.CpuNum = runtime.NumCPU()
	info.System = runtime.GOOS
	info.Arch = runtime.GOARCH

	info.Hostname,_ = os.Hostname()
	return info
}
