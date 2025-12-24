package db

import (
	"errors"
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
	rows, err := DB.Query(`SELECT catalog_id, upc, invoice_number, brand, name FROM catalog ORDER BY upc ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println(rows)

	var catalog []models.Catalog
	for rows.Next() {
		var c models.Catalog
		rows.Scan(&c.CatalogID, &c.UPC, &c.InvoiceNumber, &c.Brand, &c.Name)
		fmt.Println(c)
		catalog = append(catalog, c)
	}
	return catalog, nil
}

func DeleteCatalog(id int) error {
	s, err := DB.Exec(`DELETE FROM catalog WHERE catalog_id = ?`, id)
	if err != nil {
		return err
	}
	if r, _ := s.RowsAffected(); r == 0 {
		return errors.New("catalog entry does not exist")
	}
	return err
}

func UpdateCatalog(catalog models.Catalog) error {
	query := `
	UPDATE catalog
	SET upc = ?, invoice_number = ?, brand = ?, name = ?
	WHERE catalog_id = ?
	`
	s, err := DB.Exec(query, catalog.UPC, catalog.InvoiceNumber, catalog.Brand, catalog.Name, catalog.CatalogID)
	if err != nil {
		return err
	}
	if r, _ := s.RowsAffected(); r == 0 {
		return errors.New("catalog entry does not exist")
	}
	return nil
}
