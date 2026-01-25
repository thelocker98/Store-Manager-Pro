package db

import (
	"errors"
	"fmt"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddLocationEntry(location models.Location) error {
	query := `
	INSERT INTO locations (location_name, department_id)
	VALUES (?, ?)
	`
	_, err := DB.Exec(query, location.LocationName, location.DepartmentID)
	return err
}

func GetAllLocations(departmentid int) ([]models.Location, error) {
	query := `
	SELECT
		l.location_id,
		l.location_name,
		l.department_id,
			(SELECT COALESCE(COUNT(i.id), 0)
			 FROM inventory AS i
		     WHERE i.location_id = l.location_id) AS total_items
		FROM locations AS l
		WHERE l.location_id != 1 `
	if departmentid != 0 {
		query += `AND l.department_id = ` + fmt.Sprint(departmentid)
	}
	query += `
	ORDER BY location_id DESC
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []models.Location
	for rows.Next() {
		var l models.Location
		rows.Scan(&l.LocationID, &l.LocationName, &l.DepartmentID, &l.LocationItemCount)
		locations = append(locations, l)
	}
	return locations, nil
}

func DeleteLocationEntry(id int) error {
	s, err := DB.Exec(`DELETE FROM locations WHERE location_id = ?`, id)
	if err != nil {
		return err
	}
	if r, _ := s.RowsAffected(); r == 0 {
		return errors.New("location id does not exist")
	}
	return err
}

func UpdateLocationEntry(location models.Location) error {
	query := `
	UPDATE locations
	SET location_name = ?
	WHERE location_id = ?
	`
	s, err := DB.Exec(query, location.LocationName, location.LocationID)
	if err != nil {
		return err
	}
	if r, _ := s.RowsAffected(); r == 0 {
		return errors.New("location entry does not exist")
	}
	return err
}
