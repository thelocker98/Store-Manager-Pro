package db

import (
	"fmt"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddItem(item models.Inventory) error {
	query := `
	INSERT INTO inventory (vendor_id, upc, invoice_number, name, brand, description, price, weighed, count)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := DB.Exec(query, item.VendorID, item.UPC, item.InvoiceNumber, item.Name, item.Brand, item.Description, item.Price, item.Weighed, item.Count)
	return err
}

func GetAllItems() ([]models.InventoryAll, error) {
	rows, err := DB.Query(`
	SELECT
		i.id,
		v.vendor_id,
		v.vendor_name,
		i.upc,
		i.invoice_number,
		i.brand,
		i.name,
		i.description,
		i.price,
		i.weighed,
		i.count,
		i.deleted,
		i.arrived_at,
		i.soldout_at
	FROM inventory i
	JOIN vendors v ON i.vendor_id = v.vendor_id
	ORDER BY i.upc ASC;
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
			&i.VendorID,
			&i.VendorName,
			&i.UPC,
			&i.InvoiceNumber,
			&i.Brand,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Weighed,
			&i.Count,
			&i.Deleted,
			&i.ArrivedAt,
			&i.SoldOutAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func UpdateItem(id int, item models.Inventory) error {
	query := `
	UPDATE inventory
	SET vendor_id = ?, upc = ?, invoice_number = ?, brand = ?, name = ?, description = ?, price = ?, weighed = ?, count = ?, deleted = ?, arrived_at = ?, soldout_at = ?
	WHERE id = ?
	`
	_, err := DB.Exec(query, item.VendorID, item.UPC, item.InvoiceNumber, item.Brand, item.Name, item.Description, item.Price, item.Weighed, item.Count, item.Deleted, item.ArrivedAt, item.SoldOutAt, item.ID)
	return err
}

func DeleteItem(id int) error {
	query := `
	UPDATE inventory
	SET deleted = true, soldout_at = CURRENT_TIMESTAMP
	WHERE id = ?
	`
	_, err := DB.Exec(query, id)
	return err
}

func RestoreItem(id int) error {
	query := `
	UPDATE inventory
	SET deleted = false, soldout_at = NULL
	WHERE id = ?
	`
	_, err := DB.Exec(query, id)
	return err
}

func DeleteItemPermanent(id int) error {
	_, err := DB.Exec(`DELETE FROM inventory WHERE id = ?`, id)
	return err
}
