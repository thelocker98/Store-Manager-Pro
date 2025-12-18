package routes

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"gitea.locker98.com/locker98/Store-Manager-Pro/models"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf" // PDF library, install: go get github.com/jung-kurt/gofpdf
)

// LoadHome serves the HTML page
func LoadHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// GetItems returns all items in JSON
func GetItems(c *gin.Context) {
	items, err := db.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// AddItem adds a new item
func AddItem(c *gin.Context) {
	var item models.Item
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

	var item models.Item
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

// ExportCSV downloads all items as CSV
func ExportCSV(c *gin.Context) {
	items, err := db.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=inventory.csv")
	c.Header("Content-Type", "text/csv")

	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"ID", "UPC", "Name", "Description", "Price"})

	// Write rows
	for _, item := range items {
		writer.Write([]string{
			strconv.Itoa(item.ID),
			item.UPC,
			item.Name,
			item.Description,
			fmt.Sprintf("%.2f", item.Price),
		})
	}
}

// ExportPDF downloads all items as PDF
func ExportPDF(c *gin.Context) {
	items, err := db.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Inventory List")
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(10, 8, "ID", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "UPC", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 8, "Name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(60, 8, "Description", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "Price", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 12)
	for _, item := range items {
		pdf.CellFormat(10, 8, strconv.Itoa(item.ID), "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 8, item.UPC, "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 8, item.Name, "1", 0, "", false, 0, "")
		// MultiCell for description
		x := pdf.GetX()
		y := pdf.GetY()
		pdf.MultiCell(60, 6, item.Description, "1", "", false)
		pdf.SetXY(x+60, y) // move X to next cell
		pdf.CellFormat(30, 8, fmt.Sprintf("$%.2f", item.Price), "1", 0, "R", false, 0, "")
		pdf.Ln(-1)
	}

	c.Header("Content-Disposition", "attachment; filename=inventory.pdf")
	c.Header("Content-Type", "application/pdf")
	err = pdf.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
