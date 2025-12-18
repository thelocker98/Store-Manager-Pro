package models

type Item struct {
	ID          int     `json:"id"`
	UPC         string  `json:"upc"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
