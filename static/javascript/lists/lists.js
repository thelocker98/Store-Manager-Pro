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
                         <span class="list-card-delete material-symbol">&#xE872;</span>
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

    container.appendChild(card);
  });
}

function viewList(listId) {
  // Navigate to the list detail page
  window.location.href = `/lists/${listId}`;
}

function createNewList() {
  // Handle creating a new list
  window.location.href = "/lists/new";
}

// Load lists when page loads
document.addEventListener("DOMContentLoaded", () => {
  renderLists();
});

function deleteList(list_id) {
  if (confirm("Are you sure you want to permanently delete this list?")) {
    axios.delete(`/api/lists/${list_id}`).then(() => renderLists());
    closeInfo();
    closeItemPopup();
  }
}
