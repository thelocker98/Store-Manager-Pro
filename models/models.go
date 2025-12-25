package models

import (
	"time"
)

type Inventory struct {
	ID          int        `json:"id"`
	CatalogID   int        `json:"catalog_id"`
	VendorID    int        `json:"vendor_id"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Weighed     bool       `json:"weighed" `
	Count       int        `json:"count"`
	ArrivedAt   *time.Time `json:"arrived_at"`
	SoldAt      *time.Time `json:"sold_at,omitempty"`
}

type InventoryAll struct {
	ID            int        `json:"item_id"`
	CatalogID     int        `json:"catalog_id"`
	VendorID      int        `json:"vendor_id"`
	UPC           string     `json:"upc"`
	InvoiceNumber string     `json:"invoice_number"`
	Brand         string     `json:"brand"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Price         float64    `json:"price"`
	Weighed       bool       `json:"weighed"`
	Count         int        `json:"count"`
	ArrivedAt     *time.Time `json:"arrived_at"`
	SoldOutAt     *time.Time `json:"sold_out_at"`
	VendorName    string     `json:"vendor_name"`
}

type Catalog struct {
	CatalogID     int    `json:"catalog_id"`
	UPC           string `json:"upc"`
	InvoiceNumber string `json:"invoice_number"`
	Brand         string `json:"brand"`
	Name          string `json:"name"`
}

type Vendor struct {
	VendorID   int    `json:"vendor_id"`
	VendorName string `json:"vendor_name"`
}
