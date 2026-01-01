// Open the Add items Popup
function openAdditemsPopup() {
  // Load vendor selector
  loadVendors();
  // Set Title
  document.getElementById("itemsPopupHeader").textContent = "Add items Entry";

  // Set Buttons
  document.getElementById("itemsEntrySubmitButton").textContent = "Add";
  document.getElementById("itemsEntryFormPopup").onsubmit = () =>
    submitAdditemsEntryPopup(event);

  // Set Content
  document.getElementById("popup_items_id").value = 0;
  document.getElementById("popup_items_upc").value = "";
  document.getElementById("popup_items_invoicenumber").value = "";
  document.getElementById("popup_items_vendorselector").value = -1;
  document.getElementById("popup_items_name").value = "";
  document.getElementById("popup_items_brand").value = "";
  document.getElementById("popup_items_description").value = "";
  document.getElementById("popup_items_price").value = 0;
  document.getElementById("popup_items_weighed").value = false;
  document.getElementById("popup_items_count").value = 0;

  // Set layers
  document.getElementById("popup_div_items_arrived").style.display = "none";
  document.getElementById("popup_div_items_soldout").style.display = "none";
  document.getElementById("itemsPopup").style.display = "block";
  document.getElementById("itemsOverlay").style.display = "block";
}

// Open the edit popup and fill fields
async function openEdititemsEntry(id) {
  // Load vendor selector
  await loadVendors();

  // Check that entry exists in database
  axios.get("/api/items").then((res) => {
    const itemsEntry = res.data.find((i) => i.item_id === id);
    if (!itemsEntry) return;

    // Set Title
    document.getElementById("itemsPopupHeader").textContent =
      "Edit items Entry";

    // Set buttons
    document.getElementById("itemsEntrySubmitButton").textContent = "Save";
    document.getElementById("itemsEntryFormPopup").onsubmit = () =>
      submitEdititemsEntry(event);

    // Set Content
    document.getElementById("popup_items_id").value = itemsEntry.item_id;
    document.getElementById("popup_items_upc").value = itemsEntry.upc;
    document.getElementById("popup_items_invoicenumber").value =
      itemsEntry.invoice_number;
    document.getElementById("popup_items_vendorselector").value =
      itemsEntry.vendor_id;
    document.getElementById("popup_items_name").value = itemsEntry.name;
    document.getElementById("popup_items_brand").value = itemsEntry.brand;
    document.getElementById("popup_items_description").value =
      itemsEntry.description;
    document.getElementById("popup_items_price").value = itemsEntry.price;
    document.getElementById("popup_items_weighed").value = itemsEntry.weighed;
    document.getElementById("popup_items_count").value = itemsEntry.count;
    document.getElementById("popup_items_arrived").value = formatDateForInput(
      itemsEntry.arrived_at,
    );

    // Set layers
    document.getElementById("popup_div_items_arrived").style.display = "block";
    if (itemsEntry.deleted) {
      // only show the soldout date if the item is deleted
      document.getElementById("popup_div_items_soldout").style.display =
        "block";
    } else {
      document.getElementById("popup_div_items_soldout").style.display = "none";
    }

    document.getElementById("itemsPopup").style.display = "block";
    document.getElementById("itemsOverlay").style.display = "block";
  });
}

// Add items Entry from Popup
function submitAdditemsEntryPopup(e) {
  e.preventDefault();
  const itemsEntry = {
    upc: document.getElementById("popup_items_upc").value,
    invoice_number: document.getElementById("popup_items_invoicenumber").value,
    vendor_id: parseInt(
      document.getElementById("popup_items_vendorselector").value,
      10,
    ),
    name: document.getElementById("popup_items_name").value,
    brand: document.getElementById("popup_items_brand").value,
    description: document.getElementById("popup_items_description").value,
    price: parseFloat(document.getElementById("popup_items_price").value),
    weighed: document.getElementById("popup_items_weighed").checked,
    count: parseInt(document.getElementById("popup_items_count").value, 10),
  };

  if (itemsEntry.price == 0) {
    document.getElementById("editError").style.display = "block";
    document.getElementById("editError").textContent = "Invalid Price";
    return;
  }

  axios
    .post("/api/items", itemsEntry)
    .then(() => {
      loadItems();
      closeInfo();
      closeItemsPopup();
      document.getElementById("itemsEntryFormPopup").reset();
    })
    .catch((err) => {
      document.getElementById("editError").style.display = "block";
      if (itemsEntry.vendor_id == -1) {
        document.getElementById("editError").textContent = "Invalid Vendor";
      } else {
        document.getElementById("editError").textContent =
          "Error Processing Your Request";
      }
    });
}

// Submit edit form
function submitEdititemsEntry(e) {
  e.preventDefault();
  const id = document.getElementById("popup_items_id").value;
  const itemsEntry = {
    upc: document.getElementById("popup_items_upc").value,
    invoice_number: document.getElementById("popup_items_invoicenumber").value,
    vendor_id: parseInt(
      document.getElementById("popup_items_vendorselector").value,
      10,
    ),
    brand: document.getElementById("popup_items_brand").value,
    name: document.getElementById("popup_items_name").value,
    description: document.getElementById("popup_items_description").value,
    price: parseFloat(document.getElementById("popup_items_price").value),
    weighed: document.getElementById("popup_items_weighed").checked,
    count: parseInt(document.getElementById("popup_items_count").value),
    arrived_at: formatDateForSQL(
      document.getElementById("popup_items_arrived").value,
    ),
  };

  axios
    .put(`/api/items/${id}`, itemsEntry)
    .then(() => {
      closeItemsPopup();
      closeInfo();
      loadItems();
      document.getElementById("itemsEntryFormPopup").reset();
    })
    .catch((err) => {
      document.getElementById("editError").style.display = "block";
      if (itemsEntry.vendor_id == -1) {
        document.getElementById("editError").textContent = "Invalid Vendor";
      } else {
        document.getElementById("editError").textContent =
          "Error Processing Your Request";
      }
    });
}

// Close the Add items Popup
function closeItemsPopup() {
  document.getElementById("editError").style.display = "none";
  document.getElementById("editError").textContent = "";
  document.getElementById("itemsPopup").style.display = "none";
  document.getElementById("itemsOverlay").style.display = "none";
}

// Populate Vendors selector
function loadVendors() {
  return axios.get("/api/vendors").then((res) => {
    const select = document.getElementById("popup_items_vendorselector");
    select.innerHTML = "";

    res.data.forEach((v) => {
      const opt = document.createElement("option");
      opt.value = v.vendor_id;
      opt.text = v.vendor_name;
      select.appendChild(opt);
    });
  });
}
