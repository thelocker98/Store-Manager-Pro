package db

import (
	"fmt"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func AddItem(item models.Item) error {
	query := `
	INSERT INTO inventory (upc, brand, name, description, price)
	VALUES (?, ?, ?, ?, ?)
	`
	_, err := DB.Exec(query, item.UPC, item.Brand, item.Name, item.Description, item.Price)
	fmt.Println(err)
	return err
}

func GetAllItems() ([]models.Item, error) {
	rows, err := DB.Query(`SELECT id, upc, brand, name, description, price FROM inventory ORDER BY upc ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var i models.Item
		rows.Scan(&i.ID, &i.UPC, &i.Brand, &i.Name, &i.Description, &i.Price)
		items = append(items, i)
	}
	return items, nil
}

func DeleteItem(id int) error {
	_, err := DB.Exec(`DELETE FROM inventory WHERE id = ?`, id)
	return err
}

func UpdateItem(id int, item models.Item) error {
	query := `
	UPDATE inventory
	SET upc = ?, brand = ?, name = ?, description = ?, price = ?
	WHERE id = ?
	`
	_, err := DB.Exec(query, item.UPC, item.Brand, item.Name, item.Description, item.Price, id)
	return err
}
