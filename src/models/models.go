package models

import (
	"time"
)

type Inventory struct {
	ItemID         int        `json:"item_id"`
	VendorID       int        `json:"vendor_id"`
	LocationID     int        `json:"location_id"`
	DepartmentID   int        `json:"department_id"`
	DepartmentName string     `json:"department_name"`
	UPC            string     `json:"upc"`
	InvoiceNumber  string     `json:"invoice_number"`
	Brand          string     `json:"brand"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Price          float64    `json:"price"`
	Weighed        bool       `json:"weighed"`
	Count          int        `json:"count"`
	Deleted        bool       `json:"deleted"`
	ArrivedAt      *time.Time `json:"arrived_at"`
	SoldOutAt      *time.Time `json:"sold_out_at"`
}

type InventoryAll struct {
	ItemID         int        `json:"item_id"`
	VendorID       int        `json:"vendor_id"`
	VendorName     string     `json:"vendor_name"`
	LocationID     int        `json:"location_id"`
	LocationName   string     `json:"location_name"`
	DepartmentID   int        `json:"department_id"`
	DepartmentName string     `json:"department_name"`
	UPC            string     `json:"upc"`
	InvoiceNumber  string     `json:"invoice_number"`
	Brand          string     `json:"brand"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Price          float64    `json:"price"`
	Weighed        bool       `json:"weighed"`
	Count          int        `json:"count"`
	Deleted        bool       `json:"deleted"`
	ArrivedAt      *time.Time `json:"arrived_at"`
	SoldOutAt      *time.Time `json:"sold_out_at"`
	EntryID        int        `json:"entry_id"`
	ListCount      int        `json:"list_count"`
	Entry_Count    int        `json:"number_of_entrys"`
}

type Vendor struct {
	VendorID        int    `json:"vendor_id"`
	VendorName      string `json:"vendor_name"`
	VendorItemCount int    `json:"vendor_item_count"`
}

type Location struct {
	LocationID        int    `json:"location_id"`
	LocationName      string `json:"location_name"`
	LocationItemCount int    `json:"location_item_count"`
}

type Department struct {
	DepartmentID        int    `json:"department_id"`
	DepartmentName      string `json:"department_name"`
	DepartmentItemCount int    `json:"department_item_count"`
}

type List struct {
	ListID         int        `json:"list_id"`
	ListName       string     `json:"list_name"`
	DepartmentID   int        `json:"department_id"`
	DepartmentName string     `json:"department_name"`
	TotalCount     int        `json:"total_count"`
	CreatedAt      *time.Time `json:"created_at"`
}

type ListEntry struct {
	EntryID   int `json:"entry_id"`
	ListID    int `json:"list_id"`
	ListCount int `json:"list_count"`
	ItemID    int `json:"item_id"`
}

type ListEntryAll struct {
	EntryID       int     `json:"entry_id"`
	ListID        int     `json:"list_id"`
	ItemID        int     `json:"item_id"`
	VendorID      int     `json:"vendor_id"`
	LocationID    int     `json:"location_id"`
	ListCount     int     `json:"list_count"`
	VendorName    string  `json:"vendor_name"`
	LocationName  string  `json:"location_name"`
	UPC           string  `json:"upc"`
	InvoiceNumber string  `json:"invoice_number"`
	Brand         string  `json:"brand"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	Weighed       bool    `json:"weighed"`
	Count         int     `json:"count"`
	Deleted       bool    `json:"deleted"`
}

type Invoice struct {
	InvoiceID      int        `json:"invoice_id"`
	InvoiceName    string     `json:"invoice_name"`
	InvoiceType    string     `json:"invoice_type"`
	DepartmentID   int        `json:"department_id"`
	DepartmentName string     `json:"department_name"`
	TotalCount     int        `json:"total_count"`
	CreatedAt      *time.Time `json:"created_at"`
}

type InvoiceEntry struct {
	InvoiceDataID int     `json:"invoice_entry_id"`
	InvoiceID     int     `json:"invoice_id"`
	RawText       string  `json:"rawtext"`
	UPC           string  `json:"upc"`
	Detail        string  `json:"details"`
	QTY           int     `json:"qty"`
	ItemCost      float64 `json:"item_cost"`     // Cost of Item from retailer like Walmart
	TotalCost     float64 `json:"total_cost"`    // Total cost of all items in the row from retailer like Walmart
	DiscountCost  float64 `json:"discount_cost"` // Total price paid for the row by the store to distributor
	TrueCost      float64 `json:"true_cost"`     // Price to sell item for in the store
	EntryCount    int     `json:"number_of_entrys"`
}

type OCREntry struct {
	FileID        int        `json:"file_id"`
	InvoiceID     int        `json:"invoice_id"`
	FilePath      string     `json:"file_path"`
	Status        int        `json:"status"`
	EntriesAdded  int        `json:"entries_added"`
	EntriesFailed int        `json:"entries_failed"`
	ErrorMessage  string     `json:"error_message"`
	StartedAt     *time.Time `json:"started_at"`
	CompletedAt   *time.Time `json:"completed_at"`
	CreatedAt     *time.Time `json:"created_at"`
}

type OCRJob struct {
	FileID       int
	InvoiceID    int
	FilePath     string
	MarkupAmount int
}

type OCRStatus int

const (
	Pending OCRStatus = iota
	Processing
	Failed
	Succeeded
)

func (c OCRStatus) String() string {
	return [...]string{"Pending", "Processing", "Failed", "Succeded"}[c]
}
