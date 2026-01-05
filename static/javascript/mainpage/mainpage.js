// variables
var showDeletedItems = false;

// Load all items
async function loadItems() {
  const res = await axios.get(
    `/api/items?page=${page}&pagesize=${pageSize}&showdeleted=${showDeletedItems}`,
  );
  const items = res.data;

  if (!items) {
    if (page != 0) {
      page--;
      reloadData();
    }
    return;
  }

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
  return res.data[0]?.number_of_entrys ?? 0;
}

// Delete Item
function deleteItem(id) {
  axios.delete(`/api/items/${id}`).then(() => loadItems());
  closeInfo();
  closeItemPopup();
}
// Restore Item
function restoreItem(id) {
  axios.get(`/api/items/restore/${id}`).then(() => loadItems());
  closeInfo();
  closeItemPopup();
}
// Delete Permanently Item
function deleteItemPermanent(id) {
  if (confirm("Are you sure you want to permanently delete this item?")) {
    axios.delete(`/api/items/permanent/${id}`).then(() => loadItems());
    closeInfo();
    closeItemPopup();
  }
}

// Open the Vendor Iframe popup
function openVendorIframe() {
  document.getElementById("vendorIframeModal").style.display = "block";
  document.getElementById("iframeOverlay").style.display = "block";
}

// Close the Vendor Iframe popup
function closeVendorIframe() {
  document.getElementById("vendorIframeModal").style.display = "none";
  document.getElementById("iframeOverlay").style.display = "none";

  loadVendors();
}

// Populate Vendors selector
function loadVendors() {
  return axios.get("/api/vendors").then((res) => {
    const select = document.getElementById("popup_items_vendorselector");
    select.innerHTML = "";

    const opt = document.createElement("option");
    opt.value = -1;
    opt.text = "Select Vendor";
    select.appendChild(opt);

    res.data.forEach((v) => {
      const opt = document.createElement("option");
      opt.value = v.vendor_id;
      opt.text = v.vendor_name;
      select.appendChild(opt);
    });
  });
}

document.addEventListener("DOMContentLoaded", function () {
  const itemsPerPageSelector = document.getElementById("itemsPerPageSelector");
  const showDeletedItemsDOM = document.getElementById("showDeletedCheckbox");

  itemsPerPageSelector.addEventListener("change", (e) => {
    pageSize = e.target.value;
    // reload data and refresh page buttons
    reloadData(page);
  });

  showDeletedItemsDOM.addEventListener("change", (e) => {
    showDeletedItems = showDeletedItemsDOM.checked;
    // reload data and refresh page buttons
    reloadData(page);
  });

  // Price decimal conversion
  const priceInput = document.getElementById("popup_items_price");

  priceInput.addEventListener("input", () => {
    // Remove anything that's not a digit
    let raw = priceInput.value.replace(/\D/g, "");

    if (raw === "") {
      priceInput.value = "";
      return;
    }

    // Convert to cents → dollars
    const value = (parseInt(raw, 10) / 100).toFixed(2);
    priceInput.value = value;
  });

  // Load Inital Data
  reloadData();
});
