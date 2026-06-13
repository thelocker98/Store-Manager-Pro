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
    row.id = `tr-${entry.invoice_entry_id}`;

    var html = `
                <td>${entry.upc || ""}</td>
                <td>${entry.qty || ""}</td>
                <td>${entry.details || entry.rawtext || ""}</td>
                <td>${formatPrice(entry.true_cost, false) || ""}</td>
                `;

    row.innerHTML = html;

    tbody.appendChild(row);

    // Open Info on click
    row.onclick = () => openInvoiceEntryEdit(entry.invoice_entry_id);
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

function openInvoiceEntryEdit(invoice_id) {
  const row = document.getElementById(`tr-${invoice_id}`);
  const container = document.createElement("tr");
  container.id = `tr-div-${invoice_id}`;

  if (document.getElementById(container.id)) {
    saveInvoiceEntryEdit(invoice_id);
    return;
  }

  axios
    .get(`/api/invoiceentry/${invoice_id}`)
    .then((res) => {
      var entry = res.data;

      const inputStyle = `
        font-size: 20px;
        padding: 5px;
        text-align: center;
        border: 1px solid #ccc;
        border-radius: 6px;
        box-sizing: border-box;
        margin: 10px;
      `;

      const quantity = Math.round(entry.total_cost / entry.item_cost);

      container.innerHTML = `
      <td colspan="4" style="padding: 20px 0px;">
        <div style="width:100%; display:flex; flex-direction:column; align-items:center; border-radius: 10px; padding: 10px 0px; background-color: lightgray; ">

          <h2 style="font-size:50px; text-align: center; margin: 5px;">${entry.details}</h2>
          <h2 style="font-size:80px; text-align: center; margin: 0 50px;">Price: ${formatPrice(entry.true_cost, false)}&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;QTY: ${entry.qty}</h2>

          <div style="margin: 30px 10px; padding: 10px; width: 90%; border-radius: 10px; background-color: #FFA07A;">
            <p style="margin: 5px 20px;"><strong>Raw Invoice Line:</strong> ${entry.rawtext}</p>

            <div style="display:flex; justify-content:space-between;">
              <p style="margin: 5px 20px;"><strong>Item Cost: </strong>${formatPrice(entry.item_cost, false)}</p>
              <p style="margin: 5px 20px;"><strong>Quantity Received: </strong>${quantity}</p>
              <p style="margin: 5px 20px;"><strong>Total Cost: </strong>${formatPrice(entry.total_cost, false)}</p>
              <p style="margin: 5px 20px;"><strong>Actually Paid: </strong>${formatPrice(entry.discount_cost, false)}</p>
            </div>
          </div>

          <div style="display:flex; justify-content: space-between; width:90%;">
              <div>
                <div style="flex:1;">
                  <label style="font-size:25px;">Entry Details:</label>
                  <input type="text" id="editInvoiceDetailsInput" value="${entry.details}" style="${inputStyle} width: 700px;" />
                </div>
                <div style="flex:1;">
                  <label style="font-size:25px;">QTY:</label>
                  <input type="number" id="editInvoiceQTYInput" value="${entry.qty}" style="${inputStyle} width: 100px" />
                </div>
              </div>
              <div>
                <div style="flex:1;">
                  <label style="font-size:25px;">UPC:</label>
                  <input type="text" id="editInvoiceUPCInput" value="${entry.upc}" style="${inputStyle} width: 170px;" />
                </div>
                <div style="flex:1;">
                  <label style="font-size:25px;">Item Cost:</label>
                  <input type="number" id="editInvoiceCostInput" value="${entry.true_cost}" style="${inputStyle} width: 115px;" />
                </div>
              </div>
            </div>

            <div style="display:flex; justify-content: space-between; width:90%;">
              <button style="padding:10px; width:100px" onclick="cancelInvoiceEntryEdit(${invoice_id})">Close</button>
              <button style="padding:10px; width:100px" onclick="saveInvoiceEntryEdit(${invoice_id})">Save</button>
            </div>
          </div>
        </div>
      </td>
      `;

      row.style.display = "none";

      row.after(container);
    })
    .catch((err) => {
      return;
    });
}

function saveInvoiceEntryEdit(invoice_id, showConfirm = false) {
  exists = document.getElementById(`tr-div-${invoice_id}`);
  row = document.getElementById(`tr-${invoice_id}`);

  if (exists) {
    // Get

    var details_input = document.getElementById("editInvoiceDetailsInput").value;
    var qty_input = document.getElementById("editInvoiceQTYInput").value;
    var cost_input = document.getElementById("editInvoiceCostInput").value;
    var upc_input = document.getElementById("editInvoiceUPCInput").value;

    axios
      .get(`/api/invoiceentry/${invoice_id}`)
      .then((res) => {
        var entry = res.data;

        if (
          details_input != entry.details ||
          qty_input != entry.qty ||
          cost_input != entry.true_cost ||
          upc_input != entry.upc
        ) {
          // Check if should show confirm other wise just save
          if (showConfirm) {
            const answer = window.confirm("You have unsaved changes. Do you want to save them?");
            if (!answer) {
              exists.remove();
              row.style.display = "table-row";
              return;
            }
          }

          // Update entry and then save
          entry.details = details_input;
          entry.qty = parseInt(qty_input, 10);
          entry.true_cost = parseFloat(cost_input, 10);
          entry.upc = upc_input;
          console.log("updating");

          axios.put("/api/invoiceentrys/" + invoice_id, entry).then(() => {
            // refresh page
            document.getElementById("searchInput").value = ""; // Clear Search
            reloadData();
          });
        }

        // Reload Referances
        exists = document.getElementById(`tr-div-${invoice_id}`);
        row = document.getElementById(`tr-${invoice_id}`);
        // Preform action
        exists.remove();
        row.style.display = "table-row";

        return;
      })
      .catch((err) => {
        return;
      });
  }
}

function cancelInvoiceEntryEdit(invoice_id) {
  exists = document.getElementById(`tr-div-${invoice_id}`);
  row = document.getElementById(`tr-${invoice_id}`);

  if (exists) {
    exists.remove();
    row.style.display = "table-row";
    return;
  }
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

  document.addEventListener("click", function (event) {
    // Check if there's an open edit div
    const openEditDiv = document.querySelector('[id^="tr-div-"]');

    if (openEditDiv) {
      // Get the invoice_id from the div id
      const invoice_id = openEditDiv.id.replace("tr-div-", "");
      const originalRow = document.getElementById(`tr-${invoice_id}`);

      // Check if click is outside the edit div and not on the original row
      const isClickInsideEdit = openEditDiv.contains(event.target);
      const isClickOnOriginalRow = originalRow && originalRow.contains(event.target);

      // If click is outside edit div and not on the original row, close it
      if (!isClickInsideEdit && !isClickOnOriginalRow) {
        console.log("Saved");
        saveInvoiceEntryEdit(invoice_id, true);
      }
    }
  });
});
