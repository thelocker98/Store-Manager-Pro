package main

import (
	"log"
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
	"gitea.locker98.com/locker98/Store-Manager-Pro/utils"

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
	go func() {
		if err := r.Run(":8080"); err != nil {
			os.Exit(0)
		}
	}()

	//---------------------------------------------------
	a := app.New()
	w := a.NewWindow("Local IP Addresses")
	w.Resize(fyne.NewSize(300, 400))

	icon, err := fyne.LoadResourceFromPath("static/icon.png")
	if err != nil {
		panic(err)
	}

	w.SetIcon(icon)

	// Build IP list UI

	title := widget.NewLabelWithStyle(
		"Store Manager Pro",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)

	ipList := container.NewVBox(
		title,
		widget.NewLabel("Click on IP the interface"),
	)

	for _, ip := range utils.GetLocalIPs() {
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
			fyne.NewMenuItem("Open Database Folder", func() {
				err := utils.OpenFileExplorer(appDir)
				if err != nil {
					log.Println("Failed to open file explorer:", err)
				}
			}),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		)
		desk.SetSystemTrayMenu(menu)
	}

	// Start hidden
	w.ShowAndRun()
}
