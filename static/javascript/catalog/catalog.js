// Load Catalog
function loadCatalog() {
  axios.get("/api/catalog").then((res) => {
    const tbody = document.querySelector("#catalogTable tbody");
    tbody.innerHTML = "";
    res.data.forEach((catalog) => {
      tbody.innerHTML += `
                        <tr onclick="openInfo(${catalog.catalog_id})" style="cursor:pointer">
                                    <td>${catalog.upc}</td>
                                    <td>${catalog.invoice_number}</td>
                                    <td>${catalog.name}</td>
                                    <td>${catalog.brand}</td>
                                </tr>
                            `;
    });
  });
}

// Add Catalog Entry
function addCatalogEntry(e) {
  e.preventDefault();
  const catalogEntry = {
    upc: document.getElementById("catalog_upc").value,
    invoice_number: document.getElementById("catalog_invoicenumber").value,
    brand: document.getElementById("catalog_brand").value,
    name: document.getElementById("catalog_name").value,
  };
  axios.post("/api/catalog", catalogEntry).then(() => {
    loadCatalog();
    document.getElementById("addCatalogEntryForm").reset();
  });
}

// Delete Catalog Entry
function deleteCatalogEntry(id) {
  if (confirm("Are you sure you want to delete this Catalog Entry?")) {
    axios
      .delete(`/api/catalog/${id}`)
      .then(() => {
        loadCatalog();
        closeInfo();
      })
      .catch((err) => {
        console.error("Failed to delete catalog entry:", err);
      });
  }
}
// Initial load
loadCatalog();

// Open the edit popup and fill fields
function editCatalogEntry(id) {
  axios.get("/api/catalog").then((res) => {
    const catalogEntry = res.data.find((i) => i.catalog_id === id);
    if (!catalogEntry) return;

    document.getElementById("editCatalogId").value = catalogEntry.catalog_id;
    document.getElementById("editCatalogUPC").value = catalogEntry.upc;
    document.getElementById("editCatalogInvoiceNumber").value =
      catalogEntry.invoice_number;
    document.getElementById("editCatalogName").value = catalogEntry.name;
    document.getElementById("editCatalogBrand").value = catalogEntry.brand;

    document.getElementById("editCatalogEntryPopup").style.display = "block";
    document.getElementById("overlayEditor").style.display = "block";
  });
}

let currentInfoItem = null;

function openInfo(id) {
  axios.get("/api/catalog").then((res) => {
    const catalogEntry = res.data.find((i) => i.catalog_id === id);
    if (!catalogEntry) return;
    console.log(catalogEntry);

    currentInfoItem = catalogEntry;

    document.getElementById("infoContent").innerHTML = `
                        <strong>UPC:</strong> ${catalogEntry.upc}<br />
                        <strong>Invoice Number:</strong> ${catalogEntry.invoice_number}<br /><br />

                        <strong>Brand:</strong> ${catalogEntry.brand}<br />
                        <strong>Name:</strong> ${catalogEntry.name}<br />
                    `;

    document.getElementById("infoEditBtn").onclick = () =>
      editCatalogEntry(catalogEntry.catalog_id);
    document.getElementById("infoDeleteBtn").onclick = () =>
      deleteCatalogEntry(catalogEntry.catalog_id);

    document.getElementById("infoPopup").style.display = "block";
    document.getElementById("infooverlay").style.display = "block";
  });
}

function closeInfo() {
  document.getElementById("infoPopup").style.display = "none";
  document.getElementById("infooverlay").style.display = "none";
}

// Close the popup
function closeCatalogEntryEdit() {
  document.getElementById("editCatalogEntryPopup").style.display = "none";
  document.getElementById("overlayEditor").style.display = "none";
}

// Submit edit form
function submitCatalogEntryEdit(e) {
  e.preventDefault();
  const id = document.getElementById("editCatalogId").value;
  const catalogEntry = {
    upc: document.getElementById("editCatalogUPC").value,
    invoice_number: document.getElementById("editCatalogInvoiceNumber").value,
    brand: document.getElementById("editCatalogBrand").value,
    name: document.getElementById("editCatalogName").value,
  };

  axios.put(`/api/catalog/${id}`, catalogEntry).then(() => {
    closeCatalogEntryEdit();
    closeInfo();
    loadCatalog();
  });
}

// Open the Add Catalog Popup
function openAddCatalogPopup() {
  document.getElementById("addCatalogPopup").style.display = "block";
  document.getElementById("addCatalogOverlay").style.display = "block";
}

// Close the Add Catalog Popup
function closeAddCatalogPopup() {
  document.getElementById("addCatalogPopup").style.display = "none";
  document.getElementById("addCatalogOverlay").style.display = "none";
}

// Add Catalog Entry from Popup
function addCatalogEntryPopup(e) {
  e.preventDefault();
  const catalogEntry = {
    upc: document.getElementById("popup_catalog_upc").value,
    invoice_number: document.getElementById("popup_catalog_invoicenumber")
      .value,
    name: document.getElementById("popup_catalog_name").value,
    brand: document.getElementById("popup_catalog_brand").value,
  };
  axios.post("/api/catalog", catalogEntry).then(() => {
    loadCatalog();
    document.getElementById("addCatalogEntryFormPopup").reset();
    closeAddCatalogPopup();
  });
}

// Search catalog client-side
function searchCatalogEntries() {
  const query = document.getElementById("searchCatalog").value.toLowerCase();
  axios.get("/api/catalog").then((res) => {
    const filtered = res.data.filter(
      (item) =>
        item.name.toLowerCase().includes(query) ||
        item.brand.toLowerCase().includes(query) ||
        item.upc.includes(query) ||
        item.invoice_number.includes(query),
    );
    const tbody = document.querySelector("#catalogTable tbody");
    tbody.innerHTML = "";
    filtered.forEach((catalog) => {
      tbody.innerHTML += `
                        <tr onclick="openInfo(${catalog.catalog_id})" style="cursor:pointer">
                            <td>${catalog.upc}</td>
                            <td>${catalog.invoice_number}</td>
                            <td>${catalog.name}</td>
                            <td>${catalog.brand}</td>
                        </tr>
                    `;
    });
  });
}

function resetSearch() {
  document.getElementById("searchCatalog").value = "";
  loadCatalog();
}
