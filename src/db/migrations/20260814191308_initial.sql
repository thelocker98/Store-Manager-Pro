-- +goose Up
CREATE TABLE departments (
    department_id serial PRIMARY KEY,
    department_name text NOT NULL UNIQUE
);
INSERT INTO departments (department_name)
	VALUES ('N/A')
	ON CONFLICT (department_name) DO NOTHING;


CREATE TABLE vendors (
    vendor_id serial PRIMARY KEY,
    vendor_name text NOT NULL UNIQUE
);
INSERT INTO vendors (vendor_name)
	VALUES ('N/A')
	ON CONFLICT (vendor_name) DO NOTHING;



CREATE TABLE locations (
    location_id serial PRIMARY KEY,
    location_name text NOT NULL UNIQUE
);
INSERT INTO locations (location_name)
	VALUES ('N/A')
	ON CONFLICT (location_name) DO NOTHING;


CREATE TABLE inventory (
    id serial PRIMARY KEY,
    vendor_id integer NOT NULL DEFAULT 1 REFERENCES vendors(vendor_id),
    location_id integer NOT NULL DEFAULT 1 REFERENCES locations(location_id),
    department_id integer NOT NULL REFERENCES departments(department_id),
    upc text NOT NULL,
    invoice_number text NOT NULL,
    name text NOT NULL,
    brand text NOT NULL,
    description text NOT NULL,
    price double precision NOT NULL,
    weighed boolean NOT NULL,
    count integer NOT NULL,
    arrived_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    soldout_at timestamp without time zone,
    deleted boolean NOT NULL DEFAULT false
);

CREATE TABLE invoices (
    invoice_id serial PRIMARY KEY,
    invoice_name text NOT NULL UNIQUE,
    invoice_type text NOT NULL,
    department_id integer NOT NULL REFERENCES departments(department_id),
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE invoice_data (
    invoice_data_id serial PRIMARY KEY,
    invoice_id integer NOT NULL REFERENCES invoices(invoice_id) ON DELETE CASCADE,
    rawtext text NOT NULL,
    upc text NOT NULL,
    details text NOT NULL,
    qty integer NOT NULL,
    item_cost double precision NOT NULL,
    total_cost double precision NOT NULL,
    discount_cost double precision NOT NULL,
    true_cost double precision NOT NULL,
    UNIQUE (invoice_data_id, invoice_id)
);

CREATE TABLE lists (
    list_id serial PRIMARY KEY,
    list_name text NOT NULL UNIQUE,
    department_id integer NOT NULL REFERENCES departments(department_id),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE list_data (
    list_entry_id serial PRIMARY KEY,
    list_id integer NOT NULL REFERENCES lists(list_id) ON DELETE CASCADE,
    item_id integer NOT NULL REFERENCES inventory(id),
    list_count integer NOT NULL,
    UNIQUE (list_id, item_id)
);

CREATE TABLE ocr_files (
    file_id serial PRIMARY KEY,
    invoice_id integer NOT NULL REFERENCES invoices(invoice_id) ON DELETE CASCADE,
    filepath text NOT NULL UNIQUE,
    status integer NOT NULL DEFAULT 0,
    entries_added integer NOT NULL DEFAULT 0,
    entries_failed integer NOT NULL DEFAULT 0,
    error_message text NOT NULL DEFAULT '',
    started_at timestamp without time zone,
    completed_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE ocr_files;
DROP TABLE list_data;
DROP TABLE lists;
DROP TABLE invoice_data;
DROP TABLE invoices;
DROP TABLE inventory;
DROP TABLE locations;
DROP TABLE vendors;
DROP TABLE departments;
