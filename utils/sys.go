package utils

import (
	"fmt"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUInfo struct {
	CPUNum    int
	AvgUseage float64
	Useages   []float64
}

func GetCPUInfo(interval time.Duration) (*CPUInfo, error) {
	num, err := cpu.Counts(false)
	if err != nil {
		return nil, err
	}
	useages, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, err
	}
	var ret float64
	for _, v := range useages {
		ret += v
	}
	useage := ret / float64(len(useages))

	return &CPUInfo{
		CPUNum:    num,
		AvgUseage: useage,
		Useages:   useages,
	}, nil
}

type MemoryInfo struct {
	Total int64
	Use   int64
	Free  int64
}

func GetMemoryInfo() (*MemoryInfo, error) {
	var memInfo syscall.Sysinfo_t
	if err := syscall.Sysinfo(&memInfo); err != nil {
		// fmt.Println("获取内存信息失败:", err)
		return nil, fmt.Errorf("get memery info failed, %w", err)
	}
	memTotal := memInfo.Totalram * uint64(memInfo.Unit)
	memFree := memInfo.Freeram * uint64(memInfo.Unit)
	memUsed := memTotal - memFree
	return &MemoryInfo{
		Total: int64(memTotal),
		Use:   int64(memUsed),
		Free:  int64(memFree),
	}, nil
}
