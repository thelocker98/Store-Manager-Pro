let selectedCatalogId = null;
let allCatalog = [];

function init() {
  axios.get("/api/catalog").then((res) => {
    allCatalog = res.data;
    renderCatalog(allCatalog);
  });
  loadVendors();
}

function renderCatalog(list) {
  const tbody = document.querySelector("#catalogTable tbody");
  tbody.innerHTML = "";
  list.forEach((catalogEntry) => {
    tbody.innerHTML += `
                    <tr onclick="selectCatalog(${catalogEntry.catalog_id}, '${catalogEntry.name}', '${catalogEntry.brand}')">
                        <td>${catalogEntry.upc}</td>
                        <td>${catalogEntry.invoice_number}</td>
                        <td>${catalogEntry.name}</td>
                        <td>${catalogEntry.brand}</td>
                    </tr>
                `;
  });
}

function searchCatalog() {
  const q = document.getElementById("catalogSearch").value.toLowerCase();
  renderCatalog(
    allCatalog.filter(
      (c) =>
        c.name.toLowerCase().includes(q) ||
        c.brand.toLowerCase().includes(q) ||
        c.upc.includes(q),
    ),
  );
}

function selectCatalog(id, name, brand) {
  selectedCatalogId = id;
  document.getElementById("selectedCatalogTitle").innerText =
    `${name} (${brand})`;
  document.getElementById("itemDetails").style.display = "block";
}

function loadVendors() {
  axios.get("/api/vendors").then((res) => {
    const v = document.getElementById("vendorSelect");
    v.innerHTML = "";
    res.data.forEach((vendor) => {
      v.innerHTML += `<option value="${vendor.vendor_id}">${vendor.vendor_name}</option>`;
    });
  });
}

function openVendorIframe() {
  document.getElementById("vendorIframeModal").style.display = "block";
}

function closeVendorIframe() {
  document.getElementById("vendorIframeModal").style.display = "none";
  loadVendors();
}

function submitItem() {
  console.log(selectedCatalogId);
  if (!selectedCatalogId) return alert("Select a catalog item");
  const payload = {
    catalog_id: selectedCatalogId,
    vendor_id: parseInt(document.getElementById("vendorSelect").value),
    price: parseFloat(document.getElementById("itemPrice").value),
    count: parseInt(document.getElementById("itemCount").value),
    weighed: document.getElementById("itemWeighed").checked,
    description: document.getElementById("itemDesc").value,
  };
  axios.post("/api/items", payload).then(() => {
    closeItem();
    // Optionally inform parent page to refresh items
    if (window.parent && window.parent.loadItems) window.parent.loadItems();
  });
}

function closeItem() {
  window.parent &&
    window.parent.closeAddItemIframe &&
    window.parent.closeAddItemIframe();
}

init();
