package routes

import (
	"net/http"
	"strconv"

	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
	"github.com/gin-gonic/gin"
)

// GetCatalog returns all items in JSON
func GetCatalog(c *gin.Context) {
	catalog, err := db.GetAllVendors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, catalog)
}

// AddCatalogEntry adds a new item
func AddCatalogEntry(c *gin.Context) {
	var catalog models.Catalog
	if err := c.ShouldBindJSON(&catalog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.AddCatalog(catalog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Catalog entry added"})
}

// UpdateCatalogEntry updates an existing item by ID
func UpdateCatalogEntry(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid catalog ID"})
		return
	}

	var catalog models.Catalog
	if err := c.ShouldBindJSON(&catalog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	catalog.CatalogID = id

	if err := db.UpdateCatalog(catalog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Catalog entry updated"})
}

// DeleteCatalogEntry deletes an item by ID
func DeleteCatalogEntry(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid catalog ID"})
		return
	}

	if err := db.DeleteCatalog(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Catalog entry deleted"})
}
