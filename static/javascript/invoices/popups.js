var GlobalMarkupAmount = 0;

// Upload Popup
let refreshInterval = null; // Global variable to store interval ID
function openUploadPopup() {
  document.getElementById("uploadPopup").style.display = "flex";
  document.getElementById("uploadOverlay").style.display = "block";

  // Start refreshing every 10 seconds
  renderNotifications(); // Render immediately
  refreshInterval = setInterval(renderNotifications, 10000); // Then every 10 seconds
}

function closeUploadPopup() {
  document.getElementById("uploadPopup").style.display = "none";
  document.getElementById("uploadOverlay").style.display = "none";

  // Stop refreshing
  if (refreshInterval) {
    clearInterval(refreshInterval);
    refreshInterval = null;
  }
  // Reload Data
  reloadData();
}

// Function to upload file to Go backend
async function uploadFile(file) {
  // Validate PDF
  if (!file.name.toLowerCase().endsWith(".pdf") && file.type !== "application/pdf") {
    showUploadError("Invalid file type - PDF required");
    return false;
  }
  // Validate Markup
  var markup = document.getElementById("markupAmount").value;
  if (markup < 0) {
    showUploadError("Invalid Markup");
    return false;
  }

  // Show loading state
  const uploadArea = document.getElementById("uploadArea");
  const originalContent = uploadArea.innerHTML;
  uploadArea.innerHTML =
    '<div style="text-align: center;">Uploading...<div class="spinner"></div></div>';

  // Prepare form data
  const formData = new FormData();
  formData.append("file", file);
  formData.append("markup", markup);

  try {
    const response = await fetch(`/api/invoices/file/${InvoiceID}`, {
      method: "POST",
      body: formData,
    });

    const result = await response.json();

    if (response.ok) {
      showUploadSuccess(result.message || "File uploaded successfully!");
      uploadArea.innerHTML = originalContent;
      return true;
    } else {
      showUploadError(result.error || "Upload failed");
      uploadArea.innerHTML = originalContent;
      return false;
    }
  } catch (error) {
    showUploadError("Network error: " + error.message);
    uploadArea.innerHTML = originalContent;
    return false;
  }
}

// Function to handle file picker
function uploadFilePicker() {
  // Create hidden file input
  const fileInput = document.createElement("input");
  fileInput.type = "file";
  fileInput.accept = ".pdf";
  fileInput.style.display = "none";

  fileInput.addEventListener("change", (e) => {
    const file = e.target.files[0];
    if (file) {
      uploadFile(file);
    }
    document.body.removeChild(fileInput);
  });

  document.body.appendChild(fileInput);
  fileInput.click();
}

// Show Error Messages
function showUploadError(message) {
  const messageElement = document.getElementById("uploadMessage");
  if (messageElement) {
    messageElement.classList.add("errorMessage");
    messageElement.textContent = message;
    setTimeout(() => {
      messageElement.textContent = "";
      messageElement.classList.remove("errorMessage");
    }, 3000);
  }
  console.error("Upload error:", message);
}

// Show Success Messages
function showUploadSuccess(message) {
  const messageElement = document.getElementById("uploadMessage");
  if (messageElement) {
    messageElement.classList.add("successMessage");
    messageElement.textContent = message;
    setTimeout(() => {
      messageElement.textContent = "";
      messageElement.classList.remove("successMessage");
    }, 5000);
  }
  renderNotifications();
}

// Populate Markup Amount
function loadMarkupAmount() {
  document.getElementById("markupAmount").value = GlobalMarkupAmount;
}

