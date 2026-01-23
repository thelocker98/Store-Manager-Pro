function loadListName() {
  const display = document.getElementById("listName");

  // Load List Name from Database
  axios
    .get(`/api/lists/${ListID}`)
    .then((res) => {
      display.textContent = res.data.list_name;
    })
    .catch((err) => {
      console.log("error updating list name in database");
    });
}

function loadList() {
  const tbody = document.querySelector("#listTable tbody");

  axios.get(`/api/listentrys/${ListID}`).then((res) => {
    // assuming API returns an array of items
    const items = res.data;

    // clear table body
    tbody.innerHTML = "";

    const headsUpContainer = document.getElementById("headsUpContainer");
    if (!items) {
      headsUpContainer.style.display = "flex";
      return;
    } else {
      headsUpContainer.style.display = "none";
    }

    // Add rows
    items.forEach((item) => {
      const row = document.createElement("tr");

      row.innerHTML = `
                <td>${item.location_name || ""}</td>
                <td>${item.brand || ""}</td>
                <td>${item.name || ""}</td>
                <td>${item.description || ""}</td>
                <td>${formatPrice(item.price, item.weighed) || ""}</td>
                <td style="width: 110px; font-size: 25px;">
                    <div style="display: flex; align-items: center; justify-content: space-between;">
                        <div class="material-symbol hoverExpand" onclick="deincrementListEntry(${item.entry_id})" style="cursor: pointer; font-size:25px;">&#xE15B;</div>
                        <span>${item.list_count}</span>
                        <div class="material-symbol hoverExpand" onclick="incrementListEntry(${item.entry_id})" style="cursor: pointer; font-size:25px;">&#xE145;</div>
                        <div class="material-symbol hoverExpand" onclick="deleteListEntry(${item.entry_id})" style="cursor: pointer; font-size:25px;">&#xE872;</div>
                    </div>
                </td>
            `;

      tbody.appendChild(row);
    });
  });
}

function editListName() {
  const display = document.getElementById("listName");
  const input = document.getElementById("listNameInput");
  // Load List Name from Database
  axios
    .get(`/api/lists/${ListID}`)
    .then((res) => {
      input.value = res.data.list_name;
    })
    .catch((err) => {
      return;
    });

  display.style.display = "none";
  input.style.display = "block";
  input.focus();
  input.select();

  input.onblur = () => saveListName();

  input.addEventListener("keypress", function (e) {
    if (e.key === "Enter") saveListName();
  });
}

function saveListName() {
  const display = document.getElementById("listName");
  const input = document.getElementById("listNameInput");

  // quit if name is empty
  if (input.value == "") {
    display.style.display = "block";
    input.style.display = "none";
    return;
  }

  // Save to Database
  axios
    .put(`/api/lists/${ListID}`, { list_name: input.value })
    .then(() => {
      display.textContent = input.value;
    })
    .catch((err) => {
      console.log("error updating list name in database");
    });

  // Exit title editor
  display.style.display = "block";
  input.style.display = "none";
}

function reloadData() {
  searchListItems(false);
}

