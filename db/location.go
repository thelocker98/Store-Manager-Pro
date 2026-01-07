package db

import (
	"errors"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddLocationEntry(location models.Location) error {
	query := `
	INSERT INTO locations (location_name)
	VALUES (?)
	`
	_, err := DB.Exec(query, location.LocationName)
	return err
}

func GetAllLocations() ([]models.Location, error) {
	rows, err := DB.Query(`SELECT location_id, location_name FROM locations WHERE location_id != 1 ORDER BY location_name ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []models.Location
	for rows.Next() {
		var v models.Location
		rows.Scan(&v.LocationID, &v.LocationName)
		locations = append(locations, v)
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
