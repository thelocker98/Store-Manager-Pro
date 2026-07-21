package main

import (
	"os"

	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"gitea.locker98.com/locker98/Store-Manager-Pro/ocr"
	"gitea.locker98.com/locker98/Store-Manager-Pro/routes"
	"gitea.locker98.com/locker98/Store-Manager-Pro/utils"

	"github.com/gin-gonic/gin"
)

var InvoiceFilePath = ""

func main() {
	// ENV vars
	mode := os.Getenv("GIN_MODE")
	appDirEnv := os.Getenv("DATA_PATH")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	// Setup DB
	appDir, _ := utils.SetupFolders()
	if appDirEnv == "" {
		appDirEnv = appDir
	}
	db.InitDB(appDirEnv)

	// Find Unfinished OCR Jobs
	unfinishedEntrys, err := db.GetUnfinishedEntrys()
	if err != nil {
		panic(err)
	}
	// Remove Unfinished Jobs from DB
	for _, entry := range unfinishedEntrys {
		db.DeleteOCREntryById(entry.FileID, true)
		if err != nil {
			panic(err)
		}
	}

	// Configure Web Server
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// Register Routes
	routes.RegisterRoutes(r)

	// Start OCR Jobs
	go ocr.StartOCRWorker()

	// Start Server on Port 8080
	func() {
		if err := r.Run(":8080"); err != nil {
			os.Exit(0)
		}
	}()

}
