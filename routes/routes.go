package routes

import (
	"github.com/gin-gonic/gin"
)

// Import the handler functions from the same package
// We will implement these in routes/inventory.go
func RegisterRoutes(r *gin.Engine) {
	// Home page
	r.GET("/", LoadHome)

	// Vendors Page
	r.GET("/vendors", LoadVendor)

	// API group
	api := r.Group("/api")
	{

		// Items
		api.GET("/items", GetItems)          // List all items
		api.POST("/items", AddItem)          // Add a new item
		api.PUT("/items/:id", UpdateItem)    // Update an item
		api.DELETE("/items/:id", DeleteItem) // Delete an item

		// Vendors
		api.GET("/vendors", GetVendors)          // List all vendors
		api.POST("/vendors", AddVendor)          // Add a new vendor
		api.PUT("/vendors/:id", UpdateVendor)    // Update an vendor
		api.DELETE("/vendors/:id", DeleteVendor) // Delete an vendor

		// Catalog
		api.GET("/catalog", GetCatalog)                // List all catalog entrys
		api.POST("/catalog", AddCatalogEntry)          // Add a new catalog entry
		api.PUT("/catalog/:id", UpdateCatalogEntry)    // Update an catalog entry
		api.DELETE("/catalog/:id", DeleteCatalogEntry) // Delete an catalog entry

		// Export
		api.GET("/export/csv", ExportCSV)
		api.GET("/export/pdf", ExportPDF)

	}
}
