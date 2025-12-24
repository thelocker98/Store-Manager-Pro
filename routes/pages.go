package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoadHome serves the HTML page
func LoadHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// LoadVendor serves the HTML page
func LoadVendor(c *gin.Context) {
	c.HTML(http.StatusOK, "vendors.html", nil)
}

// LoadItemEdit serves the HTML page for editing items and creating them
func LoadItemEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "edititems.html", nil)
}

// LoadCatalog serves the HTML page for editing and adding catalog entrys
func LoadCatalog(c *gin.Context) {
	c.HTML(http.StatusOK, "catalog.html", nil)
}
