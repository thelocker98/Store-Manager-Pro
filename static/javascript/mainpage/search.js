async function searchItems() {
  const query = document.getElementById("searchInput").value;
  const tbody = document.querySelector("#itemsTable tbody");
  const noResultsContainer = document.getElementById("noResultsContainer");

  try {
    const res = await axios.get(
      `/api/search/${encodeURIComponent(query)}?showdeleted=${ShowDeleted}&orderby=${OrderBy}&page=${page}&pagesize=${pageSize}&vendor=${Vendor}&location=${Location}&department=${Department}&filterlist_id=${Listid}`,
    );

    // assuming API returns an array of items
    const items = res.data;

    // check if the items array is empty
    if (!items) {
      if (page != 1) {
        page--;
        reloadData();
      }
      // No results found -> show Add Item button
      tbody.innerHTML = "";
      noResultsContainer.style.display = "block";
      closePriceSection();

      return 1;
    } else {
      noResultsContainer.style.display = "none";
    }

    // check if their is only one result then show it large
    if (items.length == 1) {
      openPriceSection(items[0]);
    } else {
      closePriceSection();
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
                <td>${item.location_name || ""}</td>
                <td>${item.brand || ""}</td>
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

function openPriceSection(item) {
  document.getElementById("priceSectionTitle").textContent =
    item.name + " (" + item.brand + ")";
  document.getElementById("priceSectionPrice").textContent =
    "Price:" + formatPrice(item.price, item.weighed);
  document.getElementById("priceSection").style.display = "block";

  console.log(item);
}

function closePriceSection() {
  document.getElementById("priceSection").style.display = "none";
}

document.addEventListener("DOMContentLoaded", function () {
  const searchInput = document.getElementById("searchInput");

  if (searchInput) {
    searchInput.addEventListener("keypress", function (e) {
      if (e.key === "Enter") {
        reloadData();
      }
    });
  }
});
