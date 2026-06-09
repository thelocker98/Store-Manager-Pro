function loadInvoiceName() {
  const display = document.getElementById("invoiceName");

  // Load Invoice Name from Database
  axios
    .get(`/api/invoices/${InvoiceID}`)
    .then((res) => {
      display.textContent = res.data.invoice_name;

      // if department_id is not equal to all departments and it is in the wrong department relative to the invoice then exit to the invoicepage
      if (res.data.department_id != GlobalDepartment && GlobalDepartment != 0) {
        window.location.href = "/invoices";
      }

      if (GlobalDepartment == 0) {
        window.location.href = "/invoices";
      }
    })
    .catch((err) => {
      console.log("error updating invoice name in database");
    });
}

async function loadInvoice() {
  const tbody = document.querySelector("#invoiceTable tbody");

  const res = await axios.get(`/api/invoiceentrys/${InvoiceID}?page=${page}&pagesize=${pageSize}`);
  // assuming API returns an array of items
  const entrys = res.data;

  // clear table body
  tbody.innerHTML = "";

  // Check if data on page. If no data then move back a page till on page 1
  if (!entrys) {
    if (page != 1) {
      page--;
      reloadData();
    }
    return;
  }

  const headsUpContainer = document.getElementById("noResultsContainer");
  if (!entrys) {
    headsUpContainer.style.display = "flex";
    return 0;
  }

  headsUpContainer.style.display = "none";

  // Add rows
  entrys.forEach((entry) => {
    const row = document.createElement("tr");

    var html = `
                <td>${entry.upc || ""}</td>
                <td>${entry.qty || ""}</td>
                <td>${entry.details || entry.rawtext || ""}</td>
                <td>${formatPrice(entry.true_cost, false) || ""}</td>
                `;

    row.innerHTML = html;

    tbody.appendChild(row);
  });

  return entrys[0]?.count ?? 0;
}

function invoiceEdit() {
  const display = document.getElementById("invoiceName");
  const input = document.getElementById("invoiceNameInput");
  // Load Invoice Name from Database
  axios
    .get(`/api/invoices/${InvoiceID}`)
    .then((res) => {
      input.value = res.data.invoice_name;
    })
    .catch((err) => {
      return;
    });

  display.style.display = "none";
  input.style.display = "block";
  input.focus();
  input.select();

  input.onblur = () => invoiceSave();

  input.addEventListener("keypress", function (e) {
    if (e.key === "Enter") invoiceSave();
  });
}

function invoiceSave() {
  const display = document.getElementById("invoiceName");
  const input = document.getElementById("invoiceNameInput");

  // quit if name is empty
  if (input.value == "") {
    display.style.display = "block";
    input.style.display = "none";
    return;
  }

  // Save to Database
  axios
    .put(`/api/invoices/${InvoiceID}`, { invoice_name: input.value })
    .then(() => {
      display.textContent = input.value;
    })
    .catch((err) => {
      console.log("error updating invoice name in database");
    });

  // Exit title editor
  display.style.display = "block";
  input.style.display = "none";
}

document.addEventListener("DOMContentLoaded", function () {
  loadInvoiceName();
  reloadData();

  // Watch page size dropdown
  const itemsPerPageSelector = document.getElementById("itemsPerPageSelector");
  itemsPerPageSelector.addEventListener("change", (e) => {
    pageSize = e.target.value;
    reloadData();
  });
});
