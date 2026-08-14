-- name: GetVendors :many
SELECT  v.vendor_id,
        v.vendor_name,
        CAST((SELECT COALESCE(COUNT(i.id), 0)
	FROM inventory AS i
	WHERE i.vendor_id = v.vendor_id) AS INTEGER) AS total_items
	FROM vendors AS v
	WHERE vendor_id != 1
	ORDER BY vendor_id DESC;

-- name: AddVendor :execrows
INSERT INTO vendors (vendor_name)
	VALUES ($1);

-- name: UpdateVendor :execrows
UPDATE vendors
    SET vendor_name = $1
	WHERE vendor_id = $2;

-- name: DeleteVendor :execrows
DELETE FROM vendors WHERE vendor_id = $1;