async function renderNotifications() {
  const res = await axios.get(`/api/invoices/notification/${InvoiceID}`);
  const ocrEntrys = res.data;

  const container = document.getElementById("notification-container");
  container.innerHTML = "";

  if (!Array.isArray(ocrEntrys)) {
    document.getElementById("notification-header").style.display = "none";
    return;
  }
  document.getElementById("notification-header").style.display = "block";

  ocrEntrys.forEach((ocrEntry) => {
    const card = document.createElement("div");
    card.className = "notification-card";

    card.innerHTML = `
    <button id="delete-btn-${ocrEntry.file_id}" class="delete-notification-btn" onclick="deleteNotification(this, ${ocrEntry.file_id})">
      ✕
    </button>
    <div class="card-header">
      <h3>${ocrEntry.file_path}</h3>
      <p class="status-badge" id="status-badge-${ocrEntry.file_id}"></p>
    </div>

    <div class="card-row">
      <p>
      <strong>Duration: </strong>
      ${
        ocrEntry.started_at && ocrEntry.completed_at
          ? `${((new Date(ocrEntry.completed_at) - new Date(ocrEntry.started_at)) / 1000).toFixed(
              0,
            )} sec`
          : "Waiting..."
      }
      </p>
    </div>
    <div class="card-row">
      <p>
        <strong>Entrys Failed: </strong>${ocrEntry.entries_failed}
      </p>
      <p>
        <strong>Entrys Added: </strong>${ocrEntry.entries_added}
      </p>
    </div>
    <p style="color: Red" id="error-message-${ocrEntry.file_id}"></p>
    `;

    container.appendChild(card);

    error_message = document.getElementById(`error-message-${ocrEntry.file_id}`);
    if (ocrEntry.status == 2) {
      error_message.style.display = "block";
      error_message.textContent = `error: ${ocrEntry.error_message}`;
    } else {
      error_message.style.display = "none";
      error_message.textContent = "";
    }

    status_badge = document.getElementById(`status-badge-${ocrEntry.file_id}`);

    switch (ocrEntry.status) {
      case 0:
        // Handle Pending status
        status_badge.textContent = "Pending";
        status_badge.style.backgroundColor = "Yellow";
        card.style.backgroundColor = "LightYellow";

        break;

      case 1:
        // Handle Processing status
        status_badge.textContent = "Processing";
        status_badge.style.backgroundColor = "Blue";
        card.style.backgroundColor = "LightBlue";
        document.getElementById(`delete-btn-${ocrEntry.file_id}`).remove();

        break;

      case 2:
        // Handle Failed status
        status_badge.textContent = "Failed";
        status_badge.style.backgroundColor = "Red";
        card.style.backgroundColor = "LightRed";

        break;

      case 3:
        // Handle Succeeded status
        status_badge.textContent = "Succeeded";
        status_badge.style.backgroundColor = "Green";
        card.style.backgroundColor = "LightGreen";

        break;

      default:
        // Handle unknown status
        status_badge.style.display = "none";
        break;
    }
  });
}
function deleteNotification(btn, file_id) {
  axios
    .delete(`/api/invoices/notification/${file_id}`)
    .then(() => {
      const card = btn.closest(".notification-card");
      card.style.transition = "opacity 0.3s, transform 0.3s";
      card.style.opacity = "0";
      card.style.transform = "scale(0.95)";

      setTimeout(() => {
        card.remove();
      }, 300);
    })
    .catch((err) => {
      return;
    });
}

function updateMarkupAmount() {
  GlobalMarkupAmount = Number(document.getElementById("markupAmount").value);
  document.cookie = "markupAmount=" + GlobalMarkupAmount + "; max-age=315360000; path=/";
}

// Complete DOMContentLoaded with all features
document.addEventListener("DOMContentLoaded", function () {
  const dropzoneContainer = document.getElementById("uploadArea");

  if (!dropzoneContainer) {
    console.error("uploadArea element not found!");
    return;
  }

  // Make the dropzone clickable for file picker
  dropzoneContainer.addEventListener("click", () => {
    uploadFilePicker();
  });

  // Drag and drop handlers
  dropzoneContainer.addEventListener("dragover", (e) => {
    e.preventDefault();
    dropzoneContainer.classList.add("uploadDrag");
  });

  dropzoneContainer.addEventListener("drop", (e) => {
    e.preventDefault();
    dropzoneContainer.classList.remove("uploadDrag");

    let file = null;

    if (e.dataTransfer.items) {
      // Using DataTransferItemList
      const items = e.dataTransfer.items;
      for (let i = 0; i < items.length; i++) {
        if (items[i].kind === "file") {
          file = items[i].getAsFile();
          break; // Take only the first file
        }
      }
    } else {
      // Using DataTransfer (fallback)
      file = e.dataTransfer.files[0];
    }

    if (file) {
      console.log("File dropped:", file.name);
      uploadFile(file);
    } else {
      showUploadError("No file detected");
    }
  });

  dropzoneContainer.addEventListener("dragleave", (e) => {
    dropzoneContainer.classList.remove("uploadDrag");
  });

  // Markup
  GlobalMarkupAmount = Number(getCookie("markupAmount"));
  // get cookie or set the default to 30%
  if (isNaN(GlobalMarkupAmount)) {
    document.getElementById("markupAmount").value = 30;
    updateMarkupAmount();
  }

  const markupAmountSelector = document.getElementById("markupAmount");

  if (markupAmountSelector) {
    markupAmountSelector.addEventListener("change", () => {
      updateMarkupAmount();
      console.log("Markup Changed");
    });
  }

  loadMarkupAmount();
});
