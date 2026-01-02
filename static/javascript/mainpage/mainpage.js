// variables
var pageSize = 2;
var page = 1;
var numberOfPages;
var showDeletedItems = false;

// Load all items
function loadItems() {
  axios
    .get(
      `/api/items?page=${page}&pagesize=${pageSize}&showdeleted=${showDeletedItems}`,
    )
    .then((res) => {
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
  changePage(0);
  const pagesPerScreenSelector = document.getElementById(
    "pagePerScreenSelector",
  );
  const showDeletedItemsDOM = document.getElementById("showDeletedItems");

  pagesPerScreenSelector.addEventListener("change", (e) => {
    pageSize = e.target.value;
    changePage(0);
    loadItems();
  });

  showDeletedItemsDOM.addEventListener("change", (e) => {
    showDeletedItems = showDeletedItemsDOM.checked;
    changePage(0);
    loadItems();
  });
});

async function changePage(changePageBy) {
  page += changePageBy;

  // Clamp page to minimum of 1
  if (page < 1) page = 1;

  // Get total count
  const res = await axios.get(
    `/api/items/count?countdeleted=${showDeletedItems}`,
  );

  let data = res.data;
  if (typeof data === "string") {
    data = JSON.parse(data);
  }

  numberOfPages = Math.ceil(data.count / pageSize);

  // Clamp page to max
  if (page > numberOfPages) page = numberOfPages;

  // Update buttons
  document.getElementById("prevPageBtn").disabled = page <= 1;
  document.getElementById("nextPageBtn").disabled = page >= numberOfPages;

  // Update indicator
  document.getElementById("pageIndicator").textContent =
    `Page ${page} of ${numberOfPages}`;

  loadItems();
}
