package db

import (
	"fmt"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddVendor(vendor models.Vendor) error {
	query := `
	INSERT INTO vendors (vendor_name)
	VALUES (?)
	`
	_, err := DB.Exec(query, vendor.VendorName)
	fmt.Println(err)
	return err
}

func GetAllVendorss() ([]models.Vendor, error) {
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

func DeleteVendor(id int) error {
	_, err := DB.Exec(`DELETE FROM vendors WHERE vendor_id = ?`, id)
	return err
}

func UpdateVendor(vendor models.Vendor) error {
	query := `
	UPDATE vendors
	SET vendor_name = ?
	WHERE vendor_id = ?
	`
	_, err := DB.Exec(query, vendor.VendorName, vendor.VendorID)
	return err
}
