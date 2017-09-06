package system

import (
	"runtime"
	"os"
	"github.com/shirou/gopsutil/mem"
	"port-info/util"
	"github.com/shirou/gopsutil/process"
)

type InfoSystem struct {
	CpuNum      int
	System      string
	Arch        string
	Pid         int
	Hostname    string
	M_Total     string
	M_Available string
	P_rss       string
	P_vms       string
	P_swap      string
}

func GetHostInfo() *InfoSystem {
	info := new(InfoSystem)
	m, _ := mem.VirtualMemory()
	info.Pid = os.Getpid()

	ps, _ := process.NewProcess(int32(info.Pid))
	mi, _ := ps.MemoryInfo()
	info.P_rss = util.GetBytes(float64(mi.RSS))
	info.P_vms = util.GetBytes(float64(mi.VMS))
	info.P_swap = util.GetBytes(float64(mi.Swap))

	info.M_Total = util.GetBytes(float64(m.Total))
	info.M_Available = util.GetBytes(float64(m.Available))
	info.CpuNum = runtime.NumCPU()
	info.System = runtime.GOOS
	info.Arch = runtime.GOARCH

	info.Hostname, _ = os.Hostname()
	return info
}
