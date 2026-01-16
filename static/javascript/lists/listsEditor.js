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

function searchListItems() {
  const query = document.getElementById("searchInput").value;
  const tbody = document.querySelector("#searchResultsTable tbody");

  // Build API URL
  const sortBy = document.getElementById("sortBySelect").value;
  if (query == "") {
    document.getElementById("searchResultsTable").style.display = "none";
    return;
  }

  axios
    .get(
      `/api/search/${encodeURIComponent(query)}?showdeleted=false&sortby=${sortBy}&page=1&pagesize=15&list_id=${ListID}`,
    )
    .then((res) => {
      // assuming API returns an array of items
      const items = res.data;

      if (!items) {
        return;
      }

      // show table
      document.getElementById("searchResultsTable").style.display = "table";

      // clear table body
      tbody.innerHTML = "";

      // Add rows
      items.forEach((item) => {
        const row = document.createElement("tr");
        if (item.list_count != -1) {
          row.style.backgroundColor = "#00ff00";
        }

        row.innerHTML = `
                <td>${item.location_name || ""}</td>
                <td>${item.brand || ""}</td>
                <td>${item.name || ""}</td>
                <td>${item.description || ""}</td>
                <td>${formatPrice(item.price, item.weighed) || ""}</td>
            `;
        // Open Info on click
        row.onclick = () => addToList(item.item_id);

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
    })
    .catch((err) => {
      console.log("error adding item to list in database");
    });
}

document.addEventListener("DOMContentLoaded", function () {
  loadListName();

  const searchInput = document.getElementById("searchInput");
  const sortBySearchInput = document.getElementById("sortBySelect");

  if (searchInput) {
    searchInput.addEventListener("keypress", function (e) {
      if (e.key === "Enter") {
        searchListItems();
      }
    });
  }

  sortBySearchInput.addEventListener("change", (e) => {
    if (searchInput.value != "") {
      searchListItems();
    }
  });
});
