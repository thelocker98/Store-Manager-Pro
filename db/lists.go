package db

import (
	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func GetLists() ([]models.List, error) {
	query := `
	SELECT
		l.list_id,
		l.list_name,
		COALESCE(COUNT(ld.list_count), 0) AS total_items,
		l.created_at
	FROM lists l
	LEFT JOIN list_data ld ON l.list_id = ld.list_id
	GROUP BY l.list_id, l.list_name, l.created_at
	ORDER BY l.list_id DESC;
	`

	rows, err := DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []models.List
	for rows.Next() {
		var l models.List
		err := rows.Scan(
			&l.ListID,
			&l.ListName,
			&l.TotalCount,
			&l.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		lists = append(lists, l)
	}
	return lists, nil
}

func GetListById(listID int) (models.List, error) {
	query := `
	SELECT
		list_id,
		list_name,
		created_at
	FROM lists
	WHERE list_id = ?;
	`

	var list models.List
	err := DB.QueryRow(query, listID).Scan(
		&list.ListID,
		&list.ListName,
		&list.CreatedAt,
	)

	if err != nil {
		return list, err
	}

	return list, nil
}

func GetNumberOfItemsList(listID int) (int, error) {
	query := `
		SELECT SUM(list_count) AS total_items
		FROM list_data
		WHERE list_id = ?;
	`

	var totalCount int
	err := DB.QueryRow(query, listID).Scan(
		&totalCount,
	)

	if err != nil {
		return -1, err
	}

	return totalCount, nil
}

func AddList(list models.List) error {
	query := `
	INSERT INTO lists (list_name)
	VALUES (?);
	`
	_, err := DB.Exec(query, list.ListName)
	return err
}

func UpdateList(listID int, list models.List) error {
	query := `
	UPDATE lists
	SET list_name = ?
	WHERE list_id = ?;
	`
	_, err := DB.Exec(query, list.ListName, listID)
	return err
}

func DeleteList(listID int) error {
	_, err := DB.Exec(`DELETE FROM lists WHERE list_id = ?;`, listID)
	return err
}

func GetListEntrys(listID int) ([]models.ListEntryAll, error) {
	query := `
	SELECT
		ld.list_entry_id,
		ld.list_id,
		ld.item_id,
		v.vendor_id,
		l.location_id,
		ld.list_count,
		v.vendor_name,
		l.location_name,
		i.upc,
		i.invoice_number,
		i.brand,
		i.name,
		i.description,
		i.price,
		i.weighed,
		i.count,
		i.deleted
	FROM list_data ld
	JOIN lists l on ld.list_id = l.list_id
	JOIN inventory i on ld.item_id = i.id
	JOIN vendors v ON i.vendor_id = v.vendor_id
	JOIN locations l ON i.location_id = l.location_id
	WHERE ld.list_id = ?
	ORDER BY l.location_name ASC, ld.list_entry_id DESC;
	`

	rows, err := DB.Query(query, listID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []models.ListEntryAll
	for rows.Next() {
		var l models.ListEntryAll
		err := rows.Scan(
			&l.EntryID,
			&l.ListID,
			&l.ItemID,
			&l.VendorID,
			&l.LocationID,
			&l.ListCount,
			&l.VendorName,
			&l.LocationName,
			&l.UPC,
			&l.InvoiceNumber,
			&l.Brand,
			&l.Name,
			&l.Description,
			&l.Price,
			&l.Weighed,
			&l.Count,
			&l.Deleted,
		)
		if err != nil {
			return nil, err
		}
		lists = append(lists, l)
	}
	return lists, nil
}

func GetListEntryById(listID int) (models.ListEntryAll, error) {
	query := `
	SELECT
		ld.list_entry_id,
		ld.list_id,
		ld.item_id,
		v.vendor_id,
		l.location_id,
		ld.list_count,
		v.vendor_name,
		l.location_name,
		i.upc,
		i.invoice_number,
		i.brand,
		i.name,
		i.description,
		i.price,
		i.weighed,
		i.count,
		i.deleted
	FROM list_data ld
	JOIN lists l on ld.list_id = l.list_id
	JOIN inventory i on ld.item_id = i.id
	JOIN vendors v ON i.vendor_id = v.vendor_id
	JOIN locations l ON i.location_id = l.location_id
	WHERE ld.list_entry_id = ?;
	`

	var list models.ListEntryAll

	err := DB.QueryRow(query, listID).Scan(
		&list.EntryID,
		&list.ListID,
		&list.ItemID,
		&list.VendorID,
		&list.LocationID,
		&list.ListCount,
		&list.VendorName,
		&list.LocationName,
		&list.UPC,
		&list.InvoiceNumber,
		&list.Brand,
		&list.Name,
		&list.Description,
		&list.Price,
		&list.Weighed,
		&list.Count,
		&list.Deleted,
	)

	return list, err
}

func AddEntryToList(listEntry models.ListEntry) error {
	query := `
	INSERT INTO list_data (list_id,  item_id, list_count)
	VALUES (?, ?, 1);
	`
	_, err := DB.Exec(query, listEntry.ListID, listEntry.ItemID)
	return err
}

func UpdateEntryInList(entryID int, listEntry models.ListEntry) error {
	query := `
	UPDATE list_data
	SET list_count = ?
	WHERE list_entry_id = ?;
	`
	_, err := DB.Exec(query, listEntry.ListCount, entryID)
	return err
}

func DeleteEntryFromList(listEntryID int) error {
	_, err := DB.Exec(`DELETE FROM list_data WHERE list_entry_id = ?;`, listEntryID)
	return err
}
