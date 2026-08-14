// Load all Locations
function loadLocations() {
  axios.get("/api/locations").then((res) => {
    const tbody = document.querySelector("#locationsTable tbody");
    tbody.innerHTML = "";

    if (!Array.isArray(res.data)) return;

    res.data.forEach((locations) => {
      tbody.innerHTML += `
          <tr onclick="editLocation(${locations.location_id})" class="avalible">
            <td>
              <div style="display: flex; justify-content: space-between; align-items: center;">
                <span>${locations.location_name}</span>
                <span style="font-size: 15px; color: #555;">${locations.location_item_count} Items</span>
              </div>
            </td>
          </tr>
      `;
    });
  });
}

// Add Location
function addLocation(e) {
  e.preventDefault();
  const locations = {
    location_name: document.getElementById("location_name").value,
    department_id: GlobalDepartment,
  };
  axios
    .post("/api/locations", locations)
    .then(() => {
      loadLocations();
      document.getElementById("locationaddError").style.display = "none";
      document.getElementById("locationaddError").textContent = "";
      document.getElementById("addLocationForm").reset();
    })
    .catch((err) => {
      document.getElementById("locationaddError").style.display = "block";
      document.getElementById("locationaddError").textContent = "Error Processing Your Request";
    });
}

// Delete Locations
function deleteLocation(id) {
  if (confirm("Are you sure you want to delete this location?")) {
    axios
      .delete(`/api/locations/${id}`)
      .then(() => {
        closeLocationEdit();
        loadLocations();
      })
      .catch((err) => {
        document.getElementById("locationinfoError").style.display = "block";
        var response = String(err.response.data.error).toLowerCase();

        if (response.includes("foreign key")) {
          document.getElementById("locationinfoError").textContent =
            "Some items still reference this location";
        } else {
          document.getElementById("locationinfoError").textContent = "Could Not Delete Location";
        }
      });
  }
}

// Initial load
loadLocations();

// Open the edit popup and fill fields
function editLocation(id) {
  axios
    .get("/api/locations")
    .then((res) => {
      const location = res.data.find((i) => i.location_id === id);
      if (!location) return;

      document.getElementById("editLocationId").value = location.location_id;
      document.getElementById("editLocationName").value = location.location_name;
      document.getElementById("locationDeleteButton").onclick = () => {
        deleteLocation(location.location_id);
      };

      document.getElementById("editLocationPopup").style.display = "block";
      document.getElementById("overlay").style.display = "block";
    })
    .catch((err) => {
      document.getElementById("locationinfoError").style.display = "block";
      document.getElementById("locationinfoError").textContent = "Error Processing Your Request";
    });
}

// Close the popup
function closeLocationEdit() {
  document.getElementById("locationinfoError").style.display = "none";
  document.getElementById("locationinfoError").textContent = "";
  document.getElementById("editLocationPopup").style.display = "none";
  document.getElementById("overlay").style.display = "none";
}

// Submit edit form
function submitLocationEdit(e) {
  e.preventDefault();
  const id = document.getElementById("editLocationId").value;
  const location = {
    location_name: document.getElementById("editLocationName").value,
  };

  axios
    .put(`/api/locations/${id}`, location)
    .then(() => {
      closeLocationEdit();
      document.getElementById("addLocationForm").reset();
      loadLocations();
    })
    .catch((err) => {
      document.getElementById("locationinfoError").style.display = "block";
      document.getElementById("locationinfoError").textContent = "Error Processing Your Request";
    });
}
