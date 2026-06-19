package utils

import (
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var DBFolderPath string
var InvoiceFolderPath string

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

func SetupFolders() (string, string) {
	// Create App Directory
	appData, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	DBFolderPath = filepath.Join(appData, "StoreManagerPro")
	err = os.MkdirAll(DBFolderPath, 0755)
	if err != nil {
		panic(err)
	}

	// Setup Invoice Upload Folder
	InvoiceFolderPath = filepath.Join(DBFolderPath, "InvoiceFolder")
	err = os.RemoveAll(InvoiceFolderPath)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(InvoiceFolderPath, 0755)
	if err != nil {
		panic(err)
	}

	return DBFolderPath, InvoiceFolderPath
}
