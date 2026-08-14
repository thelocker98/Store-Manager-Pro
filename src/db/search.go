package db

import (
	"fmt"
	"strings"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func SearchItems(search string, list_id int, page int, pageSize int, showdeleted bool, vendor int, location int, department int, limitToList bool, orderBy string) ([]models.InventoryAll, error) {
	// Calculate pages
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 50
	}
	offset := (page - 1) * pageSize

	// Search String
	search = "%" + strings.ToLower(search) + "%"

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
			COALESCE(i.count, 0) AS count,
			i.deleted,
			i.arrived_at,
			i.soldout_at,
			COALESCE(ld.list_entry_id, -1) AS list_entry_id,
			COALESCE(ld.list_count, -1) as list_count,
			COUNT(*) OVER() AS entry_count
		FROM inventory i
		JOIN vendors v ON i.vendor_id = v.vendor_id
		JOIN locations l on l.location_id = i.location_id
		LEFT JOIN list_data ld ON i.id = ld.item_id AND ld.list_id = $1
		WHERE
			(LOWER(i.name)        LIKE $2
		 OR LOWER(i.brand)       LIKE $3
		 OR LOWER(i.description) LIKE $4
		 OR LOWER(v.vendor_name) LIKE $5
		 OR LOWER(i.upc) LIKE $6
	     OR LOWER(i.invoice_number) LIKE $7)
	`
	// Don't show deleted items
	if !showdeleted {
		query += ` AND i.deleted = FALSE
			`
	}
	if vendor != 0 {
		query += ` AND i.vendor_id = ` + fmt.Sprint(vendor)
	}
	if location != 0 {
		query += ` AND i.location_id = ` + fmt.Sprint(location)
	}
	if department != 0 {
		query += ` AND i.department_id = ` + fmt.Sprint(department)
	}
	if limitToList == true {
		query += ` AND ld.list_id = ` + fmt.Sprint(list_id)
	}

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

	query += ` LIMIT $8 OFFSET $9;`

	rows, err := DB.Query(query, list_id, search, search, search, search, search, search, pageSize, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.InventoryAll
	for rows.Next() {
		var i models.InventoryAll
		if err := rows.Scan(
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
			&i.EntryID,
			&i.ListCount,
			&i.Entry_Count,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func SearchInvoice(search string, invoice_id int, page int, pageSize int) ([]models.InvoiceEntry, error) {
	// Calculate pages
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 50
	}
	offset := (page - 1) * pageSize

	// Search String
	search = "%" + strings.ToLower(search) + "%"
	fmt.Println(invoice_id, search)

	query := `
		SELECT
			id.invoice_data_id,
			id.invoice_id,
			id.rawtext,
			id.upc,
			id.details,
			id.qty,
			id.item_cost,
			id.total_cost,
			id.discount_cost,
			id.true_cost,
			COUNT(*) OVER() AS total_items

		FROM invoice_data id
		WHERE
			invoice_id = $1 AND (
				LOWER(id.rawtext)        LIKE $2
				OR LOWER(id.upc)       LIKE $3
				OR LOWER(id.details) LIKE $4
			)
		ORDER BY id.qty DESC
		LIMIT $5 OFFSET $6;
	`

	rows, err := DB.Query(query, invoice_id, search, search, search, pageSize, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entrys []models.InvoiceEntry
	for rows.Next() {
		var i models.InvoiceEntry
		if err := rows.Scan(
			&i.InvoiceDataID,
			&i.InvoiceID,
			&i.RawText,
			&i.UPC,
			&i.Detail,
			&i.QTY,
			&i.ItemCost,
			&i.TotalCost,
			&i.DiscountCost,
			&i.TrueCost,
			&i.EntryCount,
		); err != nil {
			return nil, err
		}
		entrys = append(entrys, i)
	}
	return entrys, nil
}
