package routes

import (
	"net/http"
	"strconv"

	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"gitea.locker98.com/locker98/Store-Manager-Pro/models"

	"github.com/gin-gonic/gin"
	// PDF library, install: go get github.com/jung-kurt/gofpdf
)

// GetInvoices returns all invoices in JSON
func GetInvoices(c *gin.Context) {
	departmentID, err := strconv.Atoi(c.Param("departmentid"))

	invoices, err := db.GetInvoices(departmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, invoices)
}

// Get Single Invoice
func GetInvoiceById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("invoiceid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid invoice id"})
		return
	}
	invoice, err := db.GetInvoiceById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no invoice found"})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

// AddInvoice adds a new invoice to the db
func AddInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.AddInvoice(invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Invoice added"})
}

// UpdateInvoice updates an existing Invoice by ID
func UpdateInvoice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("invoiceid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice id"})
		return
	}

	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.UpdateInvoice(id, invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invoice updated"})
}

// DeleteInvoice deletes an invoice by ID
func DeleteInvoice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("invoiceid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice id"})
		return
	}

	if err := db.DeleteInvoice(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted"})
}

// GetInvoiceEntrys returns all entrys in a list
func GetInvoiceEntrys(c *gin.Context) {
	invoiceID, err := strconv.Atoi(c.Param("invoiceid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice id"})
		return
	}
	invoices, err := db.GetInvoiceEntrys(invoiceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, invoices)
}

// GetInvoiceEntryById returns an entrys in a invoice by id
func GetInvoiceEntryById(c *gin.Context) {
	entryID, err := strconv.Atoi(c.Param("entryid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice entry id"})
		return
	}
	invoiceEntry, err := db.GetInvoiceEntryById(entryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, invoiceEntry)
}

// AddEntryToInvoice adds a new entry to an invoice with a given id
func AddEntryToInvoice(c *gin.Context) {
	invoiceID, err := strconv.Atoi(c.Param("invoiceid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid invoice id"})
		return
	}

	var invoiceEntry models.InvoiceEntry
	if err := c.ShouldBindJSON(&invoiceEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	invoiceEntry.InvoiceDataID = invoiceID

	err = db.AddEntryToInvoice(invoiceEntry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Invoice entry added"})
}

// UpdateEntryInInvoice deletes an entry from a invoice by ID
func UpdateEntryInInvoice(c *gin.Context) {
	entryID, err := strconv.Atoi(c.Param("entryid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid entry id"})
		return
	}

	var invoiceEntry models.InvoiceEntry
	if err := c.ShouldBindJSON(&invoiceEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.UpdateEntryInInvoice(entryID, invoiceEntry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "entry updated"})
}

// DeleteEntryFromInvoice deletes an entry from a invoice by ID
func DeleteEntryFromInvoice(c *gin.Context) {
	entryID, err := strconv.Atoi(c.Param("entryid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid entry id"})
		return
	}

	if err := db.DeleteEntryFromInvoice(entryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "entry deleted from invoice"})
}