function searchListItems(UseScanner) {
  const query = document.getElementById("searchInput").value;
  const tbody = document.querySelector("#searchResultsTable tbody");

  // clear table body
  tbody.innerHTML = "";

  // Build API URL
  if (query == "") {
    return;
  }

  axios
    .get(
      `/api/search/${encodeURIComponent(query)}?showdeleted=${ShowDeleted}&orderby=${OrderBy}&page=0&pagesize=10&vendor=${Vendor}&location=${Location}&department=${Department}&filterlist_id=${Listid}&list_id=${ListID}`,
    )
    .then((res) => {
      // assuming API returns an array of items
      const items = res.data;

      if (!items) {
        return;
      }

      if (items.length == 1 && UseScanner) {
        if (items[0].entry_id == -1) {
          addToList(items[0].item_id);
        } else {
          // Item exits so add another
          incrementListEntry(items[0].entry_id);
        }
        reloadData();
        return;
      }

      // show table
      document.getElementById("searchResultsTable").style.display = "table";

      // Add rows
      items.forEach((item) => {
        const row = document.createElement("tr");

        if (item.deleted == true) {
          row.className = "deleted";
        } else {
          row.className = "avalible";
        }

        var html = `
                <td>${item.location_name || ""}</td>
                <td>${item.brand || ""}</td>
                <td>${item.name || ""}</td>
                <td>${item.description || ""}</td>
                <td>${formatPrice(item.price, item.weighed) || ""}</td>
                <td style="width: 110px; font-size: 25px;">
            `;

        if (item.list_count != -1) {
          html += `
                    <div style="display: flex; align-items: center; justify-content: space-between;">
                        <div class="material-symbol hoverExpand" onclick="deincrementListEntry(${item.entry_id})" style="cursor: pointer; font-size:25px;">&#xE15B;</div>
                        <span>${item.list_count}</span>
                        <div class="material-symbol hoverExpand" onclick="incrementListEntry(${item.entry_id})" style="cursor: pointer; font-size:25px;">&#xE145;</div>
                        <div class="material-symbol hoverExpand" onclick="deleteListEntry(${item.entry_id})" style="cursor: pointer; font-size:25px;">&#xE872;</div>
                    </div>
                </td>
            `;
        } else {
          html += `
                    <button onclick="addToList(${item.item_id});">Add Item</button>
                </td>
            `;
        }

        // Update the row html
        row.innerHTML = html;

        tbody.appendChild(row);
      });
    });
}

function resetSearchList() {
  const input = document.getElementById("searchInput");
  input.value = "";

  document.getElementById("searchResultsTable").style.display = "none";
}

function addToList(item_id) {
  // add item to list Database
  axios
    .post(`/api/listentrys/${ListID}`, { item_id: item_id })
    .then(() => {
      searchListItems();
      loadList();
    })
    .catch((err) => {
      console.log("error adding item to list in database");
    });
}

function deincrementListEntry(entryID) {
  axios
    .get(`/api/listentry/${entryID}`)
    .then((res) => {
      data = res.data;
      data.list_count--;
      if (data.list_count <= 0) {
        deleteListEntry(entryID);
      }

      axios
        .put(`/api/listentrys/${entryID}`, data)
        .then((res) => {
          searchListItems();
          loadList();
        })
        .catch((err) => {
          return;
        });
    })
    .catch((err) => {
      return;
    });
}

function incrementListEntry(entryID) {
  axios
    .get(`/api/listentry/${entryID}`)
    .then((res) => {
      data = res.data;
      data.list_count++;
      axios
        .put(`/api/listentrys/${entryID}`, data)
        .then((res) => {
          searchListItems();
          loadList();
        })
        .catch((err) => {
          return;
        });
    })
    .catch((err) => {
      return;
    });
}

function deleteListEntry(entryID) {
  axios
    .delete(`/api/listentrys/${entryID}`)
    .then((res) => {
      searchListItems();
      loadList();
    })
    .catch((err) => {
      return;
    });
}

document.addEventListener("DOMContentLoaded", function () {
  loadListName();
  loadList();
  loadFiltersDropdown();
  document.getElementById("showDeletedCheckbox").checked = false;
  applyFilter();

  const searchInput = document.getElementById("searchInput");
  const filterDropDown = document.getElementById("filter-dropdown-menu");

  if (searchInput) {
    searchInput.addEventListener("keypress", function (e) {
      if (e.key === "Enter") {
        if (document.getElementById("searchInput").value == "") {
          resetSearchList();
        } else {
          searchListItems();
        }
      }
    });
  }

  filterDropDown.addEventListener("mouseleave", () => {
    applyFilter();
  });
  filterDropDown.addEventListener("mouseenter", () => {
    loadFiltersDropdown();
  });
});
