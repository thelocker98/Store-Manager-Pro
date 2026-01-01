package db

func createTables() {
	// barcode Table
	DB.Exec("PRAGMA foreign_keys = ON;")

	// query := `
	// CREATE TABLE IF NOT EXISTS catalog (
	// 	catalog_id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	upc TEXT,
	// 	invoice_number TEXT,
	// 	brand TEXT,
	// 	name TEXT
	// );
	// `

	// _, err := DB.Exec(query)
	// if err != nil {
	// 	panic(err)
	// }

	// Inventory Table
	query := `
	CREATE TABLE IF NOT EXISTS inventory (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		vendor_id INTEGER,
		upc TEXT,
		invoice_number TEXT,
		name TEXT,
		brand TEXT,
		description TEXT,
		price REAL NOT NULL,
		weighed BOOL NOT NULL,
		count INTEGER,
		arrived_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		soldout_at DATETIME,
		deleted BOOL DEFAULT FALSE,
		FOREIGN KEY(vendor_id) REFERENCES vendors(vendor_id)
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}

	// Vendor Table
	query = `
	CREATE TABLE IF NOT EXISTS vendors (
		vendor_id INTEGER PRIMARY KEY AUTOINCREMENT,
		vendor_name TEXT NOT NULL
	);
	`

	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}
}
