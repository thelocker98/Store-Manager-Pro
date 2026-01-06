package main

import (
	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"gitea.locker98.com/locker98/Store-Manager-Pro/devices"
	"gitea.locker98.com/locker98/Store-Manager-Pro/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	routes.RegisterRoutes(r)

	go devices.Devices()

	r.Run(":8080")
}
