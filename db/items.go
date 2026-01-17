package db

import (
	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddItem(item models.Inventory) error {
	query := `
	INSERT INTO inventory (vendor_id, location_id, upc, invoice_number, name, brand, description, price, weighed, count)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	_, err := DB.Exec(query, item.VendorID, item.LocationID, item.UPC, item.InvoiceNumber, item.Name, item.Brand, item.Description, item.Price, item.Weighed, item.Count)
	return err
}

func GetAllItems(page int, pageSize int, showDeleted bool) ([]models.InventoryAll, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 50
	}
	offset := (page - 1) * pageSize

	query := `
	SELECT
		i.id,
		v.vendor_id,
		v.vendor_name,
		l.location_id,
		l.location_name,
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
		i.soldout_at,
		COUNT(*) OVER() AS entry_count
	FROM inventory i
	JOIN vendors v ON i.vendor_id = v.vendor_id
	JOIN locations l ON i.location_id = l.location_id
	`
	if !showDeleted {
		query += ` WHERE i.deleted = 0 `
	}
	query += `
		ORDER BY i.upc ASC
		LIMIT ? OFFSET ?;
	`

	rows, err := DB.Query(query, pageSize, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.InventoryAll
	for rows.Next() {
		var i models.InventoryAll
		err := rows.Scan(
			&i.ItemID,
			&i.VendorID,
			&i.VendorName,
			&i.LocationID,
			&i.LocationName,
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
			&i.Entry_Count,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func CountAllItems(countDeleted bool) (int, error) {
	query := `
	SELECT COUNT(id)
	FROM inventory
	`
	if !countDeleted {
		query += " WHERE deleted = 0;"
	}

	var count int
	err := DB.QueryRow(query).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetItemById(id int) (models.InventoryAll, error) {
	query := `
	SELECT
		i.id,
		v.vendor_id,
		v.vendor_name,
		l.location_id,
		l.location_name,
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
	JOIN locations l ON i.location_id = l.location_id
	WHERE i.id = ?;
	`

	var item models.InventoryAll
	err := DB.QueryRow(query, id).Scan(
		&item.ItemID,
		&item.VendorID,
		&item.VendorName,
		&item.LocationID,
		&item.LocationName,
		&item.UPC,
		&item.InvoiceNumber,
		&item.Brand,
		&item.Name,
		&item.Description,
		&item.Price,
		&item.Weighed,
		&item.Count,
		&item.Deleted,
		&item.ArrivedAt,
		&item.SoldOutAt,
	)

	if err != nil {
		return item, err
	}

	return item, nil
}

func UpdateItem(id int, item models.Inventory) error {
	query := `
	UPDATE inventory
	SET vendor_id = ?, location_id = ?, upc = ?, invoice_number = ?, brand = ?, name = ?, description = ?, price = ?, weighed = ?, count = ?, deleted = ?, arrived_at = ?, soldout_at = ?
	WHERE id = ?;
	`
	_, err := DB.Exec(query, item.VendorID, item.LocationID, item.UPC, item.InvoiceNumber, item.Brand, item.Name, item.Description, item.Price, item.Weighed, item.Count, item.Deleted, item.ArrivedAt, item.SoldOutAt, id)
	return err
}

func DeleteItem(id int) error {
	query := `
	UPDATE inventory
	SET deleted = true, soldout_at = CURRENT_TIMESTAMP
	WHERE id = ?;
	`
	_, err := DB.Exec(query, id)
	return err
}

func RestoreItem(id int) error {
	query := `
	UPDATE inventory
	SET deleted = false, soldout_at = NULL
	WHERE id = ?;
	`
	_, err := DB.Exec(query, id)
	return err
}

func DeleteItemPermanent(id int) error {
	_, err := DB.Exec(`DELETE FROM inventory WHERE id = ?;`, id)
	return err
}
