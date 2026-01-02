package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"gitea.locker98.com/locker98/Store-Manager-Pro/models"

	"github.com/gin-gonic/gin"
	// PDF library, install: go get github.com/jung-kurt/gofpdf
)

// GetItems returns all items in JSON
func GetItems(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "50"))

	items, err := db.GetAllItems(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetItems returns all items in JSON
func CountItems(c *gin.Context) {
	showDeleted, _ := strconv.ParseBool(c.DefaultQuery("countdeleted", "false")) // default to false
	count, err := db.CountAllItems(showDeleted)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, `{"count":`+fmt.Sprint(count)+`}`)
}

// Get Single Item
func GetItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid item ID"})
		return
	}
	items, err := db.GetItemById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no rows found"})
		return
	}
	c.JSON(http.StatusOK, items)
}

// AddItem adds a new item
func AddItem(c *gin.Context) {
	var item models.Inventory
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.AddItem(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item added"})
}

// UpdateItem updates an existing item by ID
func UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var item models.Inventory
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.UpdateItem(id, item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated"})
}

// DeleteItem deletes an item by ID
func DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.DeleteItem(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}

// DeleteItem deletes an item by ID
func RestoreItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.RestoreItem(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item restored"})
}

// DeleteItem deletes an item by ID
func DeleteItemPermanent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.DeleteItemPermanent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted permanently"})
}
