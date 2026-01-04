package models

import (
	"time"
)

type Inventory struct {
	ID            int        `json:"item_id"`
	VendorID      int        `json:"vendor_id"`
	UPC           string     `json:"upc"`
	InvoiceNumber string     `json:"invoice_number"`
	Brand         string     `json:"brand"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Price         float64    `json:"price"`
	Weighed       bool       `json:"weighed"`
	Count         int        `json:"count"`
	Deleted       bool       `json:"deleted"`
	ArrivedAt     *time.Time `json:"arrived_at"`
	SoldOutAt     *time.Time `json:"sold_out_at"`
}

type InventoryAll struct {
	ID            int        `json:"item_id"`
	VendorID      int        `json:"vendor_id"`
	VendorName    string     `json:"vendor_name"`
	UPC           string     `json:"upc"`
	InvoiceNumber string     `json:"invoice_number"`
	Brand         string     `json:"brand"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Price         float64    `json:"price"`
	Weighed       bool       `json:"weighed"`
	Count         int        `json:"count"`
	Deleted       bool       `json:"deleted"`
	ArrivedAt     *time.Time `json:"arrived_at"`
	SoldOutAt     *time.Time `json:"sold_out_at"`
	Entry_Count   int        `json:"number_of_entrys"`
}

type Vendor struct {
	VendorID   int    `json:"vendor_id"`
	VendorName string `json:"vendor_name"`
}
