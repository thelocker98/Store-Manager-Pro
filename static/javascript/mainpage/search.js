async function searchItems() {
  const query = document.getElementById("searchInput").value;
  const tbody = document.querySelector("#itemsTable tbody");
  const noResultsContainer = document.getElementById("noResultsContainer");

  // Build API URL
  const sortBy = document.getElementById("sortBySelect").value;
  const showDeleted = document.getElementById("showDeletedCheckbox").checked;

  try {
    const res = await axios.get(
      `http://localhost:8080/api/search/${encodeURIComponent(query)}?showdeleted=${showDeleted}&sortby=${sortBy}&page=${page}&pagesize=${pageSize}`,
    );

    const items = res.data; // assuming API returns an array of items

    if (!items) {
      if (page != 1) {
        page--;
        reloadData();
      }
      // No results found → show Add Item button
      tbody.innerHTML = "";
      noResultsContainer.style.display = "block";
      return 1;
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
      // Open Info on click
      row.onclick = () => openInfo(item.item_id);

      tbody.appendChild(row);
    });
    return items[0].number_of_entrys;
  } catch (error) {
    console.error("Search error:", error);
  }
}

function resetSearch() {
  noResultsContainer.style.display = "none";
  const input = document.getElementById("searchInput");
  input.value = "";

  reloadData();
}

document.addEventListener("DOMContentLoaded", function () {
  const searchInput = document.getElementById("searchInput");
  const showDeletedSearchInput = document.getElementById("showDeletedCheckbox");
  const sortBySearchInput = document.getElementById("sortBySelect");

  if (searchInput) {
    searchInput.addEventListener("keypress", function (e) {
      if (e.key === "Enter") {
        reloadData();
      }
    });
  }

  showDeletedSearchInput.addEventListener("change", (e) => {
    if (searchInput.value != "") {
      reloadData();
    }
  });

  sortBySearchInput.addEventListener("change", (e) => {
    if (searchInput.value != "") {
      reloadData();
    }
  });
});
