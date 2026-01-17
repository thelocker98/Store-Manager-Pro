package db

import (
	"strings"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func SearchItems(search string, order string, showdeleted bool, page int, pageSize int, list_id int) ([]models.InventoryAll, error) {
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
			i.count,
			i.deleted,
			i.arrived_at,
			i.soldout_at,
			COALESCE(l.list_entry_id, -1) AS list_entry_id,
			COALESCE(l.list_count, -1) as list_count,
			COUNT(*) OVER() AS entry_count
		FROM inventory i
		JOIN vendors v ON i.vendor_id = v.vendor_id
		JOIN locations l on l.location_id = i.location_id
		LEFT JOIN list_data l ON i.id = l.item_id AND l.list_id = ?
		WHERE
			(LOWER(i.name)        LIKE ?
		 OR LOWER(i.brand)       LIKE ?
		 OR LOWER(i.description) LIKE ?
		 OR LOWER(v.vendor_name) LIKE ?
		 OR LOWER(i.upc) LIKE ?
	     OR LOWER(i.invoice_number) LIKE ?)
	`
	// Don't show deleted items
	if !showdeleted {
		query += `AND i.deleted = 0
			`
	}

	// sort
	switch strings.ToLower(order) {
	case "date":
		query += `ORDER BY i.arrived_at DESC`
	case "name":
		query += `ORDER BY i.name DESC`
	default:
		query += `ORDER BY i.upc DESC`
	}

	query += ` LIMIT ? OFFSET ?;`

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
