// Variables
var page = 1;
var numberOfPages = 1;
var pageSize = 25;

function nextPage() {
  page++;
  reloadData();
}

function prevPage() {
  page--;
  reloadData();
}

async function reloadData() {
  // Variables
  const searchBox = document.getElementById("searchInput").value;
  var entryCount;

  // Do Search or Data Load
  if (searchBox != "") {
    entryCount = await searchItems();
  } else {
    entryCount = (await loadInvoice()) || 101;
  }

  // Find out number of pages for result
  numberOfPages = Math.ceil(entryCount / pageSize) || 1;

  // Clamp page to max
  if (page > numberOfPages) page = numberOfPages;
  // Clamp page to minimum of 1
  if (page < 1) page = 1;

  // Update buttons
  document.getElementById("prevPageBtn").disabled = page <= 1;
  document.getElementById("nextPageBtn").disabled = page >= numberOfPages;

  // Update indicator
  document.getElementById("pageIndicator").textContent = `Page ${page}/${numberOfPages}`;
}
