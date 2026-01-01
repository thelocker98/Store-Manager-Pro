// Info Popup
function openInfo(id) {
  axios.get(`/api/items/${id}`).then((res) => {
    const item = res.data;
    if (!item) return;

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

    if (item.count != null) {
      html += `&nbsp;&nbsp;&nbsp;&nbsp<strong>Count:</strong> ${item.count}<br />`;
    }

    html += `<br/><strong>Bought:</strong> ${formatDate(item.arrived_at)}`;

    if (item.sold_out_at != null) {
      html += `&nbsp;&nbsp<strong>Sold:</strong> ${formatDate(item.sold_out_at)}`;
    }

    document.getElementById("infoContent").innerHTML = html;

    // Check if the Item has already been marked as deleted
    if (item.deleted) {
      // Edit to Restore button
      document.getElementById("infoEditBtn").textContent = "Restore";
      document.getElementById("infoEditBtn").onclick = () =>
        restoreItem(item.item_id);
      // Delete
      document.getElementById("infoDeleteBtn").textContent =
        "Delete Permanently";
      document.getElementById("infoDeleteBtn").onclick = () =>
        deleteItemPermanent(item.item_id);
    } else {
      // Restore to Edit button
      document.getElementById("infoEditBtn").textContent = "Edit";
      document.getElementById("infoEditBtn").onclick = () =>
        openEdititemsEntry(item.item_id);
      // Delete
      document.getElementById("infoDeleteBtn").textContent = "Delete";
      document.getElementById("infoDeleteBtn").onclick = () =>
        deleteItem(item.item_id);
    }

    document.getElementById("infoPopup").style.display = "block";
    document.getElementById("infooverlay").style.display = "block";
  });
}

function closeInfo() {
  document.getElementById("editError").style.display = "none";
  document.getElementById("editError").textContent = "";
  document.getElementById("infoPopup").style.display = "none";
  document.getElementById("infooverlay").style.display = "none";
}
