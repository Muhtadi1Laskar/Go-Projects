package operations

import (
    "fmt"
    "net"

    gopsutilnet "github.com/shirou/gopsutil/v3/net"
    "github.com/shirou/gopsutil/v3/process"
)

type ConnectionInfo struct {
    LocalAddr  string
    RemoteAddr string
    Status     string
    PID        int32
    Process    string
}

func MonitorConnections() {
    connections, _ := gopsutilnet.Connections("all")
    
    for _, conn := range connections {
        info := ConnectionInfo{
            LocalAddr:  net.JoinHostPort(conn.Laddr.IP, fmt.Sprintf("%d", conn.Laddr.Port)),
            RemoteAddr: net.JoinHostPort(conn.Raddr.IP, fmt.Sprintf("%d", conn.Raddr.Port)),
            Status:     conn.Status,
            PID:        conn.Pid,
        }
        
        if conn.Pid > 0 {
            p, err := process.NewProcess(conn.Pid)
            if err == nil {
                name, _ := p.Name()
                info.Process = name
            }
        }
        
        fmt.Printf("PID: %d, Process: %s, Local: %s, Remote: %s, Status: %s\n",
            info.PID, info.Process, info.LocalAddr, info.RemoteAddr, info.Status)
    }
}