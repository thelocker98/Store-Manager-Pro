// Info Popup
function moveDepartment(id) {
  axios.get(`/api/items/${id}`).then((res) => {
    const item = res.data;
    if (!item) return;

    let html = `<h2>${item.name} (${item.brand})</h2>`;

    if (item.department_name != "") {
      html += `<strong>Current Department:</strong> ${item.department_name}<br />`;
    }

    html += `<br/>
            <div style="justify-content: space-between; display: flex;">
              <label for="newDepartmentSelector">New Department</label>
              <select id="newDepartmentSelector" style="width: 175px; padding: 0px 5px; margin-left: 10px; border-radius: 5px;">
              </select>
            </div>

            <br/>

            <div id="newLocationdiv" style="justify-content: space-between; display: flex;">
              <label for="newLocationSelector">New Location</label>
              <select id="newLocationSelector" style="width: 175px; padding: 0px 5px; margin-left: 10px; border-radius: 5px;">
              </select>
            </div>

            <p id="departmentError" class="errorText" ></p>
    `;
    // Apply the Html
    document.getElementById("infoContent").innerHTML = html;

    // Populate the Department Selector
    loadDepartmentChangerSelector(item.department_id);

    // Create a listener for the new department selector
    const newDepartmentSelector = document.getElementById("newDepartmentSelector");
    newDepartmentSelector.addEventListener("change", (e) => {
      loadLocationsForDepartment(e.target.value);
    });

    // Hide buttons
    document.getElementById("infoEditBtn").style.display = "none";
    document.getElementById("infoDeleteBtn").style.display = "none";
    document.getElementById("infoEditBtn").style.display = "none";

    document.getElementById("infoChangeDepartmenBtn").style.display = "inline-block";
    document.getElementById("infoChangeDepartmenBtn").textContent = "Back";
    document.getElementById("infoChangeDepartmenBtn").onclick = () => openInfo(item.item_id);

    // Location Selector and Submit Button
    document.getElementById("newLocationdiv").style.display = "none";
    document.getElementById("submitDepartmentChange").style.display = "none";
    document.getElementById("submitDepartmentChange").onclick = () =>
      submitDepartmentChange(item.item_id);

    document.getElementById("infoPopup").style.display = "block";
    document.getElementById("infooverlay").style.display = "block";
  });
}

function loadDepartmentChangerSelector(departmentID) {
  return axios.get("/api/departments").then((res) => {
    const select = document.getElementById("newDepartmentSelector");
    select.innerHTML = "";

    var count = 0;
    var id = 0;

    if (Array.isArray(res.data)) {
      res.data.forEach((v) => {
        if (v.department_id != departmentID) {
          count++;
          const opt = document.createElement("option");
          id = v.department_id;
          opt.value = v.department_id;
          opt.text = v.department_name;
          select.appendChild(opt);
        }
      });
    }

    // Select Default
    if (count == 1) {
      document.getElementById("newDepartmentSelector").value = id;
      loadLocationsForDepartment(id);
    } else {
      document.getElementById("newDepartmentSelector").value = -1;
    }
  });
}

// Load all the locations based on the Department
function loadLocationsForDepartment(departmentID) {
  // Show Location Selector and Submit Button
  document.getElementById("newLocationdiv").style.display = "flex";
  document.getElementById("submitDepartmentChange").style.display = "inline-block";

  // Get Data for locations
  return axios.get("/api/locations").then((res) => {
    const select = document.getElementById("newLocationSelector");
    select.innerHTML = "";

    if (res.data.length > 1) {
      const opt = document.createElement("option");
      opt.value = 1;
      opt.text = "Not Assigned";
      select.appendChild(opt);
    }

    if (Array.isArray(res.data)) {
      res.data.forEach((v) => {
        const opt = document.createElement("option");
        opt.value = v.location_id;
        opt.text = v.location_name;
        select.appendChild(opt);
      });
    }

    if (res.data.length == 1) {
      const opt = document.createElement("option");
      opt.value = 1;
      opt.text = "Not Assigned";
      select.appendChild(opt);
    }
  });
}

function submitDepartmentChange(itemID) {
  const itemsEntry = {
    location_id: parseInt(document.getElementById("newLocationSelector").value, 10),
    department_id: parseInt(document.getElementById("newDepartmentSelector").value, 10),
  };

  axios
    .put(`/api/items/department/${itemID}`, itemsEntry)
    .then(() => {
      closeInfo();
      reloadData();
    })
    .catch((err) => {
      document.getElementById("departmentError").style.display = "block";

      document.getElementById("departmentError").textContent = "Error Processing Your Request";
    });
}
