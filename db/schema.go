package db

func createTables() {
	// barcode Table
	DB.Exec("PRAGMA foreign_keys = ON;")

	// Inventory Table
	query := `
	CREATE TABLE IF NOT EXISTS inventory (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		vendor_id INTEGER,
		location_id INTEGER,
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
		FOREIGN KEY(vendor_id) REFERENCES vendors(vendor_id),
		FOREIGN KEY(location_id) REFERENCES locations(location_id)
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
		vendor_name TEXT NOT NULL UNIQUE
	);
	`

	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}

	query = `
	INSERT OR IGNORE INTO vendors (vendor_id, vendor_name)
	VALUES (1, 'N/A');
	`

	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}

	// Location Table
	query = `
	CREATE TABLE IF NOT EXISTS locations (
		location_id INTEGER PRIMARY KEY AUTOINCREMENT,
		location_name TEXT NOT NULL UNIQUE
	);
	`
	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}

	query = `
	INSERT OR IGNORE INTO locations (location_id, location_name)
	VALUES (1, 'N/A');
	`

	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}

	// List Table
	query = `
		CREATE TABLE IF NOT EXISTS lists (
			list_id INTEGER PRIMARY KEY AUTOINCREMENT,
			list_name TEXT NOT NULL UNIQUE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
		`

	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}

	// List Entrys Table
	query = `
		CREATE TABLE IF NOT EXISTS list_data (
			list_entry_id INTEGER PRIMARY KEY AUTOINCREMENT,
			list_id INTEGER,
			item_id INTEGER,
			list_count INTEGER NOT NULL,
			FOREIGN KEY(list_id) REFERENCES lists(list_id) ON DELETE CASCADE,
			FOREIGN KEY(item_id) REFERENCES inventory(id),
			UNIQUE (list_id, item_id)
		);
		`

	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}
}
