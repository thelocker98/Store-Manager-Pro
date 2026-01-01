package routes

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

// ExportCSV downloads all items as CSV
func ExportCSV(c *gin.Context) {
	items, err := db.GetAllItems(1, 5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=inventory.csv")
	c.Header("Content-Type", "text/csv")

	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"UPC", "Brand", "Name", "Description", "Price"})

	// Write rows
	for _, item := range items {
		writer.Write([]string{
			item.UPC,
			item.Brand,
			item.Name,
			item.Description,
			fmt.Sprintf("%.2f", item.Price),
		})
	}
}

// ExportPDF downloads all items as PDF
func ExportPDF(c *gin.Context) {
	items, err := db.GetAllItems(1,5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AliasNbPages("")

	/* Footer with page numbers */
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15) // 15mm from bottom
		pdf.SetFont("Arial", "", 9)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d of {nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 18)
	pdf.Cell(0, 10, "Inventory List")
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(25, 8, "UPC", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 8, "Brand", "1", 0, "C", false, 0, "")
	pdf.CellFormat(55, 8, "Name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(55, 8, "Description", "1", 0, "C", false, 0, "")
	pdf.CellFormat(16, 8, "Price", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10)
	for _, item := range items {
		pdf.CellFormat(25, 8, item.UPC, "1", 0, "L", false, 0, "")
		pdf.CellFormat(40, 8, item.Brand, "1", 0, "L", false, 0, "")
		pdf.CellFormat(55, 8, item.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(55, 8, item.Description, "1", 0, "L", false, 0, "")
		pdf.CellFormat(16, 8, fmt.Sprintf("$%.2f", item.Price), "1", 0, "R", false, 0, "")

		pdf.Ln(-1)
	}

	c.Header("Content-Disposition", "attachment; filename=inventory.pdf")
	c.Header("Content-Type", "application/pdf")
	err = pdf.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
