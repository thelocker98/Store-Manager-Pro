// Load all items
function loadItems() {
  axios.get("/api/items").then((res) => {
    const tbody = document.querySelector("#itemsTable tbody");
    tbody.innerHTML = "";
    res.data.forEach((item) => {
      tbody.innerHTML += `
                        <tr onclick="openInfo(${item.item_id})" style="cursor:pointer">
                            <td>${item.upc}</td>
                            <td>${item.brand}</td>
                            <td>${item.vendor_name}</td>
                            <td>${item.name}</td>
                            <td>${item.description}</td>
                            <td>${formatPrice(item.price)}</td>
                            <td>${formatDate(item.arrived_at)}</td>
                        </tr>
                        `;
    });
  });
}

// Add Item
function addItem(e) {
  e.preventDefault();
  const item = {
    upc: document.getElementById("upc").value,
    brand: document.getElementById("brand").value,
    name: document.getElementById("name").value,
    description: document.getElementById("description").value,
    price: parseFloat(document.getElementById("price").value),
  };
  axios.post("/api/items", item).then(() => {
    loadItems();
    document.getElementById("addForm").reset();
  });
}

// Delete Item
function deleteItem(id) {
  if (confirm("Are you sure you want to delete this item?")) {
    axios.delete(`/api/items/${id}`).then(() => loadItems());
  }
}

// Search (simple client-side filter)
function searchItems() {
  const query = document.getElementById("searchInput").value.toLowerCase();
  axios.get("/api/items").then((res) => {
    const filtered = res.data.filter(
      (item) =>
        item.name.toLowerCase().includes(query) ||
        item.brand.toLowerCase().includes(query) ||
        item.upc.includes(query),
    );
    const tbody = document.querySelector("#itemsTable tbody");
    tbody.innerHTML = "";
    filtered.forEach((item) => {
      tbody.innerHTML += `
                        <tr>
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

function openInfo(id) {
  axios.get("/api/items").then((res) => {
    const item = res.data.find((i) => i.item_id === id);
    if (!item) return;

    console.log(item);
    let html = `<h2>${item.name} (${item.brand})</h2>`;

    if (item.upc != "") {
      html += `<strong>UPC:</strong> ${item.upc}<br />`;
    }
    if (item.invoice_number != "") {
      html += `<strong>Invoice Number:</strong> ${item.invoice_number}<br />`;
    }

    html += `<strong>Description:</strong><br/> ${item.description}
              <br />
              <br />
              <strong>Price:</strong> ${formatPrice(item.price, item.weighed)}
    `;

    console.log(item.count);
    if (item.count != null) {
      html += `&nbsp;&nbsp;&nbsp;&nbsp<strong>Count:</strong> ${item.count}<br />`;
    }

    html += `<br/><strong>Bought:</strong> ${formatDate(item.arrived_at)}`;

    if (item.sold_out_at != null) {
      html += `&nbsp;&nbsp<strong>Sold:</strong> ${formatDate(item.sold_out_at)}`;
    }

    document.getElementById("infoContent").innerHTML = html;

    document.getElementById("infoEditBtn").onclick = () =>
      editItem(item.item_id);
    document.getElementById("infoDeleteBtn").onclick = () =>
      deleteItem(item.item_id);

    document.getElementById("infoPopup").style.display = "block";
    document.getElementById("infooverlay").style.display = "block";
  });
}

function closeInfo() {
  document.getElementById("infoPopup").style.display = "none";
  document.getElementById("infooverlay").style.display = "none";
}

// Open the edit popup and fill fields
function editItem(id) {
  axios.get("/api/items").then((res) => {
    const item = res.data.find((i) => i.item_id === id);
    if (!item) return;

    document.getElementById("editId").value = item.item_id;
    document.getElementById("editUpc").value = item.upc;
    document.getElementById("editBrand").value = item.brand;
    document.getElementById("editName").value = item.name;
    document.getElementById("editDescription").value = item.description;
    document.getElementById("editPrice").value = item.price;

    document.getElementById("editPopup").style.display = "block";
    document.getElementById("editoverlay").style.display = "block";
  });
}
// Submit edit form
function submitEdit(e) {
  e.preventDefault();
  const id = document.getElementById("editId").value;
  const item = {
    upc: document.getElementById("editUpc").value,
    brand: document.getElementById("editBrand").value,
    name: document.getElementById("editName").value,
    description: document.getElementById("editDescription").value,
    price: parseFloat(document.getElementById("editPrice").value),
  };

  axios.put(`/api/items/${id}`, item).then(() => {
    closeEdit();
    loadItems();
  });
}

// Close the popup
function closeEdit() {
  document.getElementById("editPopup").style.display = "none";
  document.getElementById("editoverlay").style.display = "none";
}

function openAddItemIframe() {
  const modal = document.getElementById("addItemIframeModal");
  const iframe = modal.querySelector("iframe");

  // Reset the iframe src to force reload
  iframe.src = "/items";

  modal.style.display = "block";
}

function closeAddItemIframe() {
  document.getElementById("addItemIframeModal").style.display = "none";
}
