package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", nil)

	api := r.Group("/api")
	{
		api.GET("/items", nil)
		api.POST("/items", nil)
		api.PUT("/items/:id", nil)
		api.DELETE("/items/:id", nil)
	}
}
