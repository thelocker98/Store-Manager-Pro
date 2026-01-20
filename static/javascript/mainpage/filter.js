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

  listFilter.options[listFilter.value].textContent;

  reloadData();
}

function removeFilter(filterKey) {
  applyFilter(); // re-render filters
  console.log("test");
}

function loadFiltersDropdown() {
  axios.get("/api/vendors").then((res) => {
    const select = document.getElementById("vendorSelector");
    select.innerHTML = "";

    var opt = document.createElement("option");
    opt.value = 0;
    opt.text = "All";
    select.appendChild(opt);

    opt = document.createElement("option");
    opt.value = 1;
    opt.text = "Not Assigned";
    select.appendChild(opt);

    res.data.forEach((v) => {
      const opt = document.createElement("option");
      opt.value = v.vendor_id;
      opt.text = v.vendor_name;
      select.appendChild(opt);
    });
    select.value = Vendor;
  });

  axios.get("/api/locations").then((res) => {
    const select = document.getElementById("locationSelector");
    select.innerHTML = "";

    var opt = document.createElement("option");
    opt.value = 0;
    opt.text = "All";
    select.appendChild(opt);

    opt = document.createElement("option");
    opt.value = 1;
    opt.text = "Not Assigned";
    select.appendChild(opt);

    res.data.forEach((l) => {
      const opt = document.createElement("option");
      opt.value = l.location_id;
      opt.text = l.location_name;
      select.appendChild(opt);
    });
    select.value = Location;
  });

  axios.get("/api/departments").then((res) => {
    const select = document.getElementById("departmentSelector");
    select.innerHTML = "";

    var opt = document.createElement("option");
    opt.value = 0;
    opt.text = "All";
    select.appendChild(opt);

    opt = document.createElement("option");
    opt.value = 1;
    opt.text = "Not Assigned";
    select.appendChild(opt);

    res.data.forEach((d) => {
      const opt = document.createElement("option");
      opt.value = d.department_id;
      opt.text = d.department_name;
      select.appendChild(opt);
    });
    select.value = Department;
  });

  axios.get("/api/lists").then((res) => {
    const select = document.getElementById("listSelector");
    select.innerHTML = "";

    var opt = document.createElement("option");
    opt.value = 0;
    opt.text = "All";
    select.appendChild(opt);

    res.data.forEach((l) => {
      const opt = document.createElement("option");
      opt.value = l.list_id;
      opt.text = l.list_name;
      select.appendChild(opt);
    });
    select.value = Listid;
  });
}
