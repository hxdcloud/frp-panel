package utils

import (
	"fmt"
	"frp-panel/model"
	"github.com/Erope/goss"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"math"
	"runtime"
	"syscall"
)

// GetHost 获取主机硬件信息
func GetHost() *model.Host {
	hi, _ := host.Info()
	var cpuType string
	if hi.VirtualizationSystem != "" {
		cpuType = "Virtual"
	} else {
		cpuType = "Physical"
	}
	cpuModelCount := make(map[string]int)
	ci, _ := cpu.Info()
	for i := 0; i < len(ci); i++ {
		cpuModelCount[ci[i].ModelName]++
	}
	var cpus []string
	for model, count := range cpuModelCount {
		cpus = append(cpus, fmt.Sprintf("%s %d %s Core", model, count, cpuType))
	}
	mv, _ := mem.VirtualMemory()

	var swapMemTotal uint64
	if runtime.GOOS == "windows" {
		ms, _ := mem.SwapMemory()
		swapMemTotal = ms.Total
	} else {
		swapMemTotal = mv.SwapTotal
	}

	return &model.Host{
		Platform:        hi.Platform,
		PlatformVersion: hi.PlatformVersion,
		CPU:             cpus,
		MemTotal:        mv.Total,
		SwapTotal:       swapMemTotal,
		Arch:            hi.KernelArch,
		Virtualization:  hi.VirtualizationSystem,
		BootTime:        hi.BootTime,
	}
}

func GetState() *model.HostState {
	var procs []int32
	procs, _ = process.Pids()

	mv, _ := mem.VirtualMemory()

	var swapMemUsed uint64
	var swapPercent float64
	if runtime.GOOS == "windows" {
		// gopsutil 在 Windows 下不能正确取 swap
		ms, _ := mem.SwapMemory()
		swapMemUsed = ms.Used
		swapPercent = math.Round(ms.UsedPercent)
	} else {
		swapMemUsed = mv.SwapTotal - mv.SwapFree
	}

	var cpuPercent float64
	cp, err := cpu.Percent(0, false)
	if err == nil {
		cpuPercent = cp[0]
	}
	loadStat, _ := load.Avg()

	var tcpConnCount, udpConnCount uint64
	ss_err := true
	if runtime.GOOS == "linux" {
		tcpStat, err_tcp := goss.ConnectionsWithProtocol(syscall.IPPROTO_TCP)
		udpStat, err_udp := goss.ConnectionsWithProtocol(syscall.IPPROTO_UDP)
		if err_tcp == nil && err_udp == nil {
			ss_err = false
			tcpConnCount = uint64(len(tcpStat))
			udpConnCount = uint64(len(udpStat))
		}
	}
	if ss_err {
		conns, _ := net.Connections("all")
		for i := 0; i < len(conns); i++ {
			switch conns[i].Type {
			case syscall.SOCK_STREAM:
				tcpConnCount++
			case syscall.SOCK_DGRAM:
				udpConnCount++
			}
		}
	}

	return &model.HostState{
		CPU:          math.Round(cpuPercent),
		MemPercent:   math.Round(mv.UsedPercent),
		MemUsed:      mv.Total - mv.Available,
		SwapPercent:  swapPercent,
		SwapUsed:     swapMemUsed,
		Load1:        math.Trunc(loadStat.Load1*1e2+0.5) * 1e-2,
		Load5:        math.Trunc(loadStat.Load5*1e2+0.5) * 1e-2,
		Load15:       math.Trunc(loadStat.Load15*1e2+0.5) * 1e-2,
		TcpConnCount: tcpConnCount,
		UdpConnCount: udpConnCount,
		ProcessCount: uint64(len(procs)),
	}
}
