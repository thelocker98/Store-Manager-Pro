package main

import (
	"os"
	"path/filepath"

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
	r.Run(":8080")
}
