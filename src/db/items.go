package db

import (
	"fmt"
	"strings"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddItem(item models.Inventory) error {
	query := `
	INSERT INTO inventory (vendor_id, location_id, department_id, upc, invoice_number, name, brand, description, price, weighed, count)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
	`
	_, err := DB.Exec(query, item.VendorID, item.LocationID, item.DepartmentID, item.UPC, item.InvoiceNumber, item.Name, item.Brand, item.Description, item.Price, item.Weighed, item.Count)
	return err
}

func GetAllItems(page int, pageSize int, showDeleted bool, vendor int, location int, department int, listId int, orderBy string) ([]models.InventoryAll, error) {
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
		i.department_id,
		d.department_name,
		i.upc,
		i.invoice_number,
		i.brand,
		i.name,
		i.description,
		i.price,
		i.weighed,
		COALESCE(i.count, 0) AS count,
		i.deleted,
		i.arrived_at,
		i.soldout_at,
		COUNT(*) OVER() AS entry_count
	FROM inventory i
	JOIN vendors v ON i.vendor_id = v.vendor_id
	JOIN locations l ON i.location_id = l.location_id
	JOIN departments d ON i.department_id = d.department_id
	`
	if listId != 0 {
		query += `LEFT JOIN list_data ld ON i.id = ld.item_id AND ld.list_id = ` + fmt.Sprint(listId)
	}

	query += ` WHERE `
	if !showDeleted {
		query += `i.deleted = FALSE AND `
	}
	if vendor != 0 {
		query += `i.vendor_id = ` + fmt.Sprint(vendor) + ` AND `
	}
	if location != 0 {
		query += `i.location_id = ` + fmt.Sprint(location) + ` AND `
	}
	if department != 0 {
		query += `i.department_id = ` + fmt.Sprint(department) + ` AND `
	}
	if listId != 0 {
		query += `ld.list_id = ` + fmt.Sprint(listId) + ` AND `
	}

	query += `1=1 `

	// sort
	switch strings.ToLower(orderBy) {
	case "date":
		query += ` ORDER BY i.arrived_at DESC`
	case "name":
		query += ` ORDER BY i.name DESC`
	case "upc":
		query += ` ORDER BY i.upc DESC`
	default:
	}

	query += ` LIMIT $1 OFFSET $2;`

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
			&i.DepartmentID,
			&i.DepartmentName,
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
		query += " WHERE deleted = FALSE;"
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
		i.department_id,
		d.department_name,
		i.upc,
		i.invoice_number,
		i.brand,
		i.name,
		i.description,
		i.price,
		i.weighed,
		COALESCE(i.count, 0) AS count,
		i.deleted,
		i.arrived_at,
		i.soldout_at
	FROM inventory i
	JOIN vendors v ON i.vendor_id = v.vendor_id
	JOIN locations l ON i.location_id = l.location_id
	JOIN departments d ON i.department_id = d.department_id
	WHERE i.id = $1;
	`

	var item models.InventoryAll
	err := DB.QueryRow(query, id).Scan(
		&item.ItemID,
		&item.VendorID,
		&item.VendorName,
		&item.LocationID,
		&item.LocationName,
		&item.DepartmentID,
		&item.DepartmentName,
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
	if item.DepartmentID == 0 {
		query := `
		UPDATE inventory
		SET vendor_id = $1, location_id = $2, upc = $3, invoice_number = $4, brand = $5, name = $6, description = $7, price = $8, weighed = $9, count = $10, deleted = $11, arrived_at = $12, soldout_at = $13
		WHERE id = $14;
		`
		_, err := DB.Exec(query, item.VendorID, item.LocationID, item.UPC, item.InvoiceNumber, item.Brand, item.Name, item.Description, item.Price, item.Weighed, item.Count, item.Deleted, item.ArrivedAt, item.SoldOutAt, id)
		return err
	} else {
		query := `
		UPDATE inventory
		SET department_id = $1, vendor_id = $2, location_id = $3, upc = $4, invoice_number = $5, brand = $6, name = $7, description = $8, price = $9, weighed = $10, count = $11, deleted = $12, arrived_at = $13, soldout_at = $14
		WHERE id = $15;
		`
		_, err := DB.Exec(query, item.DepartmentID, item.VendorID, item.LocationID, item.UPC, item.InvoiceNumber, item.Brand, item.Name, item.Description, item.Price, item.Weighed, item.Count, item.Deleted, item.ArrivedAt, item.SoldOutAt, id)
		return err
	}
}

func UpdateItemDepartment(id int, item models.Inventory) error {
	fmt.Println(id, item)
	query := `
	UPDATE inventory
	SET location_id = $1, department_id = $2
	WHERE id = $3;
	`
	_, err := DB.Exec(query, item.LocationID, item.DepartmentID, id)
	return err
}

func DeleteItem(id int) error {
	query := `
	UPDATE inventory
	SET deleted = true, soldout_at = CURRENT_TIMESTAMP
	WHERE id = $1;
	`
	_, err := DB.Exec(query, id)
	return err
}

func RestoreItem(id int) error {
	query := `
	UPDATE inventory
	SET deleted = false, soldout_at = NULL
	WHERE id = $1;
	`
	_, err := DB.Exec(query, id)
	return err
}

func DeleteItemPermanent(id int) error {
	_, err := DB.Exec(`DELETE FROM inventory WHERE id = $1;`, id)
	return err
}
