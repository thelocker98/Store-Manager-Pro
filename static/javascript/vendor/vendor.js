// Load all Vendors
function loadVendors() {
  axios.get("/api/vendors").then((res) => {
    const tbody = document.querySelector("#vendorsTable tbody");
    tbody.innerHTML = "";
    res.data.forEach((vendors) => {
      tbody.innerHTML += `
          <tr onclick="editVendor(${vendors.vendor_id})" class="avalible">
              <td>${vendors.vendor_name}</td>
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
  axios
    .post("/api/vendors", vendors)
    .then(() => {
      loadVendors();
      document.getElementById("vendoraddError").style.display = "none";
      document.getElementById("vendoraddError").textContent = "";
      document.getElementById("addVendorForm").reset();
    })
    .catch((err) => {
      document.getElementById("vendoraddError").style.display = "block";
      document.getElementById("vendoraddError").textContent =
        "Error Processing Your Request";
    });
}

// Delete Vendors
function deleteVendor(id) {
  if (confirm("Are you sure you want to delete this vendor?")) {
    axios
      .delete(`/api/vendors/${id}`)
      .then(() => {
        closeVendorEdit();
        loadVendors();
      })
      .catch((err) => {
        document.getElementById("vendorinfoError").style.display = "block";
        var response = String(err.response.data.error).toLowerCase();

        if (response.includes("foreign key")) {
          document.getElementById("vendorinfoError").textContent =
            "Some items still reference this vendor";
        } else {
          document.getElementById("vendorinfoError").textContent =
            "Could Not Delete Vendor";
        }
      });
  }
}

// Initial load
loadVendors();

// Open the edit popup and fill fields
function editVendor(id) {
  axios
    .get("/api/vendors")
    .then((res) => {
      const vendor = res.data.find((i) => i.vendor_id === id);
      if (!vendor) return;

      document.getElementById("editVendorId").value = vendor.vendor_id;
      document.getElementById("editVendorName").value = vendor.vendor_name;
      document.getElementById("vendorDeleteButton").onclick = () => {
        deleteVendor(vendor.vendor_id);
      };

      document.getElementById("editVendorPopup").style.display = "block";
      document.getElementById("overlay").style.display = "block";
    })
    .catch((err) => {
      document.getElementById("vendorinfoError").style.display = "block";
      document.getElementById("vendorinfoError").textContent =
        "Error Processing Your Request";
    });
}

// Close the popup
function closeVendorEdit() {
  document.getElementById("vendorinfoError").style.display = "none";
  document.getElementById("vendorinfoError").textContent = "";
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

  axios
    .put(`/api/vendors/${id}`, vendor)
    .then(() => {
      closeVendorEdit();
      document.getElementById("addVendorForm").reset();
      loadVendors();
    })
    .catch((err) => {
      document.getElementById("vendorinfoError").style.display = "block";
      document.getElementById("vendorinfoError").textContent =
        "Error Processing Your Request";
    });
}
