package main

import (
	"net"
	"net/url"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"gitea.locker98.com/locker98/Store-Manager-Pro/devices"
	"gitea.locker98.com/locker98/Store-Manager-Pro/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	appData, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appDir := filepath.Join(appData, "StoreManagerPro")
	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		panic(err)
	}

	// Setup DB
	db.InitDB(appDir)

	// Configure Web Server
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// Register Routes
	routes.RegisterRoutes(r)

	// Start device drivers
	go devices.Devices()

	// Start Server on Port 8080
	go r.Run(":8080")

	//---------------------------------------------------
	a := app.New()
	w := a.NewWindow("Local IP Addresses")
	w.Resize(fyne.NewSize(300, 400))

	// Build IP list UI
	ipList := container.NewVBox(
		widget.NewLabel("Click an IP to open http://IP:8080"),
	)

	for _, ip := range getLocalIPs() {
		ip := ip // capture
		btn := widget.NewButton(ip, func() {
			u, _ := url.Parse("http://" + ip + ":8080")
			a.OpenURL(u)
		})
		ipList.Add(btn)
	}

	w.SetContent(container.NewVScroll(ipList))

	// Hide window instead of closing
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	// System tray
	if desk, ok := a.(desktop.App); ok {
		menu := fyne.NewMenu("IP Tray",
			fyne.NewMenuItem("Show Window", func() {
				w.Show()
				w.RequestFocus()
			}),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		)
		desk.SetSystemTrayMenu(menu)
	}

	// Start hidden
	w.Hide()
	a.Run()
}

func getLocalIPs() []string {
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
