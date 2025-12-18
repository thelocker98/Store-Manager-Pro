package db

func createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS inventory (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		upc TEXT NOT NULL,
		name TEXT NOT NULL,
		description TEXT,
		price REAL NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
}
