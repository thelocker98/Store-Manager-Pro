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
  document.getElementById("popup_items_id").value = "";
  document.getElementById("popup_items_upc").value = "";
  document.getElementById("popup_items_invoicenumber").value = "";
  document.getElementById("popup_items_name").value = "";
  document.getElementById("popup_items_brand").value = "";

  // Set layers
  document.getElementById("popup_div_items_arrived").style.display = "none";
  document.getElementById("popup_div_items_soldout").style.display = "none";
  document.getElementById("itemsPopup").style.display = "block";
  document.getElementById("itemsOverlay").style.display = "block";
}

// Open the edit popup and fill fields
function openEdititemsEntry(id) {
  // Load vendor selector
  loadVendors();

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
    document.getElementById("popup_items_id").value = itemsEntry.items_id;
    document.getElementById("popup_items_upc").value = itemsEntry.upc;
    document.getElementById("popup_items_invoicenumber").value =
      itemsEntry.invoice_number;
    document.getElementById("popup_items_vendorselector").value =
      itemsEntry.vendor_id;
    document.getElementById("popup_items_name").value = itemsEntry.name;
    document.getElementById("popup_items_brand").value = itemsEntry.brand;

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
    ),
    name: document.getElementById("popup_items_name").value,
    brand: document.getElementById("popup_items_brand").value,
  };
  axios.post("/api/items", itemsEntry).then(() => {
    loaditems();
    document.getElementById("itemsEntryFormPopup").reset();
    closeitemsPopup();
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
    ),
    brand: document.getElementById("popup_items_brand").value,
    name: document.getElementById("popup_items_name").value,
  };

  axios.put(`/api/items/${id}`, itemsEntry).then(() => {
    closeitemsPopup();
    closeInfo();
    loaditems();
  });
}

// Close the Add items Popup
function closeitemsPopup() {
  document.getElementById("itemsPopup").style.display = "none";
  document.getElementById("itemsOverlay").style.display = "none";
}

// Populate Vendors selector
function loadVendors() {
  axios.get("/api/vendors").then((res) => {
    const v = document.getElementById("popup_items_vendorselector");
    v.innerHTML = "";
    res.data.forEach((vendor) => {
      v.innerHTML += `<option value="${vendor.vendor_id}">${vendor.vendor_name}</option>`;
    });
  });
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
