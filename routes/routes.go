package routes

import (
	"github.com/gin-gonic/gin"
)

// Import the handler functions from the same package
// We will implement these in routes/inventory.go
func RegisterRoutes(r *gin.Engine) {
	// Page
	r.GET("/", LoadHome)           // Home page
	r.GET("/vendors", LoadVendor)  // Vendors Page
	r.GET("/export", LoadExport)   // Export Page
	r.GET("/barcode", LoadBarcode) // Barcode Page

	// API group
	api := r.Group("/api")
	{
		// Items
		api.GET("/items", GetItems)                             // List all items
		api.GET("/items/count", CountItems)                     // Count of all items
		api.GET("/items/:id", GetItemById)                      // Get item by id
		api.POST("/items", AddItem)                             // Add a new item
		api.PUT("/items/:id", UpdateItem)                       // Update an item
		api.DELETE("/items/:id", DeleteItem)                    // Delete an item
		api.GET("/items/restore/:id", RestoreItem)              // Delete an item
		api.DELETE("/items/permanent/:id", DeleteItemPermanent) // Delete an item

		// Vendors
		api.GET("/vendors", GetVendors)          // List all vendors
		api.POST("/vendors", AddVendor)          // Add a new vendor
		api.PUT("/vendors/:id", UpdateVendor)    // Update an vendor
		api.DELETE("/vendors/:id", DeleteVendor) // Delete an vendor

		// Search
		api.GET("/search/:q", SearchItems)

		// Barcode Websocket
		api.GET("/ws", BarcodeReaderWS)

		// Export
		api.GET("/export/csv", ExportCSV)
		api.GET("/export/pdf", ExportPDF)

	}
}
