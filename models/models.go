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
	DepartmentID      int    `json:"department_id"`
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
