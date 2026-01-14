async function renderLists() {
  const res = await axios.get(`/api/lists`);
  const lists = res.data;

  const container = document.getElementById("listsContainer");
  container.innerHTML = "";

  lists.forEach((list) => {
    const card = document.createElement("div");
    card.className = "list-card";
    card.onclick = () => viewList(list.list_id);
    //#xE872;
    card.innerHTML = `
                     <div class="list-card-header">
                         <h3 class="list-card-title">${list.list_name}</h3>
                         <div>
                          <span class="list-card-print material-symbol">&#xE8AD;</span>
                          <span class="list-card-delete material-symbol">&#xE872;</span>
                         </div>
                     </div>
                     <div class="list-card-footer">
                         <span>Created ${formatDate(list.created_at)}</span>
                     </div>
                 `;
    // Add delete handler with stopPropagation
    const deleteBtn = card.querySelector(".list-card-delete");
    deleteBtn.onclick = (e) => {
      e.stopPropagation(); // Prevents the card click from firing
      deleteList(list.list_id);
    };

    // Add print handler with stopPropagation
    const printBtn = card.querySelector(".list-card-print");
    printBtn.onclick = (e) => {
      e.stopPropagation(); // Prevents the card click from firing
      printList(list.list_id);
    };

    // Add list card to the grid
    container.appendChild(card);
  });
}

// Function to delete List
function deleteList(list_id) {
  if (confirm("Are you sure you want to permanently delete this list?")) {
    axios.delete(`/api/lists/${list_id}`).then(() => renderLists());
  }
}

// Function to print the list
function printList(list_id) {
  console.log("List ID to print:", list_id);
  printPDF();
}

// Function to go to the list
function viewList(listId) {
  // Navigate to the list detail page
  window.location.href = `/lists/${listId}`;
}

// Load lists when page loads
document.addEventListener("DOMContentLoaded", () => {
  renderLists();
});

// list Popup
function openList(id) {
  document.getElementById("listError").style.display = "none";
  document.getElementById("listPopup").style.display = "block";
  document.getElementById("listOverlay").style.display = "block";
}

function closeList() {
  document.getElementById("listPopup").style.display = "none";
  document.getElementById("listOverlay").style.display = "none";
}

function createNewList() {
  const list = {
    list_name: document.getElementById("popup_list_title").value,
  };

  axios
    .post("/api/lists", list)
    .then(() => {
      closeList();
      document.getElementById("popup_list_title").value = "";
      renderLists();
    })
    .catch((err) => {
      document.getElementById("listError").style.display = "block";
      document.getElementById("listError").textContent = "Error Creating List";
    });
}

async function printPDF() {
  try {
    // Fetch the PDF
    const response = await fetch("/static/inventory.pdf");
    const blob = await response.blob();

    // Create a URL for the blob
    const blobUrl = URL.createObjectURL(blob);

    // Open in new window
    const printWindow = window.open(blobUrl, "_blank");

    if (!printWindow) {
      alert("Please allow popups to print the PDF");
      URL.revokeObjectURL(blobUrl);
      return;
    }

    // Wait for PDF to load, then print
    printWindow.onload = function () {
      printWindow.print();

      // Clean up when window closes
      printWindow.onbeforeunload = function () {
        URL.revokeObjectURL(blobUrl);
      };
    };
  } catch (error) {
    console.error("Error printing PDF:", error);
    alert("Failed to load PDF for printing");
  }
}
