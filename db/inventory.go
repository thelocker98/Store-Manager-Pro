package db

import (
	"fmt"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddItem(item models.Inventory) error {
	query := `
	INSERT INTO inventory (catalog_id, vendor_id, description, price, weighed, count)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := DB.Exec(query, item.CatalogID, item.VendorID, item.Description, item.Price, item.Weighed, item.Count)
	fmt.Println(err)
	return err
}

func GetAllItems() ([]models.InventoryAll, error) {
	rows, err := DB.Query(`
	SELECT
		i.id,
		c.catalog_id,
		v.vendor_id,
		c.upc,
		c.invoice_number,
		c.brand,
		c.name,
		i.description,
		i.price,
		i.weighed,
		i.count,
		i.arived_at,
		i.soldout_at,
		v.vendor_name
	FROM inventory i
	JOIN catalog c ON i.catalog_id = c.catalog_id
	LEFT JOIN vendors v ON i.vendor_id = v.vendor_id
	ORDER BY c.upc ASC;
	`)

	if err != nil {
		fmt.Println("first", err)
		return nil, err
	}
	defer rows.Close()

	var items []models.InventoryAll
	for rows.Next() {
		var i models.InventoryAll
		err := rows.Scan(
			&i.ID,
			&i.CatalogID,
			&i.VendorID,
			&i.UPC,
			&i.InvoiceNumber,
			&i.Brand,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Weighed,
			&i.Count,
			&i.ArrivedAt,
			&i.SoldOutAt,
			&i.VendorName,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func DeleteItem(id int) error {
	_, err := DB.Exec(`DELETE FROM inventory WHERE id = ?`, id)
	return err
}

func UpdateItem(id int, item models.Inventory) error {
	query := `
	UPDATE inventory
	SET catalog_id = ?, vendor_id = ?, description = ?, price = ?, weighed = ?, count = ?, ArrivedAt = ?, SoldAt = ?
	WHERE id = ?
	`
	_, err := DB.Exec(query, item.CatalogID, item.VendorID, item.Description, item.Price, item.Weighed, item.Count, item.ArrivedAt, item.SoldAt, item.ID)
	return err
}
