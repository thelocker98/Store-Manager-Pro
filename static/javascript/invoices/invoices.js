async function renderInvoices() {
  if (GlobalDepartment == 0) {
    document.getElementById("addInvoiceButton").disabled = true;
  } else {
    document.getElementById("addInvoiceButton").disabled = false;
  }

  const res = await axios.get(`/api/invoices/d/${GlobalDepartment}`);
  const invoices = res.data;

  const container = document.getElementById("invoiceContainers");
  container.innerHTML = "";

  if (!Array.isArray(invoices)) return;

  invoices.forEach((invoice) => {
    const card = document.createElement("div");
    card.className = "invoice-card";
    card.onclick = () => viewInvoice(invoice.invoice_id);
    card.innerHTML = `
                     <div class="invoice-card-header">
                         <h3 class="invoice-card-title">${invoice.invoice_name}</h3>
                         <div>
                          <span class="invoice-card-delete material-symbol">&#xE872;</span>
                         </div>
                     </div>
                     <div class="invoice-card-footer">
                         <span>Created: ${formatDate(invoice.created_at)}</span>
                         <span>Item Count: ${invoice.total_count}</span>
                     </div>
                 `;
    // Add delete handler with stopPropagation
    const deleteBtn = card.querySelector(".invoice-card-delete");
    deleteBtn.onclick = (e) => {
      e.stopPropagation(); // Prevents the card click from firing
      deleteInvoice(invoice.invoice_id);
    };

    // Add invoice card to the grid
    container.appendChild(card);
  });
}

// Function to delete invoice
function deleteInvoice(invoice_id) {
  if (confirm("Are you sure you want to permanently delete this invoice?")) {
    axios.delete(`/api/invoices/${invoice_id}`).then(() => renderInvoices());
  }
}

// Function to go to the invoice
function viewInvoice(invoiceId) {
  // Navigate to the invoice detail page
  window.location.href = `/invoices/${invoiceId}`;
}

// Load invoices when page loads
document.addEventListener("DOMContentLoaded", () => {
  renderInvoices();
});

// Invoice Popup
function openInvoice(id) {
  document.getElementById("invoiceError").style.display = "none";
  document.getElementById("invoicePopup").style.display = "block";
  document.getElementById("invoiceOverlay").style.display = "block";
}

function closeInvoice() {
  document.getElementById("invoicePopup").style.display = "none";
  document.getElementById("invoiceOverlay").style.display = "none";
}

function createNewInvoice() {
  const invoice = {
    invoice_name: document.getElementById("popup_invoice_title").value,
    department_id: GlobalDepartment,
  };

  axios
    .post("/api/invoices", invoice)
    .then(() => {
      closeInvoice();
      document.getElementById("popup_invoice_title").value = "";
      renderInvoices();
    })
    .catch((err) => {
      document.getElementById("invoiceError").style.display = "block";
      document.getElementById("invoiceError").textContent = "Error Creating invoice";
    });
}
