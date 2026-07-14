PRAGMA foreign_keys = OFF;


--------------------------------------------------
-- 1. Rename old tables
--------------------------------------------------
ALTER TABLE locations RENAME TO locations_old;
ALTER TABLE inventory RENAME TO inventory_old;

--------------------------------------------------
-- 2. Create new locations table
--------------------------------------------------
CREATE TABLE locations (
    location_id INTEGER PRIMARY KEY AUTOINCREMENT,
    location_name TEXT NOT NULL UNIQUE
);

--------------------------------------------------
-- 3. Insert unique locations
--------------------------------------------------
INSERT INTO locations (location_name)
SELECT DISTINCT location_name
FROM locations_old;

--------------------------------------------------
-- 4. Create new inventory table
--------------------------------------------------
CREATE TABLE inventory (
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

--------------------------------------------------
-- 5. Copy inventory with remapped location IDs
--------------------------------------------------
INSERT INTO inventory (
    id,
    vendor_id,
    location_id,
    department_id,
    upc,
    invoice_number,
    name,
    brand,
    description,
    price,
    weighed,
    count,
    arrived_at,
    soldout_at,
    deleted
)
SELECT
    inv.id,
    inv.vendor_id,
    newloc.location_id,
    inv.department_id,
    inv.upc,
    inv.invoice_number,
    inv.name,
    inv.brand,
    inv.description,
    inv.price,
    inv.weighed,
    inv.count,
    inv.arrived_at,
    inv.soldout_at,
    inv.deleted
FROM inventory_old inv
JOIN locations_old oldloc
    ON inv.location_id = oldloc.location_id
JOIN locations newloc
    ON oldloc.location_name = newloc.location_name;

--------------------------------------------------
-- 6. Drop old tables
--------------------------------------------------
DROP TABLE inventory_old;
DROP TABLE locations_old;



PRAGMA foreign_keys = ON;
