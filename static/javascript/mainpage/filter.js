// variables
var Vendor = 0;
var Location = 0;
var Department = 0;
var Listid = 0;
var OrderBy = "upc";
var ShowDeleted = true;

const VendorLabels = new Map();
const LocationLabels = new Map();
const DepartmentLabels = new Map();
const ListLabels = new Map();

function applyFilter() {
  const filterDiv = document.getElementById("filterSection");
  if (!filterDiv) return;

  const vendorFilter = document.getElementById("vendorSelector");
  const locationFilter = document.getElementById("locationSelector");
  const listFilter = document.getElementById("listSelector");
  const orderByFilter = document.getElementById("orderBySelector");
  const showDeletedFilter = document.getElementById("showDeletedCheckbox");

  Vendor = vendorFilter.value || 0;
  Location = locationFilter.value || 0;
  Listid = listFilter.value || 0;
  OrderBy = orderByFilter.value || 0;
  ShowDeleted = showDeletedFilter.checked || 0;

  filterDiv.innerHTML = "";
  var html = "";

  if (Vendor != 0) {
    html += `<span class="filters" onclick="removeFilter('vendor')">${VendorLabels.get(Number(Vendor))}</span>`;
  }
  if (Location != 0) {
    html += `<span class="filters" onclick="removeFilter('location')">${LocationLabels.get(Number(Location))}</span>`;
  }
  if (Listid != 0) {
    html += `<span class="filters" onclick="removeFilter('list')">${ListLabels.get(Number(Listid))}</span>`;
  }
  if (ShowDeleted != 0) {
    html += `<span class="filters" onclick="removeFilter('showdeleted')">Show Deleted</span>`;
  }

  filterDiv.innerHTML = html;
  reloadData();
}

function removeFilter(filterKey) {
  if (filterKey == "vendor") {
    document.getElementById("vendorSelector").value = 0;
  }
  if (filterKey == "location") {
    document.getElementById("locationSelector").value = 0;
  }
  if (filterKey == "list") {
    document.getElementById("listSelector").value = 0;
  }
  if (filterKey == "showdeleted") {
    document.getElementById("showDeletedCheckbox").checked = 0;
  }

  applyFilter(); // re-render filters
}

async function loadFiltersDropdown() {
  const [vendorsRes, locationsRes, listsRes] = await Promise.all([
    axios.get("/api/vendors"),
    axios.get("/api/locations/" + String(GlobalDepartment)),
    axios.get("/api/lists/d/" + String(GlobalDepartment)),
  ]);

  // Vendors
  {
    const select = document.getElementById("vendorSelector");
    select.innerHTML = "";

    select.append(new Option("All", 0));
    VendorLabels.set(0, "All");
    select.append(new Option("Not Assigned To Vendor", 1));
    VendorLabels.set(1, "Not Assigned To Vendor");

    if (Array.isArray(vendorsRes.data)) {
      vendorsRes.data.forEach((v) => {
        select.append(new Option(v.vendor_name, v.vendor_id));
        VendorLabels.set(v.vendor_id, v.vendor_name);
      });
    }

    select.value = Vendor;
  }

  // Locations
  {
    const select = document.getElementById("locationSelector");
    select.innerHTML = "";

    select.append(new Option("All", 0));
    LocationLabels.set(0, "All");
    select.append(new Option("Not Assigned To Location", 1));
    LocationLabels.set(1, "Not Assigned To Location");

    if (Array.isArray(locationsRes.data)) {
      locationsRes.data.forEach((l) => {
        select.append(new Option(l.location_name, l.location_id));
        LocationLabels.set(l.location_id, l.location_name);
      });
    }

    select.value = Location;
  }

  // Lists
  {
    const select = document.getElementById("listSelector");
    select.innerHTML = "";

    select.append(new Option("All", 0));
    ListLabels.set(0, "All");

    if (Array.isArray(listsRes.data)) {
      listsRes.data.forEach((l) => {
        select.append(new Option(l.list_name, l.list_id));
        ListLabels.set(l.list_id, l.list_name);
      });
    }

    select.value = Listid;
  }
}
