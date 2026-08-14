-- name: GetLocations :many
SELECT  l.location_id,
        l.location_name,
        CAST((SELECT COALESCE(COUNT(i.id), 0)
	FROM inventory AS i
	WHERE i.location_id = l.location_id) AS INTEGER) AS total_items
	FROM locations AS l
	WHERE l.location_id != 1
	ORDER BY location_id DESC;

-- name: AddLocation :execrows
INSERT INTO locations (location_name)
	VALUES ($1);

-- name: UpdateLocation :execrows
UPDATE locations
    SET location_name = $1
	WHERE location_id = $2;

-- name: DeleteLocation :execrows
DELETE FROM locations WHERE location_id = $1;
