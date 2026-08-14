-- name: GetInvoices :many
SELECT
	v.invoice_id,
	v.invoice_name,
	v.invoice_type,
	v.department_id,
	d.department_name,
	COALESCE(COUNT(vd.qty), 0)::bigint AS total_count,
	v.created_at
FROM invoices v
LEFT JOIN invoice_data vd ON v.invoice_id = vd.invoice_id
LEFT JOIN departments d ON v.department_id = d.department_id
WHERE (sqlc.arg(department_id)::integer = 0 OR v.department_id = sqlc.arg(department_id))
GROUP BY v.invoice_id, v.invoice_name, v.invoice_type, v.department_id, d.department_name, v.created_at
ORDER BY v.invoice_id DESC;

-- name: GetInvoiceById :one
SELECT
	v.invoice_id,
	v.invoice_name,
	v.invoice_type,
	v.department_id,
	d.department_name,
	COALESCE(COUNT(vd.qty), 0)::bigint AS total_count,
	v.created_at
FROM invoices v
LEFT JOIN invoice_data vd ON v.invoice_id = vd.invoice_id
JOIN departments d ON v.department_id = d.department_id
WHERE v.invoice_id = $1
GROUP BY v.invoice_id, v.invoice_name, v.invoice_type, v.department_id, d.department_name, v.created_at;

-- name: AddInvoice :execrows
INSERT INTO invoices (invoice_name, invoice_type, department_id)
	VALUES ($1, $2, $3);

-- name: UpdateInvoice :execrows
UPDATE invoices
	SET invoice_name = $1
	WHERE invoice_id = $2;

-- name: DeleteInvoice :execrows
DELETE FROM invoices WHERE invoice_id = $1;

-- name: DeleteInvoiceData :execrows
DELETE FROM invoice_data WHERE invoice_id = $1;

-- name: GetInvoiceEntrys :many
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
WHERE id.invoice_id = $1
ORDER BY id.qty DESC
LIMIT $2 OFFSET $3;

-- name: CountAllInvoiceEntrys :one
SELECT COUNT(invoice_data_id)
FROM invoice_data
WHERE invoice_id = $1;

-- name: GetInvoiceEntryById :one
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
WHERE id.invoice_data_id = $1;

-- name: AddEntryToInvoice :execrows
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
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: UpdateEntryInInvoice :execrows
UPDATE invoice_data
	SET
		rawtext = $1,
		upc = $2,
		details = $3,
		qty = $4,
		item_cost = $5,
		total_cost = $6,
		discount_cost = $7,
		true_cost = $8
	WHERE invoice_data_id = $9;

-- name: DeleteEntryFromInvoice :execrows
DELETE FROM invoice_data WHERE invoice_data_id = $1;
