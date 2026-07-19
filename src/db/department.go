package db

import (
	"errors"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddDepartmentEntry(department models.Department) error {
	query := `
	INSERT INTO departments (department_name)
	VALUES (?)
	`
	_, err := DB.Exec(query, department.DepartmentName)
	return err
}

func GetAllDepartments() ([]models.Department, error) {
	query := `
	SELECT
		d.department_id,
		d.department_name,
		(SELECT COALESCE(COUNT(i.id), 0)
		 FROM inventory AS i
	     WHERE d.department_id = i.department_id) AS total_items
	FROM departments AS d
	WHERE department_id != 1
	ORDER BY department_id DESC
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []models.Department
	for rows.Next() {
		var d models.Department
		rows.Scan(&d.DepartmentID, &d.DepartmentName, &d.DepartmentItemCount)
		departments = append(departments, d)
	}
	return departments, nil
}

func DeleteDepartmentEntry(id int) error {
	s, err := DB.Exec(`DELETE FROM departments WHERE department_id = ?`, id)
	if err != nil {
		return err
	}
	if r, _ := s.RowsAffected(); r == 0 {
		return errors.New("department id does not exist")
	}
	return err
}

func UpdateDepartmentEntry(department models.Department) error {
	query := `
	UPDATE departments
	SET department_name = ?
	WHERE department_id = ?
	`
	s, err := DB.Exec(query, department.DepartmentName, department.DepartmentID)
	if err != nil {
		return err
	}
	if r, _ := s.RowsAffected(); r == 0 {
		return errors.New("department entry does not exist")
	}
	return err
}
