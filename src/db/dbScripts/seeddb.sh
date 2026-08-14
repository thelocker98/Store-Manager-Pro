#!/usr/bin/env bash
set -euo pipefail

DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-storemanager}"
DB_NAME="${DB_NAME:-storemanager}"
DB_PASSWORD="${DB_PASSWORD:-changeme}"

export PGPASSWORD="$DB_PASSWORD"
PSQL="psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -v ON_ERROR_STOP=1 -t -A"

pick() {
    local arr=("$@")
    echo "${arr[$((RANDOM % ${#arr[@]}))]}"
}

sqlq() {
    local s="$1"
    s="${s//\'/\'\'}"
    echo "$s"
}

# --- Departments -----------------------------------------------------------
$PSQL -c "INSERT INTO departments (department_name) VALUES ('Freezer'), ('Refrigerator') ON CONFLICT (department_name) DO NOTHING;"

# --- Locations --------------------------------------------------------------
$PSQL -c "INSERT INTO locations (location_name) VALUES ('Refrigerator'), ('T1'), ('T2') ON CONFLICT (location_name) DO NOTHING;"

# --- Vendors ----------------------------------------------------------------
$PSQL -c "INSERT INTO vendors (vendor_name) VALUES ('G&C'), ('Natural Choice'), ('Hornings'), ('Denver Cold Storage'), ('Dug') ON CONFLICT (vendor_name) DO NOTHING;"

# --- Look up IDs by name ----------------------------------------------------
dept_freezer=$($PSQL -c "SELECT department_id FROM departments WHERE department_name = 'Freezer';")
dept_refrig=$($PSQL -c "SELECT department_id FROM departments WHERE department_name = 'Refrigerator';")

mapfile -t location_ids < <($PSQL -c "SELECT location_id FROM locations WHERE location_name IN ('Refrigerator','T1','T2') ORDER BY location_id;")
mapfile -t vendor_ids < <($PSQL -c "SELECT vendor_id FROM vendors WHERE vendor_name IN ('G&C','Natural Choice','Hornings','Denver Cold Storage','Dug') ORDER BY vendor_id;")

# --- Items ------------------------------------------------------------------
# Format: department|name|upc|price|weighed|brand|description
items=(
"Freezer|TicTac min candies|8005012|1.29|true|Ferrero|Mint candy in a small plastic container"
"Freezer|nutella chocolate spread|8013546|5.95|false|Ferrero|Chocolate hazelnut spread"
"Freezer|Mollers dobbel fish oil capsules|7011242|25.09|false|Moller|Fish oil omega-3 capsules"
"Freezer|Camel Candy|7011242|4.49|false|Camel|Hard candy assortment"
"Refrigerator|Colloring Pens|04252891217|4.25|false|Crayola|Set of coloring pens"
"Refrigerator|Polaroid Photo film|07410014661|1.59|false|Polaroid|Instant photo film pack"
"Refrigerator|Olympus Photo Camera|05033212701|9.29|false|Olympus|35mm point-and-shoot film camera"
"Refrigerator|Contact Lenses|78581010215|3.29|false|Acuvue|Contact lenses"
)

for item in "${items[@]}"; do
    IFS='|' read -r dept name upc price weighed brand description <<< "$item"

    if [[ "$dept" == "Freezer" ]]; then
        dep_id="$dept_freezer"
    else
        dep_id="$dept_refrig"
    fi

    vendor_id="$(pick "${vendor_ids[@]}")"
    location_id="$(pick "${location_ids[@]}")"
    count=$((RANDOM % 50 + 1))

    $PSQL -c "INSERT INTO inventory (vendor_id, location_id, department_id, upc, invoice_number, name, brand, description, price, weighed, count)
    VALUES ($vendor_id, $location_id, $dep_id, '$(sqlq "$upc")', '', '$(sqlq "$name")', '$(sqlq "$brand")', '$(sqlq "$description")', $price, $weighed, $count);"

    echo "  + $name  (dept=$dept, vendor=$vendor_id, location=$location_id, qty=$count, \$"$price")"
done

echo "Seed complete."
