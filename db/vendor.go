package db

import (
	"errors"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddVendorEntry(vendor models.Vendor) error {
	query := `
	INSERT INTO vendors (vendor_name)
	VALUES (?)
	`
	_, err := DB.Exec(query, vendor.VendorName)
	return err
}

func GetAllVendors() ([]models.Vendor, error) {
	query := `
	SELECT
		v.vendor_id,
		v.vendor_name,
		(SELECT COALESCE(COUNT(i.id), 0)
		 FROM inventory AS i
	     WHERE i.vendor_id = v.vendor_id) AS total_items
	FROM vendors AS v
	WHERE vendor_id != 1
	ORDER BY vendor_id DESC
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendors []models.Vendor
	for rows.Next() {
		var v models.Vendor
		rows.Scan(&v.VendorID, &v.VendorName, &v.VendorItemCount)
		vendors = append(vendors, v)
	}
	return vendors, nil
}

func DeleteVendorEntry(id int) error {
	s, err := DB.Exec(`DELETE FROM vendors WHERE vendor_id = ?`, id)
	if err != nil {
		return err
	}
	if r, _ := s.RowsAffected(); r == 0 {
		return errors.New("vendor id does not exist")
	}
	return err
}

func UpdateVendorEntry(vendor models.Vendor) error {
	query := `
	UPDATE vendors
	SET vendor_name = ?
	WHERE vendor_id = ?
	`
	s, err := DB.Exec(query, vendor.VendorName, vendor.VendorID)
	if err != nil {
		return err
	}
	if r, _ := s.RowsAffected(); r == 0 {
		return errors.New("vendor entry does not exist")
	}
	return err
}
