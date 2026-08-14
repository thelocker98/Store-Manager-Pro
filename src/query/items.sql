-- name: AddItem :execrows
INSERT INTO inventory (vendor_id, location_id, department_id, upc, invoice_number, name, brand, description, price, weighed, count)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);

-- name: GetAllItems :many
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
LEFT JOIN list_data ld ON i.id = ld.item_id AND ld.list_id = sqlc.arg(list_id)
WHERE (sqlc.arg(show_deleted)::boolean = true OR i.deleted = false)
	AND (sqlc.arg(vendor)::integer = 0 OR i.vendor_id = sqlc.arg(vendor))
	AND (sqlc.arg(location)::integer = 0 OR i.location_id = sqlc.arg(location))
	AND (sqlc.arg(department)::integer = 0 OR i.department_id = sqlc.arg(department))
	AND (sqlc.arg(list_id) = 0 OR ld.list_id = sqlc.arg(list_id))
ORDER BY CASE sqlc.arg(order_by)::text
	WHEN 'date' THEN CAST(i.arrived_at AS TEXT)
	WHEN 'name' THEN i.name
	ELSE i.upc
END DESC
LIMIT sqlc.arg(page_size) OFFSET sqlc.arg(page_offset);

-- name: CountAllItems :one
SELECT COUNT(id)
FROM inventory
WHERE (sqlc.arg(count_deleted)::boolean = true OR deleted = false);

-- name: GetItemById :one
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

-- name: UpdateItem :execrows
UPDATE inventory
	SET department_id = $1, vendor_id = $2, location_id = $3, upc = $4, invoice_number = $5, brand = $6, name = $7, description = $8, price = $9, weighed = $10, count = $11, deleted = $12, arrived_at = $13, soldout_at = $14
	WHERE id = $15;

-- name: UpdateItemNoDepartment :execrows
UPDATE inventory
	SET vendor_id = $1, location_id = $2, upc = $3, invoice_number = $4, brand = $5, name = $6, description = $7, price = $8, weighed = $9, count = $10, deleted = $11, arrived_at = $12, soldout_at = $13
	WHERE id = $14;

-- name: UpdateItemDepartment :execrows
UPDATE inventory
	SET location_id = $1, department_id = $2
	WHERE id = $3;

-- name: DeleteItem :execrows
UPDATE inventory
	SET deleted = true, soldout_at = CURRENT_TIMESTAMP
	WHERE id = $1;

-- name: RestoreItem :execrows
UPDATE inventory
	SET deleted = false, soldout_at = NULL
	WHERE id = $1;

-- name: DeleteItemPermanent :execrows
DELETE FROM inventory WHERE id = $1;
