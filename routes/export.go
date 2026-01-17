package routes

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"codeberg.org/go-pdf/fpdf"
	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"github.com/gin-gonic/gin"
)

// ExportInventoryCSV downloads all items as CSV
func ExportInventoryCSV(c *gin.Context) {
	items, err := db.GetAllItems(1, 9999999, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=inventory.csv")
	c.Header("Content-Type", "text/csv")

	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"UPC", "Invoice Number", "Brand", "Name", "Description", "Price", "Deleted", "Count", "Arrived Date", "Sold Out Date"})

	// Write rows
	for _, item := range items {
		var arrived_time, sold_time string
		if item.ArrivedAt != nil {
			arrived_time = fmt.Sprintf("%d/%d/%d", item.ArrivedAt.Local().Day(), int64(item.ArrivedAt.Local().Month()), item.ArrivedAt.Local().Year())
		}

		if item.SoldOutAt != nil {
			sold_time = fmt.Sprintf("%d/%d/%d", item.SoldOutAt.Local().Day(), int64(item.SoldOutAt.Local().Month()), item.SoldOutAt.Local().Year())
		}

		writer.Write([]string{
			item.UPC,
			item.InvoiceNumber,
			item.Brand,
			item.Name,
			item.Description,
			formatPrice(item.Price, item.Weighed),
			strconv.FormatBool(item.Deleted),
			strconv.FormatInt(int64(item.Count), 10),
			arrived_time,
			sold_time,
		})
	}
}

