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
	rows, err := DB.Query(`SELECT department_id, department_name FROM departments WHERE department_id != 1 ORDER BY department_id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []models.Department
	for rows.Next() {
		var v models.Department
		rows.Scan(&v.DepartmentID, &v.DepartmentName)
		departments = append(departments, v)
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
