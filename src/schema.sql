CREATE TABLE IF NOT EXISTS vendors (
		vendor_id SERIAL PRIMARY KEY,
		vendor_name TEXT NOT NULL UNIQUE
	);

CREATE TABLE IF NOT EXISTS departments (
		department_id SERIAL PRIMARY KEY,
		department_name TEXT NOT NULL UNIQUE
	);

CREATE TABLE IF NOT EXISTS lists (
    list_id SERIAL PRIMARY KEY,
    list_name TEXT NOT NULL UNIQUE,
    department_id INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (department_id)
        REFERENCES departments(department_id)
);

CREATE TABLE IF NOT EXISTS invoices (
			invoice_id SERIAL PRIMARY KEY,
			invoice_name TEXT NOT NULL UNIQUE,
			invoice_type TEXT NOT NULL,
			department_id INTEGER NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(department_id) REFERENCES departments(department_id)
		);

CREATE TABLE IF NOT EXISTS invoice_data (
			invoice_data_id SERIAL PRIMARY KEY,
			invoice_id INTEGER NOT NULL,
			rawtext TEXT NOT NULL,
			upc TEXT NOT NULL,
			details TEXT NOT NULL,
			qty INTEGER NOT NULL,
			item_cost DOUBLE PRECISION NOT NULL,
			total_cost DOUBLE PRECISION NOT NULL,
			discount_cost DOUBLE PRECISION NOT NULL,
			true_cost DOUBLE PRECISION NOT NULL,
			FOREIGN KEY(invoice_id) REFERENCES invoices(invoice_id) ON DELETE CASCADE,
			UNIQUE (invoice_data_id, invoice_id)
		);

CREATE TABLE IF NOT EXISTS ocr_files (
		    file_id          SERIAL PRIMARY KEY,
		    invoice_id       INTEGER NOT NULL,
		    filepath         TEXT NOT NULL UNIQUE,
		    status           INTEGER NOT NULL DEFAULT 0,
		    entries_added    INTEGER NOT NULL DEFAULT 0,
		    entries_failed   INTEGER NOT NULL DEFAULT 0,
		    error_message    TEXT NOT NULL DEFAULT '',
		    started_at       TIMESTAMPTZ,
		    completed_at     TIMESTAMPTZ,
		    created_at       TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    FOREIGN KEY (invoice_id) REFERENCES invoices(invoice_id) ON DELETE CASCADE
		);

CREATE TABLE IF NOT EXISTS locations (
    location_id SERIAL PRIMARY KEY,
    location_name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS inventory (
    id SERIAL PRIMARY KEY,
    vendor_id INTEGER DEFAULT 1,
    location_id INTEGER DEFAULT 1,
    department_id INTEGER NOT NULL,
    upc TEXT NOT NULL,
    invoice_number TEXT NOT NULL,
    name TEXT NOT NULL,
    brand TEXT NOT NULL,
    description TEXT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    weighed BOOLEAN NOT NULL,
    count INTEGER NOT NULL,
    arrived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    soldout_at TIMESTAMPTZ,
    deleted BOOLEAN DEFAULT FALSE,

    FOREIGN KEY(vendor_id) REFERENCES vendors(vendor_id),
    FOREIGN KEY(location_id) REFERENCES locations(location_id),
    FOREIGN KEY(department_id) REFERENCES departments(department_id)
);

CREATE TABLE IF NOT EXISTS list_data (
    list_entry_id SERIAL PRIMARY KEY,
    list_id INTEGER NOT NULL,
    item_id INTEGER NOT NULL,
    list_count INTEGER NOT NULL,
    FOREIGN KEY(list_id) REFERENCES lists(list_id) ON DELETE CASCADE,
    FOREIGN KEY(item_id) REFERENCES inventory(id),
    UNIQUE (list_id, item_id)
);
