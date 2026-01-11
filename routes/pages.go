package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoadHome serves the HTML page
func LoadHome(c *gin.Context) {
	c.HTML(http.StatusOK, "mainpage.html", nil)
}

// LoadVendor serves the HTML page
func LoadVendor(c *gin.Context) {
	c.HTML(http.StatusOK, "vendors.html", nil)
}

// LoadLocation serves the HTML page
func LoadLocation(c *gin.Context) {
	c.HTML(http.StatusOK, "locations.html", nil)
}

// LoadExport serves the HTML page
func LoadExport(c *gin.Context) {
	c.HTML(http.StatusOK, "export.html", nil)
}

// LoadBarcode serves the HTML page
func LoadBarcode(c *gin.Context) {
	c.HTML(http.StatusOK, "barcode.html", nil)
}

// LoadLists serves the HTML page
func LoadLists(c *gin.Context) {
	c.HTML(http.StatusOK, "lists.html", nil)
}
