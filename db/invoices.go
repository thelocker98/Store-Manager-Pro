package db

import (
	"fmt"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func GetInvoices(departmentid int) ([]models.Invoice, error) {
	query := `
	SELECT
		v.invoice_id,
		v.invoice_name,
		v.invoice_type,
		v.department_id,
		d.department_name,
		COALESCE(COUNT(vd.qty), 0) AS total_items,
		v.created_at
	FROM invoices v
	LEFT JOIN  invoice_data vd ON v.invoice_id = vd.invoice_id
	LEFT JOIN departments d ON v.department_id = d.department_id
	`
	if departmentid != 0 {
		query += `WHERE v.department_id = ` + fmt.Sprint(departmentid)
	}

	query += `
	GROUP BY v.invoice_id, v.invoice_name, v.created_at
	ORDER BY v.invoice_id DESC;
	`

	rows, err := DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []models.Invoice
	for rows.Next() {
		var v models.Invoice
		err := rows.Scan(
			&v.InvoiceID,
			&v.InvoiceName,
			&v.InvoiceType,
			&v.DepartmentID,
			&v.DepartmentName,
			&v.TotalCount,
			&v.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, v)
	}
	return invoices, nil
}

func GetInvoiceById(InvoiceID int) (models.Invoice, error) {
	query := `
	SELECT
		v.invoice_id,
		v.invoice_name,
		v.invoice_type,
		v.department_id,
		d.department_name,
		v.created_at
	FROM invoices v
	JOIN departments d ON v.department_id = d.department_id
	WHERE v.invoice_id= ?;
	`

	var invoice models.Invoice
	err := DB.QueryRow(query, InvoiceID).Scan(
		&invoice.InvoiceID,
		&invoice.InvoiceName,
		&invoice.InvoiceType,
		&invoice.DepartmentID,
		&invoice.DepartmentName,
		&invoice.CreatedAt,
	)

	if err != nil {
		return invoice, err
	}

	return invoice, nil
}

func AddInvoice(invoice models.Invoice) error {
	query := `
	INSERT INTO lists (invoice_name, invoice_type, department_id)
	VALUES (?, ?, ?);
	`
	_, err := DB.Exec(query, invoice.InvoiceName, invoice.InvoiceType, invoice.DepartmentID)
	return err
}

func UpdateInvoice(invoiceID int, invoice models.Invoice) error {
	query := `
	UPDATE invoice
	SET invoice_name = ?
	WHERE invoice_id = ?;
	`
	_, err := DB.Exec(query, invoice.InvoiceName, invoiceID)
	return err
}

func DeleteInvoice(invoiceID int) error {
	_, err := DB.Exec(`DELETE FROM invoice_data WHERE invoice_id = ?;`, invoiceID)
	if err != nil {
		return err
	}
	_, err = DB.Exec(`DELETE FROM invoices WHERE invoice_id = ?;`, invoiceID)
	return err
}

func GetInvoiceEntrys(invoiceID int) ([]models.InvoiceEntry, error) {
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
		id.true_cost
	FROM invoice_data id
	WHERE id.invoice_id = ?
	ORDER BY id.qty DESC;
	`

	rows, err := DB.Query(query, invoiceID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoiceEntrys []models.InvoiceEntry
	for rows.Next() {
		var i models.InvoiceEntry
		err := rows.Scan(
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
		)
		if err != nil {
			return nil, err
		}
		invoiceEntrys = append(invoiceEntrys, i)
	}
	return invoiceEntrys, nil
}

func GetInvoiceEntryById(invoiceID int) (models.InvoiceEntry, error) {
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
		id.true_cost
	FROM invoice_data id
	WHERE id.invoice_data_id = ?;
	`

	var invoiceEntry models.InvoiceEntry

	err := DB.QueryRow(query, invoiceID).Scan(
		&invoiceEntry.InvoiceDataID,
		&invoiceEntry.InvoiceID,
		&invoiceEntry.RawText,
		&invoiceEntry.UPC,
		&invoiceEntry.Detail,
		&invoiceEntry.QTY,
		&invoiceEntry.ItemCost,
		&invoiceEntry.TotalCost,
		&invoiceEntry.DiscountCost,
		&invoiceEntry.TrueCost,
	)

	return invoiceEntry, err
}

func AddEntryToInvoice(invoiceEntry models.InvoiceEntry) error {
	query := `
	INSERT INTO invoice_data (
    	invoice_id,
     	rawtext,
      	upc,
       	details,
        qty,
        item_cost,
        total_cost,
        discount_cost,
        true_cost
)
VALUES
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	_, err := DB.Exec(query,
		invoiceEntry.InvoiceDataID,
		invoiceEntry.RawText,
		invoiceEntry.UPC,
		invoiceEntry.Detail,
		invoiceEntry.QTY,
		invoiceEntry.ItemCost,
		invoiceEntry.TotalCost,
		invoiceEntry.DiscountCost,
		invoiceEntry.TrueCost,
	)

	return err
}

func UpdateEntryInInvoice(entryID int, invoiceEntry models.InvoiceEntry) error {
	query := `
	UPDATE invoice_data
	SET
     	rawtext = ?,
      	upc = ?,
       	details = ?,
        qty = ?,
        item_cost = ?,
        total_cost = ?,
        discount_cost = ?,
        true_cost = ?
    WHERE invoice_data_id = ?;
	`
	_, err := DB.Exec(query,
		invoiceEntry.InvoiceDataID,
		invoiceEntry.RawText,
		invoiceEntry.UPC,
		invoiceEntry.Detail,
		invoiceEntry.QTY,
		invoiceEntry.ItemCost,
		invoiceEntry.TotalCost,
		invoiceEntry.DiscountCost,
		invoiceEntry.TrueCost,
		entryID,
	)

	return err
}

func DeleteEntryFromInvoice(invoiceEntryID int) error {
	_, err := DB.Exec(`DELETE FROM invoice_data WHERE invoice_data_id = ?;`, invoiceEntryID)
	return err
}
