-- +goose Up
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS inventory (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	vendor_id INTEGER DEFAULT 1,
	location_id INTEGER DEFAULT 1,
	department_id INTEGER NOT NULL,
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
    FOREIGN KEY(location_id) REFERENCES locations(location_id),
	FOREIGN KEY(department_id) REFERENCES departments(department_id)
);

CREATE TABLE IF NOT EXISTS vendors (
	vendor_id INTEGER PRIMARY KEY AUTOINCREMENT,
	vendor_name TEXT NOT NULL UNIQUE
);

INSERT OR IGNORE INTO vendors (vendor_id, vendor_name)
VALUES (1, 'N/A');

CREATE TABLE IF NOT EXISTS departments (
	department_id INTEGER PRIMARY KEY AUTOINCREMENT,
	department_name TEXT NOT NULL UNIQUE
);

INSERT OR IGNORE INTO departments (department_id, department_name)
VALUES (1, 'N/A');


CREATE TABLE IF NOT EXISTS locations (
	location_id INTEGER PRIMARY KEY AUTOINCREMENT,
	location_name TEXT NOT NULL UNIQUE
);

INSERT OR IGNORE INTO locations (location_id, location_name)
VALUES (1, 'N/A');


CREATE TABLE IF NOT EXISTS lists (
	list_id INTEGER PRIMARY KEY AUTOINCREMENT,
	list_name TEXT NOT NULL UNIQUE,
	department_id INTEGER NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(department_id) REFERENCES departments(department_id)
);

CREATE TABLE IF NOT EXISTS list_data (
	list_entry_id INTEGER PRIMARY KEY AUTOINCREMENT,
	list_id INTEGER,
	item_id INTEGER,
	list_count INTEGER NOT NULL,
	FOREIGN KEY(list_id) REFERENCES lists(list_id) ON DELETE CASCADE,
	FOREIGN KEY(item_id) REFERENCES inventory(id),
	UNIQUE (list_id, item_id)
);


CREATE TABLE IF NOT EXISTS invoices (
	invoice_id INTEGER PRIMARY KEY AUTOINCREMENT,
	invoice_name TEXT NOT NULL UNIQUE,
	invoice_type TEXT NOT NULL,
	department_id INTEGER NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(department_id) REFERENCES departments(department_id)
);


CREATE TABLE IF NOT EXISTS invoice_data (
	invoice_data_id INTEGER PRIMARY KEY AUTOINCREMENT,
	invoice_id INTEGER NOT NULL,
	rawtext STRING NOT NULL,
	upc STRING,
	details STRING,
	qty INTEGER,
	item_cost FLOAT,
	total_cost FLOAT,
	discount_cost FLOAT,
	true_cost FLOAT,
	FOREIGN KEY(invoice_id) REFERENCES invoices(invoice_id) ON DELETE CASCADE,
	UNIQUE (invoice_data_id, invoice_id)
);


CREATE TABLE IF NOT EXISTS ocr_files (
    file_id          INTEGER PRIMARY KEY AUTOINCREMENT,
    invoice_id       INTEGER NOT NULL,
    filepath         TEXT NOT NULL UNIQUE,
    status           INTEGER DEFAULT 0,
    entries_added    INTEGER DEFAULT 0,
    entries_failed   INTEGER DEFAULT 0,
    error_message    TEXT DEFAULT '',
    started_at       DATETIME,
    completed_at     DATETIME,
    created_at       DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (invoice_id) REFERENCES invoices(invoice_id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE IF EXISTS ocr_files;
DROP TABLE IF EXISTS invoice_data;
DROP TABLE IF EXISTS invoices;
DROP TABLE IF EXISTS list_data;
DROP TABLE IF EXISTS lists;
DROP TABLE IF EXISTS inventory;
DROP TABLE IF EXISTS locations;
DROP TABLE IF EXISTS departments;
DROP TABLE IF EXISTS vendors;
