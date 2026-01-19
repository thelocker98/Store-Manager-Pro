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
			i.count,
			i.deleted,
			i.arrived_at,
			i.soldout_at,
			COALESCE(ld.list_entry_id, -1) AS list_entry_id,
			COALESCE(ld.list_count, -1) as list_count,
			COUNT(*) OVER() AS entry_count
		FROM inventory i
		JOIN vendors v ON i.vendor_id = v.vendor_id
		JOIN locations l on l.location_id = i.location_id
		LEFT JOIN list_data ld ON i.id = ld.item_id AND ld.list_id = ?
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
		query += ` AND i.deleted = 0
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
