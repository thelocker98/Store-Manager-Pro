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

// GetLists returns all lists in JSON
func GetLists(c *gin.Context) {
	departmentID, err := strconv.Atoi(c.Param("departmentid"))

	lists, err := db.GetLists(departmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lists)
}

// Get Single List
func GetListById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("listid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid list id"})
		return
	}
	lists, err := db.GetListById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no list found"})
		return
	}
	c.JSON(http.StatusOK, lists)
}

func GetNumberOfItemsList(c *gin.Context) {
	listID, err := strconv.Atoi(c.Param("listid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid list id"})
		return
	}
	totalCount, err := db.GetNumberOfItemsList(listID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no list found"})
		return
	}
	c.JSON(http.StatusOK, map[string]int{"total_count": totalCount})
}

// AddList adds a new list
func AddList(c *gin.Context) {
	var list models.List
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	if err := db.AddList(list); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "List added"})
}

// UpdateList updates an existing list by ID
func UpdateList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("listid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid list id"})
		return
	}

	var list models.List
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.UpdateList(id, list); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "list updated"})
}

// DeleteList deletes an list by ID
func DeleteList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("listid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid list id"})
		return
	}

	if err := db.DeleteList(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "List deleted"})
}

// GetListEntrys returns all entrys in a list
func GetListEntrys(c *gin.Context) {
	listID, err := strconv.Atoi(c.Param("listid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid list id"})
		return
	}
	lists, err := db.GetListEntrys(listID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lists)
}

// GetListEntryById returns all entrys in a list
func GetListEntryById(c *gin.Context) {
	entryID, err := strconv.Atoi(c.Param("entryid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid list entry id"})
		return
	}
	lists, err := db.GetListEntryById(entryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lists)
}

// AddEntryToList adds a new list
func AddEntryToList(c *gin.Context) {
	listID, err := strconv.Atoi(c.Param("listid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid list id"})
		return
	}

	var listEntry models.ListEntry
	if err := c.ShouldBindJSON(&listEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	listEntry.ListID = listID

	entryId, err := db.AddEntryToList(listEntry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "List entry added", "entry_id": entryId})
}

// UpdateEntryInList deletes an entry from a list by ID
func UpdateEntryInList(c *gin.Context) {
	entryID, err := strconv.Atoi(c.Param("entryid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid entry id"})
		return
	}

	var list models.ListEntry
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.UpdateEntryInList(entryID, list); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "entry updated"})
}

// DeleteEntryFromList deletes an entry from a list by ID
func DeleteEntryFromList(c *gin.Context) {
	entryID, err := strconv.Atoi(c.Param("entryid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid entry id"})
		return
	}

	if err := db.DeleteEntryFromList(entryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "entry deleted from list"})
}
