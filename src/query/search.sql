-- name: SearchItems :many
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
	COALESCE(ld.list_count, -1) AS list_count,
	COUNT(*) OVER() AS entry_count
FROM inventory i
JOIN vendors v ON i.vendor_id = v.vendor_id
JOIN locations l ON l.location_id = i.location_id
LEFT JOIN list_data ld ON i.id = ld.item_id AND ld.list_id = sqlc.arg(list_id)
WHERE
	(LOWER(i.name) LIKE sqlc.arg(search)
	 OR LOWER(i.brand) LIKE sqlc.arg(search)
	 OR LOWER(i.description) LIKE sqlc.arg(search)
	 OR LOWER(v.vendor_name) LIKE sqlc.arg(search)
	 OR LOWER(i.upc) LIKE sqlc.arg(search)
	 OR LOWER(i.invoice_number) LIKE sqlc.arg(search))
	AND (sqlc.arg(show_deleted)::boolean = true OR i.deleted = false)
	AND (sqlc.arg(vendor)::integer = 0 OR i.vendor_id = sqlc.arg(vendor))
	AND (sqlc.arg(location)::integer = 0 OR i.location_id = sqlc.arg(location))
	AND (sqlc.arg(department)::integer = 0 OR i.department_id = sqlc.arg(department))
	AND (sqlc.arg(limit_to_list)::boolean = false OR ld.list_id = sqlc.arg(list_id))
ORDER BY CASE sqlc.arg(order_by)::text
	WHEN 'date' THEN CAST(i.arrived_at AS TEXT)
	WHEN 'name' THEN i.name
	ELSE i.upc
END DESC
LIMIT sqlc.arg(page_size) OFFSET sqlc.arg(page_offset);

-- name: SearchInvoice :many
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
		LOWER(id.rawtext) LIKE $2
		OR LOWER(id.upc) LIKE $2
		OR LOWER(id.details) LIKE $2
	)
ORDER BY id.qty DESC
LIMIT $3 OFFSET $4;
