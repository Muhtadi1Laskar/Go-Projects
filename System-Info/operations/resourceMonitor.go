package operations

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type SystemStats struct {
	Timestamp   time.Time
	CPUUsage    float64
	MemoryUsed  float64
	MemoryTotal uint64
	DiskUsage   float64
	NetworkSent uint64
	NetworkRecv uint64
}

func collectSystemStats() SystemStats {
	cpuPercents, _ := cpu.Percent(time.Second, false)

	vmStat, _ := mem.VirtualMemory()

	diskStat, _ := disk.Usage("/")

	netStats, _ := net.IOCounters(true)
	var sent, recv uint64

	for _, netStat := range netStats {
		sent += netStat.BytesSent
		recv += netStat.BytesRecv
	}

	return SystemStats{
		Timestamp:   time.Now(),
		CPUUsage:    cpuPercents[0],
		MemoryUsed:  vmStat.UsedPercent,
		MemoryTotal: vmStat.Total,
		DiskUsage:   diskStat.UsedPercent,
		NetworkSent: sent,
		NetworkRecv: recv,
	}
}

func StartMonitoring() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		stats := collectSystemStats()

		fmt.Printf("[%s] CPU: %.1f%%, Memory: %.1f%%, Disk: %.1f%%, Network Sent: %d, Network Recieved: %d\n",
			stats.Timestamp.Format("15:04:05"),
			stats.CPUUsage,
			stats.MemoryUsed,
			stats.DiskUsage,
			stats.NetworkSent,
			stats.NetworkRecv,
		)
	}
}
