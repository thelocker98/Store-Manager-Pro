// Variables
let currentInfoItem = null;

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
// Initial load
loadCatalog();

// Open the Add Catalog Popup
function openAddCatalogPopup() {
  document.getElementById("catalogPopupHeader").textContent =
    "Add Catalog Entry";

  document.getElementById("catalogEntryFormPopup").onsubmit = () =>
    submitAddCatalogEntryPopup(event);
  document.getElementById("catalogPopup").style.display = "block";
  document.getElementById("catalogOverlay").style.display = "block";
}

// Add Catalog Entry from Popup
function submitAddCatalogEntryPopup(e) {
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
    document.getElementById("catalogEntryFormPopup").reset();
    closeCatalogPopup();
  });
}

// Open the edit popup and fill fields
function openEditCatalogEntry(id) {
  axios.get("/api/catalog").then((res) => {
    const catalogEntry = res.data.find((i) => i.catalog_id === id);
    if (!catalogEntry) return;

    document.getElementById("catalogPopupHeader").textContent =
      "Edit Catalog Entry";

    document.getElementById("catalogEntrySubmitButton").textContent = "Save";
    document.getElementById("catalogEntryFormPopup").onsubmit = () =>
      submitEditCatalogEntry(event);

    document.getElementById("popup_catalog_id").value = catalogEntry.catalog_id;
    document.getElementById("popup_catalog_upc").value = catalogEntry.upc;
    document.getElementById("popup_catalog_invoicenumber").value =
      catalogEntry.invoice_number;
    document.getElementById("popup_catalog_name").value = catalogEntry.name;
    document.getElementById("popup_catalog_brand").value = catalogEntry.brand;

    document.getElementById("catalogPopup").style.display = "block";
    document.getElementById("catalogOverlay").style.display = "block";
  });
}

// Submit edit form
function submitEditCatalogEntry(e) {
  e.preventDefault();
  const id = document.getElementById("popup_catalog_id").value;
  const catalogEntry = {
    upc: document.getElementById("popup_catalog_upc").value,
    invoice_number: document.getElementById("popup_catalog_invoicenumber")
      .value,
    brand: document.getElementById("popup_catalog_brand").value,
    name: document.getElementById("popup_catalog_name").value,
  };

  axios.put(`/api/catalog/${id}`, catalogEntry).then(() => {
    closeCatalogPopup();
    closeInfo();
    loadCatalog();
  });
}

// Close the Add Catalog Popup
function closeCatalogPopup() {
  document.getElementById("catalogPopup").style.display = "none";
  document.getElementById("catalogOverlay").style.display = "none";
}

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
      openEditCatalogEntry(catalogEntry.catalog_id);
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
