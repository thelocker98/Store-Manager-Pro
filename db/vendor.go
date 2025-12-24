package db

import (
	"errors"
	"fmt"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddVendorEntry(vendor models.Vendor) error {
	query := `
	INSERT INTO vendors (vendor_name)
	VALUES (?)
	`
	_, err := DB.Exec(query, vendor.VendorName)
	fmt.Println(err)
	return err
}

func GetAllVendors() ([]models.Vendor, error) {
	rows, err := DB.Query(`SELECT vendor_id, vendor_name FROM vendors ORDER BY vendor_name ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendors []models.Vendor
	for rows.Next() {
		var v models.Vendor
		rows.Scan(&v.VendorID, &v.VendorName)
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
