// Open the Add items Popup
function openAdditemPopup() {
  // Load vendor and location selector
  loadVendors();
  loadLocations();
  // Set Title
  document.getElementById("itemPopupHeader").textContent = "Add items Entry";

  // Set Buttons
  document.getElementById("itemsEntrySubmitButton").textContent = "Add";
  document.getElementById("itemEntryFormPopup").onsubmit = () =>
    submitAdditemsEntryPopup(event);

  // This chunk of code checks if the search field has only numbers and
  // if their is no result found and it this is true it will automatically
  // fill in the upc field on the create item popup

  if (BarcodeValue.slice(0, 1) == "2") {
    searchupc = BarcodeValue.slice(0, BarcodeValue.length - 6) + "00000";
  } else {
    searchupc = BarcodeValue.slice(0, BarcodeValue.length - 1);
  }

  if (
    !(
      /^\d+$/.test(searchupc) &&
      document.getElementById("noResultsContainer").style.display == "block"
    )
  ) {
    searchupc = "";
  }

  // Set Content
  document.getElementById("popup_items_id").value = 0;
  // Check if searchupc has text
  document.getElementById("popup_items_upc").value = searchupc;
  document.getElementById("popup_items_invoicenumber").value = "";
  document.getElementById("popup_items_vendorselector").value = 1;
  document.getElementById("popup_items_locationselector").value = 1;
  document.getElementById("popup_items_name").value = "";
  document.getElementById("popup_items_brand").value = "";
  document.getElementById("popup_items_description").value = "";
  document.getElementById("popup_items_price").value = 0;
  document.getElementById("popup_items_weighed").value = false;
  document.getElementById("popup_items_count").value = 0;

  // Set layers
  document.getElementById("popup_div_items_arrived").style.display = "none";
  document.getElementById("popup_div_items_soldout").style.display = "none";
  document.getElementById("itemPopup").style.display = "block";
  document.getElementById("itemsOverlay").style.display = "block";
}

// Open the edit popup and fill fields
async function openEdititemsEntry(id) {
  // Load vendor selector
  await loadVendors();

  try {
    // Check that entry exists in database
    const res = await axios.get(`/api/items/${id}`);

    const itemsEntry = res.data;
    if (!itemsEntry) return;

    // Load locations based on department
    await loadLocations(itemsEntry.department_id);

    // Set Title
    document.getElementById("itemPopupHeader").textContent = "Edit Item Entry";

    // Set buttons
    document.getElementById("itemsEntrySubmitButton").textContent = "Save";
    document.getElementById("itemEntryFormPopup").onsubmit = () =>
      submitEdititemsEntry(event);

    // Set Content
    document.getElementById("popup_items_id").value = itemsEntry.item_id;
    document.getElementById("popup_items_upc").value = itemsEntry.upc;
    document.getElementById("popup_items_invoicenumber").value =
      itemsEntry.invoice_number;
    document.getElementById("popup_items_vendorselector").value =
      itemsEntry.vendor_id;

    document.getElementById("popup_items_locationselector").value =
      itemsEntry.location_id;
    document.getElementById("popup_items_name").value = itemsEntry.name;
    document.getElementById("popup_items_brand").value = itemsEntry.brand;
    document.getElementById("popup_items_description").value =
      itemsEntry.description;
    document.getElementById("popup_items_price").value = itemsEntry.price;
    document.getElementById("popup_items_weighed").checked = itemsEntry.weighed;
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

    document.getElementById("itemPopup").style.display = "block";
    document.getElementById("itemsOverlay").style.display = "block";
  } catch (err) {
    console.error(err);
  }
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
    location_id: parseInt(
      document.getElementById("popup_items_locationselector").value,
      10,
    ),
    department_id: GlobalDepartment,
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
      closeInfo();
      closeItemPopup();
      document.getElementById("itemEntryFormPopup").reset();

      // clear the search box of a barcode search
      searchupc = document.getElementById("searchInput").value;
      if (
        /^\d+$/.test(searchupc) &&
        document.getElementById("noResultsContainer").style.display == "block"
      ) {
        document.getElementById("searchInput").value = "";
      }
      reloadData();
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
    location_id: parseInt(
      document.getElementById("popup_items_locationselector").value,
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
      closeItemPopup();
      closeInfo();
      reloadData();
      document.getElementById("itemEntryFormPopup").reset();
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
function closeItemPopup() {
  document.getElementById("editError").style.display = "none";
  document.getElementById("editError").textContent = "";
  document.getElementById("itemPopup").style.display = "none";
  document.getElementById("itemsOverlay").style.display = "none";
}
