async function searchItems() {
  const query = document.getElementById("searchInput").value;
  const tbody = document.querySelector("#invoiceTable tbody");
  const noResultsContainer = document.getElementById("noResultsContainer");

  try {
    const res = await axios.get(
      `/api/invoice/search/${encodeURIComponent(query)}?page=${page}&pagesize=${pageSize}`,
    );

    // assuming API returns an array of invoice entrys
    const entrys = res.data;

    // check if the invoice entrys array is empty
    if (!entrys) {
      if (page != 1) {
        page--;
        reloadData();
      }
      // No results found -> show Add Item button
      tbody.innerHTML = "";
      noResultsContainer.style.display = "block";
      closeInvoiceDetailSection();

      return 1;
    } else {
      noResultsContainer.style.display = "none";
    }

    // check if their is only one result then show it large
    if (entrys.length == 1) {
      openInvoiceDetailSection(entrys[0]);
    } else {
      closeInvoiceDetailSection();
    }

    // clear table body
    tbody.innerHTML = ""; // clear previous results

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
      // Open Info on click
      row.onclick = () => openInfo(entry.invoice_id);

      tbody.appendChild(row);
    });
    return entrys[0].number_of_entrys;
  } catch (error) {
    console.error("Search error:", error);
  }
}

function resetSearch() {
  noResultsContainer.style.display = "none";
  const input = document.getElementById("searchInput");
  input.value = "";

  reloadData();
}

function openInvoiceDetailSection(entry) {
  document.getElementById("invoiceDetailSectionTitle").textContent = entry.name;
  document.getElementById("invoiceDetailSectionPrice").textContent =
    "Price:" + formatPrice(entry.price, false);
  document.getElementById("invoiceDetailSection").style.display = "block";
}

function closeInvoiceDetailSection() {
  document.getElementById("invocieDetailSection").style.display = "none";
}

document.addEventListener("DOMContentLoaded", function () {
  const searchInput = document.getElementById("searchInput");

  if (searchInput) {
    searchInput.addEventListener("keypress", function (e) {
      if (e.key === "Enter") {
        reloadData();
      }
    });
  }
});
