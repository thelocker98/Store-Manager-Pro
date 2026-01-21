const VendorLabels = new Map();
const LocationLabels = new Map();
const DepartmentLabels = new Map();
const ListLabels = new Map();

function applyFilter() {
  const filterDiv = document.getElementById("filterSection");
  if (!filterDiv) return;

  const vendorFilter = document.getElementById("vendorSelector");
  const locationFilter = document.getElementById("locationSelector");
  const departmentFilter = document.getElementById("departmentSelector");
  const listFilter = document.getElementById("listSelector");
  const orderByFilter = document.getElementById("orderBySelector");
  const showDeletedFilter = document.getElementById("showDeletedCheckbox");

  Vendor = vendorFilter.value || 0;
  Location = locationFilter.value || 0;
  Department = departmentFilter.value || 0;
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
  if (Department != 0) {
    html += `<span class="filters" onclick="removeFilter('department')">${DepartmentLabels.get(Number(Department))}</span>`;
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
  console.log(filterKey);
  if (filterKey == "vendor") {
    document.getElementById("vendorSelector").value = 0;
  }
  if (filterKey == "location") {
    document.getElementById("locationSelector").value = 0;
  }
  if (filterKey == "department") {
    document.getElementById("departmentSelector").value = 0;
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
  const [vendorsRes, locationsRes, departmentsRes, listsRes] =
    await Promise.all([
      axios.get("/api/vendors"),
      axios.get("/api/locations"),
      axios.get("/api/departments"),
      axios.get("/api/lists"),
    ]);

  // Vendors
  {
    const select = document.getElementById("vendorSelector");
    select.innerHTML = "";

    select.append(new Option("All", 0));
    VendorLabels.set(0, "All");
    select.append(new Option("Not Assigned To Vendor", 1));
    VendorLabels.set(1, "Not Assigned To Vendor");

    vendorsRes.data.forEach((v) => {
      select.append(new Option(v.vendor_name, v.vendor_id));
      VendorLabels.set(v.vendor_id, v.vendor_name);
    });

    select.value = Vendor;
  }

  // Locations
  {
    const select = document.getElementById("locationSelector");
    select.innerHTML = "";

    select.append(new Option("All", 0));
    LocationLabels.set(0, "All");
    select.append(new Option("Not Assigned To List", 1));
    LocationLabels.set(1, "Not Assigned To List");

    locationsRes.data.forEach((l) => {
      select.append(new Option(l.location_name, l.location_id));
      LocationLabels.set(l.location_id, l.location_name);
    });

    select.value = Location;
  }

  // Departments
  {
    const select = document.getElementById("departmentSelector");
    select.innerHTML = "";

    select.append(new Option("All", 0));
    DepartmentLabels.set(0, "All");
    select.append(new Option("Not Assigned To Department", 1));
    DepartmentLabels.set(1, "Not Assigned To Department");

    departmentsRes.data.forEach((d) => {
      select.append(new Option(d.department_name, d.department_id));
      DepartmentLabels.set(d.department_id, d.department_name);
    });

    select.value = Department;
  }

  // Lists
  {
    const select = document.getElementById("listSelector");
    select.innerHTML = "";

    select.append(new Option("All", 0));
    ListLabels.set(0, "All");

    listsRes.data.forEach((l) => {
      select.append(new Option(l.list_name, l.list_id));
      ListLabels.set(l.list_id, l.list_name);
    });

    select.value = Listid;
  }
}