// ExportInventoryPDF downloads all items as PDF
func ExportInventoryPDF(c *gin.Context) {
	widths := []float64{90, 125, 125, 160, 60}

	items, err := db.GetAllItems(1, 100000000, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pdf := fpdf.New("P", "pt", "Letter", "") // 560
	pdf.AliasNbPages("")

	/* Footer with page numbers */
	pdf.SetFooterFunc(func() {
		pdf.SetY(-30) // 15mm from bottom
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d of {nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(0, 10, "Inventory List")
	pdf.Ln(30)

	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(widths[0], 18, "UPC", "1", 0, "C", false, 0, "")
	pdf.CellFormat(widths[2], 18, "Name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(widths[1], 18, "Brand", "1", 0, "C", false, 0, "")
	pdf.CellFormat(widths[3], 18, "Description", "1", 0, "C", false, 0, "")
	pdf.CellFormat(widths[4], 18, "Price", "1", 0, "C", false, 0, "")

	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10)

	for _, item := range items {
		var price string
		if item.Weighed {
			price = fmt.Sprintf("$%.2f/lb", item.Price)
		} else {
			price = fmt.Sprintf("$%.2f", item.Price)
		}

		if item.Deleted {
			pdf.SetTextColor(150, 150, 150)
		} else {
			pdf.SetTextColor(0, 0, 0)
		}

		lineHeight := 12.0

		texts := []string{
			item.UPC,
			item.Name,
			item.Brand,
			item.Description,
			price,
		}

		rowHeight := RowHeight(pdf, lineHeight, widths, texts)

		// Page break protection
		_, h := pdf.GetPageSize()
		if pdf.GetY()+rowHeight > h-60 {
			pdf.AddPage()
		}

		x, y := pdf.GetXY()

		// Draw borders FIRST (important)
		curX := x
		for _, w := range widths {
			pdf.Rect(curX, y, w, rowHeight, "")
			curX += w
		}

		// Write text
		pdf.MultiCell(widths[0], lineHeight, texts[0], "", "L", false)
		pdf.SetXY(x+widths[0], y)

		pdf.MultiCell(widths[1], lineHeight, texts[1], "", "L", false)
		pdf.SetXY(x+widths[0]+widths[1], y)

		pdf.MultiCell(widths[2], lineHeight, texts[2], "", "L", false)
		pdf.SetXY(x+widths[0]+widths[1]+widths[2], y)

		pdf.MultiCell(widths[3], lineHeight, texts[3], "", "L", false)
		pdf.SetXY(x+widths[0]+widths[1]+widths[2]+widths[3], y)

		pdf.CellFormat(widths[4], rowHeight, texts[4], "", 0, "R", false, 0, "")

		pdf.Ln(-1)
	}

	c.Header("Content-Disposition", "attachment; filename=inventory.pdf")
	c.Header("Content-Type", "application/pdf")
	err = pdf.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// ExportListPDF downloads all items in a list as PDF
func ExportListPDF(c *gin.Context) {
	listID, err := strconv.Atoi(c.Param("listid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid entry id"})
		return
	}

	// Get the list name and date information
	list, err := db.GetListById(listID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the list name and date information
	list_data, err := db.GetListEntrys(listID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// setup colum sizes on pdf
	widths := []float64{130, 200, 155, 70}

	// Setup pdf page
	pdf := fpdf.New("P", "pt", "Letter", "") // 560
	pdf.AliasNbPages("")

	/* Footer with page numbers */
	pdf.SetFooterFunc(func() {
		pdf.SetY(-30) // 15mm from bottom
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d of {nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})

	// Set up page title
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(0, 10, "List Name: "+list.ListName)
	pdf.Ln(20)

	// setup created and printed date
	pdf.SetFont("Arial", "", 15)
	pdf.Cell(180, 10, "Created: "+list.CreatedAt.Local().Format("01/02/2006"))

	pdf.Cell(0, 10, "Printed: "+time.Now().Format("01/02/2006"))
	pdf.Ln(30)

	// Create collum names
	pdf.SetFont("Arial", "B", 18)
	pdf.CellFormat(widths[0], 20, "Location", "1", 0, "C", false, 0, "")
	pdf.CellFormat(widths[1], 20, "Name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(widths[2], 20, "Brand", "1", 0, "C", false, 0, "")
	pdf.CellFormat(widths[3], 20, "Price", "1", 0, "C", false, 0, "")

	// Move to the next line and set font type and size
	pdf.Ln(-1)
	pdf.SetFont("Arial", "", 16)

	for _, list_datum := range list_data {
		var price string
		if list_datum.Weighed {
			price = fmt.Sprintf("$%.2f/lb", list_datum.Price)
		} else {
			price = fmt.Sprintf("$%.2f", list_datum.Price)
		}

		if list_datum.Deleted {
			pdf.SetTextColor(150, 150, 150)
		} else {
			pdf.SetTextColor(0, 0, 0)
		}
		lineHeight := 19.0

		texts := []string{
			list_datum.LocationName,
			list_datum.Name,
			list_datum.Brand,
			price,
		}

		rowHeight := RowHeight(pdf, lineHeight, widths, texts)

		// Page break protection
		_, h := pdf.GetPageSize()
		if pdf.GetY()+rowHeight > h-60 {
			pdf.AddPage()
		}

		x, y := pdf.GetXY()

		// Draw borders FIRST (important)
		curX := x
		for _, w := range widths {
			pdf.Rect(curX, y, w, rowHeight, "")
			curX += w
		}

		// Write text
		pdf.MultiCell(widths[0], lineHeight, texts[0], "", "L", false)
		pdf.SetXY(x+widths[0], y)

		pdf.MultiCell(widths[1], lineHeight, texts[1], "", "L", false)
		pdf.SetXY(x+widths[0]+widths[1], y)

		pdf.MultiCell(widths[2], lineHeight, texts[2], "", "L", false)
		pdf.SetXY(x+widths[0]+widths[1]+widths[2], y)

		pdf.CellFormat(widths[3], rowHeight, texts[3], "", 0, "R", false, 0, "")

		pdf.Ln(-1)
	}

	// Setup download Type
	c.Header("Content-Disposition", "attachment; filename="+list.ListName+".pdf")
	c.Header("Content-Type", "application/pdf")
	err = pdf.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func RowHeight(pdf *fpdf.Fpdf, lineHeight float64, widths []float64, texts []string) float64 {
	max := 0.0
	for i, txt := range texts {
		lines := pdf.SplitLines([]byte(txt), widths[i])
		h := float64(len(lines)) * lineHeight
		if h > max {
			max = h
		}
	}
	return max
}

func formatPrice(price float64, weighed bool) string {
	if weighed {
		return fmt.Sprintf("$%.2f/lb", price)
	} else {
		return fmt.Sprintf("$%.2f", price)
	}
}
