package models

type Item struct {
	ID          int     `json:"id"`
	UPC         string  `json:"upc"`
	Brand       string  `json:"brand"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
