package routes

import (
	"net/http"
	"strconv"

	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"github.com/gin-gonic/gin"
)

func SearchItems(c *gin.Context) {
	limitToList := false
	// Search Settings
	searchTerm := c.Param("q")

	// Page Settings
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "50"))
	listID, _ := strconv.Atoi(c.DefaultQuery("list_id", "0"))

	// sort
	showDeletedStr, _ := strconv.ParseBool(c.DefaultQuery("showdeleted", "false")) // default to false
	vendorID, _ := strconv.Atoi(c.DefaultQuery("vendor", "0"))
	locationID, _ := strconv.Atoi(c.DefaultQuery("location", "0"))
	departmentID, _ := strconv.Atoi(c.DefaultQuery("department", "0"))
	filterListID, _ := strconv.Atoi(c.DefaultQuery("filterlist_id", "0"))
	orderBy := c.DefaultQuery("orderby", "upc") // default to upc

	if filterListID != 0 {
		limitToList = true
		listID = filterListID
	}

	items, err := db.SearchItems(searchTerm, listID, page, pageSize, showDeletedStr, vendorID, locationID, departmentID, limitToList, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func SearchInvoice(c *gin.Context) {
	// Search Settings
	searchTerm := c.Param("q")

	// Page Settings
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "50"))
	invoiceID, _ := strconv.Atoi(c.DefaultQuery("invoice_id", "0"))

	entrys, err := db.SearchInvoice(searchTerm, invoiceID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entrys)
}
