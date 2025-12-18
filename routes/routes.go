package routes

import (
	"github.com/gin-gonic/gin"
)

// Import the handler functions from the same package
// We will implement these in routes/inventory.go
func RegisterRoutes(r *gin.Engine) {
	// Home page
	r.GET("/", LoadHome)

	// API group
	api := r.Group("/api")
	{
		api.GET("/items", GetItems)          // List all items
		api.POST("/items", AddItem)          // Add a new item
		api.PUT("/items/:id", UpdateItem)    // Update an item
		api.DELETE("/items/:id", DeleteItem) // Delete an item

		api.GET("/export/csv", ExportCSV)
		api.GET("/export/pdf", ExportPDF)

	}
}
