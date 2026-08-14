-- name: GetDepartments :many
SELECT  d.department_id,
        d.department_name,
        CAST((SELECT COALESCE(COUNT(i.id), 0)
	FROM inventory AS i
	WHERE d.department_id = i.department_id) AS INTEGER) AS total_items
	FROM departments AS d
	WHERE department_id != 1
	ORDER BY department_id DESC;

-- name: GetDepartment :one
SELECT  d.department_id,
        d.department_name,
        CAST((SELECT COALESCE(COUNT(i.id), 0) FROM inventory AS i WHERE d.department_id = i.department_id) AS INTEGER) AS total_items
	FROM departments AS d
	WHERE d.department_id = $1 AND d.department_id != 1;

-- name: AddDepartment :execrows
INSERT INTO departments (department_name)
	VALUES ($1);

-- name: UpdateDepartment :execrows
UPDATE departments
    SET department_name = $1
	WHERE department_id = $2;

-- name: DeleteDepartment :execrows
DELETE FROM departments WHERE department_id = $1;
