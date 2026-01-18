// Load all Departments
function loadDepartments() {
  axios.get("/api/departments").then((res) => {
    const tbody = document.querySelector("#departmentsTable tbody");
    tbody.innerHTML = "";

    if (!res.data) {
      return;
    }

    res.data.forEach((departments) => {
      tbody.innerHTML += `
          <tr onclick="editDepartment(${departments.department_id})" class="avalible">
              <td>${departments.department_name}</td>
          </tr>
      `;
    });
  });
}

// Add Department
function addDepartment(e) {
  e.preventDefault();
  const departments = {
    department_name: document.getElementById("department_name").value,
  };
  axios
    .post("/api/department", departments)
    .then(() => {
      loadDepartments();
      document.getElementById("departmentaddError").style.display = "none";
      document.getElementById("departmentaddError").textContent = "";
      document.getElementById("addDepartmentForm").reset();
    })
    .catch((err) => {
      document.getElementById("departmentaddError").style.display = "block";
      document.getElementById("departmentaddError").textContent =
        "Error Processing Your Request";
    });
}

// Delete Departments
function deleteDepartment(id) {
  if (confirm("Are you sure you want to delete this department?")) {
    axios
      .delete(`/api/department/${id}`)
      .then(() => {
        closeDepartmentEdit();
        loadDepartments();
      })
      .catch((err) => {
        document.getElementById("departmentinfoError").style.display = "block";
        var response = String(err.response.data.error).toLowerCase();

        if (response.includes("foreign key")) {
          document.getElementById("departmentinfoError").textContent =
            "Some items still reference this department";
        } else {
          document.getElementById("departmentinfoError").textContent =
            "Could Not Delete Department";
        }
      });
  }
}

// Initial load
loadDepartments();

// Open the edit popup and fill fields
function editDepartment(id) {
  axios
    .get("/api/departments")
    .then((res) => {
      const department = res.data.find((i) => i.department_id === id);
      if (!department) return;

      document.getElementById("editDepartmentId").value =
        department.department_id;
      document.getElementById("editDepartmentName").value =
        department.department_name;
      document.getElementById("departmentDeleteButton").onclick = () => {
        deleteDepartment(department.department_id);
      };

      document.getElementById("editDepartmentPopup").style.display = "block";
      document.getElementById("overlay").style.display = "block";
    })
    .catch((err) => {
      document.getElementById("departmentinfoError").style.display = "block";
      document.getElementById("departmentinfoError").textContent =
        "Error Processing Your Request";
    });
}

// Close the popup
function closeDepartmentEdit() {
  document.getElementById("departmentinfoError").style.display = "none";
  document.getElementById("departmentinfoError").textContent = "";
  document.getElementById("editDepartmentPopup").style.display = "none";
  document.getElementById("overlay").style.display = "none";
}

// Submit edit form
function submitDepartmentEdit(e) {
  e.preventDefault();
  const id = document.getElementById("editDepartmentId").value;
  const department = {
    department_name: document.getElementById("editDepartmentName").value,
  };

  axios
    .put(`/api/department/${id}`, department)
    .then(() => {
      closeDepartmentEdit();
      document.getElementById("addDepartmentForm").reset();
      loadDepartments();
    })
    .catch((err) => {
      document.getElementById("departmentinfoError").style.display = "block";
      document.getElementById("departmentinfoError").textContent =
        "Error Processing Your Request";
    });
}
