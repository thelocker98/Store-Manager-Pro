// Load all Vendors
function loadVendors() {
  axios.get("/api/vendors").then((res) => {
    const tbody = document.querySelector("#vendorsTable tbody");
    tbody.innerHTML = "";
    res.data.forEach((vendors) => {
      tbody.innerHTML += `
                                <tr>
                                    <td>${vendors.vendor_name}</td>
                                    <td>
                                        <button onclick="editVendor(${vendors.vendor_id})">Edit</button>
                                        <button onclick="deleteVendor(${vendors.vendor_id})">Delete</button>
                                    </td>
                                </tr>
                            `;
    });
  });
}

// Add Vendor
function addVendor(e) {
  e.preventDefault();
  const vendors = {
    vendor_name: document.getElementById("vendor_name").value,
  };
  axios.post("/api/vendors", vendors).then(() => {
    loadVendors();
    document.getElementById("addForm").reset();
  });
}

// Delete Vendors
function deleteVendor(id) {
  if (confirm("Are you sure you want to delete this Vendor?")) {
    axios.delete(`/api/vendors/${id}`).then(() => loadVendors());
  }
}

// Initial load
loadVendors();

// Open the edit popup and fill fields
function editVendor(id) {
  axios.get("/api/vendors").then((res) => {
    const vendor = res.data.find((i) => i.vendor_id === id);
    if (!vendor) return;

    document.getElementById("editVendorId").value = vendor.vendor_id;
    document.getElementById("editVendorName").value = vendor.vendor_name;

    document.getElementById("editVendorPopup").style.display = "block";
    document.getElementById("overlay").style.display = "block";
  });
}

// Close the popup
function closeVendorEdit() {
  document.getElementById("editVendorPopup").style.display = "none";
  document.getElementById("overlay").style.display = "none";
}

// Submit edit form
function submitVendorEdit(e) {
  e.preventDefault();
  const id = document.getElementById("editVendorId").value;
  const vendor = {
    vendor_name: document.getElementById("editVendorName").value,
  };

  axios.put(`/api/vendors/${id}`, vendor).then(() => {
    closeVendorEdit();
    loadVendors();
  });
}
