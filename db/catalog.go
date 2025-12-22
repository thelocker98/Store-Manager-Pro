package db

import (
	"fmt"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddCatalog(catalog models.Catalog) error {
	query := `
	INSERT INTO catalog (upc, invoice_number, brand, name)
	VALUES (?, ?, ?, ?)
	`
	_, err := DB.Exec(query, catalog.UPC, catalog.InvoiceNumber, catalog.Brand, catalog.Name)
	fmt.Println(err)
	return err
}

func GetAllCatalog() ([]models.Catalog, error) {
	rows, err := DB.Query(`SELECT id, upc, brand, name, description, price FROM catalog ORDER BY upc ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var catalog []models.Catalog
	for rows.Next() {
		var c models.Catalog
		rows.Scan(&c.CatalogID, &c.UPC, &c.InvoiceNumber, &c.Brand, &c.Name)
		catalog = append(catalog, c)
	}
	return catalog, nil
}

func DeleteCatalog(id int) error {
	_, err := DB.Exec(`DELETE FROM catalog WHERE catalog_id = ?`, id)
	return err
}

func UpdateCatalog(catalog models.Catalog) error {
	query := `
	UPDATE catalog
	SET upc = ?, invoice_number = ?, brand = ?, name = ?, description = ?, price = ?
	WHERE id = ?
	`
	_, err := DB.Exec(query, catalog.UPC, catalog.InvoiceNumber, catalog.Brand, catalog.Name, catalog.CatalogID)
	return err
}
