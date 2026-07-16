var GlobalDepartment = getCookie("defaultDepartment");

// Populate Department selector
function loadDepartmentToggle() {
  return axios.get("/api/departments").then((res) => {
    const select = document.getElementById("departmentPageSelector");
    select.innerHTML = "";

    const opt = document.createElement("option");
    opt.value = 0;
    opt.text = "All Departments";
    select.appendChild(opt);

    if (Array.isArray(res.data)) {
      res.data.forEach((v) => {
        const opt = document.createElement("option");
        opt.value = v.department_id;
        opt.text = v.department_name;
        select.appendChild(opt);
      });
    }

    // Select Default
    document.getElementById("departmentPageSelector").value = GlobalDepartment;
  });
}

function updateDepartment() {
  GlobalDepartment = Number(
    document.getElementById("departmentPageSelector").value,
  );

  document.cookie =
    "defaultDepartment=" + GlobalDepartment + "; max-age=315360000; path=/";
  window.location.reload();
}

document.addEventListener("DOMContentLoaded", async function () {
  GlobalDepartment = Number(getCookie("defaultDepartment"));

  const pageDepartmentSelector = document.getElementById(
    "departmentPageSelector",
  );

  if (pageDepartmentSelector) {
    pageDepartmentSelector.addEventListener("change", () => {
      updateDepartment();
    });
  }

  loadDepartmentToggle();
});
