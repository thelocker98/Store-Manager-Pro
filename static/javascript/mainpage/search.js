async function searchItems() {
  const query = document.getElementById("searchInput").value;
  const tbody = document.querySelector("#itemsTable tbody");
  const noResultsContainer = document.getElementById("noResultsContainer");

  // If input is empty, turn search off
  if (!query) {
    noResultsContainer.style.display = "none";
    loadItems();
    return;
  }

  // Build API URL
  const sortBy = document.getElementById("sortBySelect").value;
  const showDeleted = document.getElementById("showDeletedCheckbox").checked;
  const url = `http://localhost:8080/api/search/${encodeURIComponent(query)}?showdeleted=${showDeleted}&sortby=${sortBy}&page=1&pagesize=50`;

  try {
    const res = await axios.get(url);
    const items = res.data; // assuming API returns an array of items

    if (!items) {
      // No results found → show Add Item button
      tbody.innerHTML = "";
      noResultsContainer.style.display = "block";
      return;
    } else {
      noResultsContainer.style.display = "none";
    }

    // clear table body
    tbody.innerHTML = ""; // clear previous results

    // Add rows
    items.forEach((item) => {
      const row = document.createElement("tr");

      if (item.deleted == true) {
        row.className = "deleted";
      } else {
        row.className = "avalible";
      }

      row.innerHTML = `
                <td>${item.upc || ""}</td>
                <td>${item.brand || ""}</td>
                <td>${item.vendor_name || ""}</td>
                <td>${item.name || ""}</td>
                <td>${item.description || ""}</td>
                <td>${formatPrice(item.price, item.weighed) || ""}</td>
                <td>${formatDate(item.arrived_at) || ""}</td>
            `;

      // optional: open info popup on row click
      row.onclick = () => openInfo(item.item_id);

      tbody.appendChild(row);
    });
  } catch (error) {
    console.error("Search error:", error);
  }
}

function resetSearch() {
  noResultsContainer.style.display = "none";
  const input = document.getElementById("searchInput");
  input.value = "";

  loadItems();
}

document.addEventListener("DOMContentLoaded", function () {
  const noResultsContainer = document.getElementById("noResultsContainer");
  const searchInput = document.getElementById("searchInput");

  if (searchInput) {
    searchInput.addEventListener("keypress", function (e) {
      if (e.key === "Enter") {
        searchItems();
      }
    });
  }
});
