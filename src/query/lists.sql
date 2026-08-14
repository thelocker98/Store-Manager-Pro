-- name: GetLists :many
SELECT
	l.list_id,
	l.list_name,
	l.department_id,
	COALESCE(COUNT(ld.list_count), 0)::bigint AS total_items,
	l.created_at
FROM lists l
LEFT JOIN list_data ld ON l.list_id = ld.list_id
WHERE (sqlc.arg(department_id)::integer = 0 OR l.department_id = sqlc.arg(department_id))
GROUP BY l.list_id, l.list_name, l.department_id, l.created_at
ORDER BY l.list_id DESC;

-- name: GetListById :one
SELECT
	l.list_id,
	l.list_name,
	l.department_id,
	d.department_name,
	COALESCE(COUNT(ld.list_count), 0)::bigint AS total_items,
	l.created_at
FROM lists l
JOIN departments d ON l.department_id = d.department_id
LEFT JOIN list_data ld ON l.list_id = ld.list_id
WHERE l.list_id = $1
GROUP BY l.list_id, l.list_name, l.department_id, d.department_name, l.created_at;

-- name: GetNumberOfItemsList :one
SELECT COALESCE(SUM(list_count), 0)::bigint AS total_items
FROM list_data
WHERE list_id = $1;

-- name: AddList :execrows
INSERT INTO lists (list_name, department_id)
	VALUES ($1, $2);

-- name: UpdateList :execrows
UPDATE lists
	SET list_name = $1
	WHERE list_id = $2;

-- name: DeleteList :execrows
DELETE FROM lists WHERE list_id = $1;

-- name: GetListEntrys :many
SELECT
	ld.list_entry_id,
	ld.list_id,
	ld.item_id,
	v.vendor_id,
	loc.location_id,
	ld.list_count,
	v.vendor_name,
	loc.location_name,
	i.upc,
	i.invoice_number,
	i.brand,
	i.name,
	i.description,
	i.price,
	i.weighed,
	COALESCE(ld.list_count, 0) AS count,
	i.deleted
FROM list_data ld
JOIN lists lst ON ld.list_id = lst.list_id
JOIN inventory i ON ld.item_id = i.id
JOIN vendors v ON i.vendor_id = v.vendor_id
JOIN locations loc ON i.location_id = loc.location_id
WHERE ld.list_id = $1
ORDER BY loc.location_name ASC, ld.list_entry_id DESC;

-- name: GetListEntryById :one
SELECT
	ld.list_entry_id,
	ld.list_id,
	ld.item_id,
	v.vendor_id,
	loc.location_id,
	ld.list_count,
	v.vendor_name,
	loc.location_name,
	i.upc,
	i.invoice_number,
	i.brand,
	i.name,
	i.description,
	i.price,
	i.weighed,
	COALESCE(i.count, 0) AS count,
	i.deleted
FROM list_data ld
JOIN lists lst ON ld.list_id = lst.list_id
JOIN inventory i ON ld.item_id = i.id
JOIN vendors v ON i.vendor_id = v.vendor_id
JOIN locations loc ON i.location_id = loc.location_id
WHERE ld.list_entry_id = $1;

-- name: AddEntryToList :one
INSERT INTO list_data (list_id, item_id, list_count)
	VALUES ($1, $2, 1)
	RETURNING list_entry_id;

-- name: UpdateEntryInList :execrows
UPDATE list_data
	SET list_count = $1
	WHERE list_entry_id = $2;

-- name: DeleteEntryFromList :execrows
DELETE FROM list_data WHERE list_entry_id = $1;
