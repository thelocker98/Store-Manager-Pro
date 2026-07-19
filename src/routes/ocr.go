package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
	"gitea.locker98.com/locker98/Store-Manager-Pro/ocr"
	"gitea.locker98.com/locker98/Store-Manager-Pro/utils"
	"github.com/gin-gonic/gin"
)

// Manage File Uploads
func InvoiceFile(c *gin.Context) {
	invoiceID, err := strconv.Atoi(c.Param("invoiceid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice id"})
		return
	}

	// Get file from form data
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	// Get markup amount from form data
	markupStr := c.DefaultPostForm("markup", "30")
	markupAmount, err := strconv.Atoi(markupStr)

	if err != nil || markupAmount < 0 || markupAmount > 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Markup"})
		return
	}

	// Save file to ./temp
	timestamp := time.Now().UnixNano()
	filePath := fmt.Sprintf("%s/%d_%d_%s", utils.InvoiceFolderPath, invoiceID, timestamp, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	fileId, err := db.CreateOCREntry(invoiceID, filePath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "failed to add to db: " + err.Error()})
		return
	}

	job := models.OCRJob{
		FileID:       fileId,
		InvoiceID:    invoiceID,
		FilePath:     filePath,
		MarkupAmount: markupAmount,
	}

	// Send to channel
	select {
	case ocr.OCRQueue <- job:
		c.JSON(http.StatusOK, gin.H{"message": "file uploaded successfully"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "file failed to enqueue"})
	}
}

// Get OCR notifications for invoices
func GetInvoiceNotifications(c *gin.Context) {
	invoiceID, err := strconv.Atoi(c.Param("invoiceid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice id"})
		return
	}
	ocrEntrys, err := db.GetOCREntrysByInvoiceID(invoiceID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ocrEntrys)

}

// Deleteds Invoice notification
func DeleteInvoiceNotification(c *gin.Context) {
	fileID, err := strconv.Atoi(c.Param("fileid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.DeleteOCREntryById(fileID, false)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
