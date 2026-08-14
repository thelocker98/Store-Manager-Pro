package routes

import (
	"gitea.locker98.com/locker98/Store-Manager-Pro/utils"
	"github.com/gin-gonic/gin"
)

// Import the handler functions from the same package
// We will implement these in routes/inventory.go
func RegisterRoutes(r *gin.Engine) {
	// Page
	r.GET("/", LoadHome)                             // Home page
	r.GET("/vendors", LoadVendor)                    // Vendors Page
	r.GET("/locations", LoadLocation)                // Locations Page
	r.GET("/departments", LoadDepartments)           // Department Page
	r.GET("/lists", LoadLists)                       // List Page
	r.GET("/lists/:listid", LoadListEditor)          // List Entry Page
	r.GET("/invoices", LoadInvoices)                 // Invoice Page
	r.GET("/invoices/:invoiceid", LoadInvoiceEditor) // Invoice Entry Page
	r.GET("/export", LoadExport)                     // Export Page

	// API group
	api := r.Group("/api")
	{
		// Items
		api.GET("/items", GetItems)                             // List all items
		api.GET("/items/count", CountItems)                     // Count of all items
		api.GET("/items/:id", GetItemById)                      // Get item by id
		api.POST("/items", AddItem)                             // Add a new item
		api.PUT("/items/:id", UpdateItem)                       // Update an item
		api.PUT("/items/department/:id", UpdateItemDepartment)  // Update Deparmtent and Location on item
		api.DELETE("/items/:id", DeleteItem)                    // Delete an item
		api.GET("/items/restore/:id", RestoreItem)              // Restore an item
		api.DELETE("/items/permanent/:id", DeleteItemPermanent) // Permentently Delete an item

		// Vendors
		api.GET("/vendors", GetVendors)          // List all vendors
		api.POST("/vendors", AddVendor)          // Add a new vendor
		api.PUT("/vendors/:id", UpdateVendor)    // Update an vendor
		api.DELETE("/vendors/:id", DeleteVendor) // Delete an vendor

		// Locations
		api.GET("/locations/", GetLocations)         // List all locations
		api.POST("/locations", AddLocation)          // Add a new location
		api.PUT("/locations/:id", UpdateLocation)    // Update an location
		api.DELETE("/locations/:id", DeleteLocation) // Delete an location

		// Lists
		api.GET("/lists/d/:departmentid", GetLists)           // Get All Lists with their name, id, date
		api.GET("/lists/:listid", GetListById)                // Get List Using its ID
		api.GET("/lists/count/:listid", GetNumberOfItemsList) // Get Number of Items in a given list
		api.POST("/lists", AddList)                           // Add a new List
		api.PUT("/lists/:listid", UpdateList)                 // Update a List name
		api.DELETE("/lists/:listid", DeleteList)              // Delete a List

		// List Entrys
		api.GET("/listentrys/:listid", GetListEntrys)          // Get All Entrys in a List id, list_id, item_id
		api.GET("/listentry/:entryid", GetListEntryById)       // Get All Entrys in a List id, list_id, item_id
		api.POST("/listentry/:listid", AddEntryToList)         // Add a new Entry to an existing List
		api.PUT("/listentry/:entryid", UpdateEntryInList)      // Update Entry in an existing List
		api.DELETE("/listentry/:entryid", DeleteEntryFromList) // Delete Entry From A existing List

		// Invoices
		api.GET("/invoices/d/:departmentid", GetInvoices)                      // Get All Invoices with their name, Id, date
		api.GET("/invoices/:invoiceid", GetInvoiceById)                        // Get Invoice Using its ID
		api.POST("/invoices", AddInvoice)                                      // Add a new Invoice
		api.PUT("/invoices/:invoiceid", UpdateInvoice)                         // Update a Invoice name
		api.DELETE("/invoices/:invoiceid", DeleteInvoice)                      // Delete a Invoice
		api.GET("/invoices/notification/:invoiceid", GetInvoiceNotifications)  // Get OCR notifications for invoices
		api.DELETE("invoices/notification/:fileid", DeleteInvoiceNotification) // Deleteds Invoice notification

		// Invoice Files
		api.POST("/invoices/file/:invoiceid", InvoiceFile) // Add a new Invoice

		// Invoice Entrys
		api.GET("/invoiceentrys/:invoiceid", GetInvoiceEntrys)           // Get All Entrys in a Invoice
		api.GET("/invoiceentrys/count/:invoiceid", GetInvoiceEntryCount) // Get Count of how many Entrys in Invoice
		api.GET("/invoiceentry/:entryid", GetInvoiceEntryById)           // Get Individual Entry by entry id in a Invoice
		api.POST("/invoiceentrys/:invoiceid", AddEntryToInvoice)         // Add a new Entry to an existing Invoice
		api.PUT("/invoiceentrys/:entryid", UpdateEntryInInvoice)         // Update Entry in an existing Invoice
		api.DELETE("/invoiceentrys/:entryid", DeleteEntryFromInvoice)    // Delete Entry From A existing Invoice

		// Departments
		api.GET("/departments", GetDepartments)         // List all departments
		api.POST("/department", AddDepartment)          // Add a new department
		api.PUT("/department/:id", UpdateDepartment)    // Update an department
		api.DELETE("/department/:id", DeleteDepartment) // Delete an department

		// Search
		api.GET("/search/:q", SearchItems)
		api.GET("/search/invoice/:q", SearchInvoice)

		// Export
		api.GET("/export/csv", ExportInventoryCSV)
		api.GET("/export/pdf", ExportInventoryPDF)
		api.GET("/export/list/:listid", ExportListPDF)

		// Version
		api.GET("/version", func(c *gin.Context) {
			c.JSON(200, gin.H{"version": utils.Version, "commit": utils.Commit})
		})
	}
}
