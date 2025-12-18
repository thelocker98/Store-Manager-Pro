package db

import "gitea.locker98.com/locker98/Store-Manager-Pro/models"

func AddItem(item models.Item) error {
	query := `
	INSERT INTO inventory (upc, name, description, price)
	VALUES (?, ?, ?, ?)
	`
	_, err := DB.Exec(query, item.UPC, item.Name, item.Description, item.Price)
	return err
}

func GetAllItems() ([]models.Item, error) {
	rows, err := DB.Query(`SELECT id, upc, name, description, price FROM inventory`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var i models.Item
		rows.Scan(&i.ID, &i.UPC, &i.Name, &i.Description, &i.Price)
		items = append(items, i)
	}
	return items, nil
}

func DeleteItem(id int) error {
	_, err := DB.Exec(`DELETE FROM inventory WHERE id = ?`, id)
	return err
}
