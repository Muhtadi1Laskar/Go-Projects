package operations

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"time"
)

type ProcessInfo struct {
	PID    int32
	Name   string
	CPU    float64
	Memory float32
	Status string
}

func MonitorProcesses() {
	for {
		processes, _ := process.Processes()

		for _, p := range processes {
			name, _ := p.Name()
			cpu, _ := p.CPUPercent()
			mem, _ := p.MemoryPercent()
			status, _ := p.Status()

			info := ProcessInfo{
				PID:    p.Pid,
				Name:   name,
				CPU:    cpu,
				Memory: mem,
				Status: status[0],
			}

			fmt.Printf("PID: %d, Name: %s, CPU: %.2f%%, Memory: %.2f%%\n",
				info.PID, info.Name, info.CPU, info.Memory)
		}

		time.Sleep(2 * time.Second)
		fmt.Println()
		fmt.Println()
	}
}
