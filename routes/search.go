package routes

import (
	"net/http"
	"strconv"

	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"github.com/gin-gonic/gin"
)

func SearchItems(c *gin.Context) {
	// Search Settings
	searchTerm := c.Param("q")
	showDeletedStr, _ := strconv.ParseBool(c.DefaultQuery("showdeleted", "false")) // default to false
	sortBy := c.DefaultQuery("sortby", "upc")                                      // default to upc

	// Page Settings
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "50"))

	items, err := db.SearchItems(searchTerm, sortBy, showDeletedStr, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
