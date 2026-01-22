package utils

import (
	"net"
	"os/exec"
	"runtime"
)

func GetLocalIPs() []string {
	var ips []string

	ifaces, err := net.Interfaces()
	if err != nil {
		return ips
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue
			}

			ips = append(ips, ip.String())
		}
	}

	return ips
}

func OpenFileExplorer(path string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		// Opens Explorer at the path
		cmd = exec.Command("explorer", path)

	case "linux":
		// Uses the user's default file manager
		cmd = exec.Command("dolphin", path)

	case "darwin":
		// macOS Finder
		cmd = exec.Command("open", path)

	default:
		return nil
	}

	return cmd.Start()
}
