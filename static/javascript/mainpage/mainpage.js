// Load all items
function loadItems() {
  axios.get("/api/items").then((res) => {
    const tbody = document.querySelector("#itemsTable tbody");

    tbody.innerHTML = "";

    res.data.forEach((item) => {
      var class_val = "";

      if (item.deleted) {
        class_val = "deleted";
      } else {
        class_val = "avalible";
      }
      tbody.innerHTML += `<tr onclick="openInfo(${item.item_id})" style="cursor:pointer" class="${class_val}">
                            <td>${item.upc}</td>
                            <td>${item.brand}</td>
                            <td>${item.vendor_name}</td>
                            <td>${item.name}</td>
                            <td>${item.description}</td>
                            <td>${formatPrice(item.price, item.weighed)}</td>
                            <td>${formatDate(item.arrived_at)}</td>
                          </tr>
                          `;
    });
  });
}

// Initial load
loadItems();

// Delete Item
function deleteItem(id) {
  axios.delete(`/api/items/${id}`).then(() => loadItems());
  closeInfo();
  closeItemsPopup();
}
// Restore Item
function restoreItem(id) {
  axios.get(`/api/items/restore/${id}`).then(() => loadItems());
  closeInfo();
  closeItemsPopup();
}
// Delete Permanently Item
function deleteItemPermanent(id) {
  if (confirm("Are you sure you want to permanently delete this item?")) {
    axios.delete(`/api/items/permanent/${id}`).then(() => loadItems());
    closeInfo();
    closeItemsPopup();
  }
}

// Open the Vendor Iframe popup
function openVendorIframe() {
  document.getElementById("vendorIframeModal").style.display = "block";
}

// Close the Vendor Iframe popup
function closeVendorIframe() {
  document.getElementById("vendorIframeModal").style.display = "none";
  loadVendors();
}
