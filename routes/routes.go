package routes

import (
	"github.com/gin-gonic/gin"
)

// Import the handler functions from the same package
// We will implement these in routes/inventory.go
func RegisterRoutes(r *gin.Engine) {
	// Page
	r.GET("/", LoadHome)              // Home page
	r.GET("/vendors", LoadVendor)     // Vendors Page
	r.GET("/locations", LoadLocation) // Locations Page
	r.GET("/lists", LoadLists)        // Barcode Page
	r.GET("/barcode", LoadBarcode)    // Barcode Page
	r.GET("/export", LoadExport)      // Export Page

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

		// Locations
		api.GET("/locations", GetLocations)          // List all locations
		api.POST("/locations", AddLocation)          // Add a new location
		api.PUT("/locations/:id", UpdateLocation)    // Update an location
		api.DELETE("/locations/:id", DeleteLocation) // Delete an location

		// Lists
		api.GET("/lists", GetLists)          // Get All Lists with their name, id, date
		api.GET("/lists/:id", GetListById)   // Get List Using its ID
		api.POST("/lists", AddList)          // Add a new List
		api.PUT("/lists/:id", UpdateList)    // Update a List name
		api.DELETE("/lists/:id", DeleteList) // Delete a List
		// List Entrys
		api.GET("/listentrys/:listid", GetListEntrys)           // Get All Entrys in a List id, list_id, item_id
		api.POST("/listentrys/:listid", AddEntryToList)         // Add a new Entry to an existing List
		api.PUT("/listentrys/:entryid", UpdateEntryInList)      // Update Entry in an existing List
		api.DELETE("/listentrys/:entryid", DeleteEntryFromList) // Delete Entry From A existing List

		// Search
		api.GET("/search/:q", SearchItems)

		// Barcode Websocket
		api.GET("/ws", BarcodeReaderWS)

		// Export
		api.GET("/export/csv", ExportCSV)
		api.GET("/export/pdf", ExportPDF)

	}
}
