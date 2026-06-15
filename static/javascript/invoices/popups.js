// Upload Popup

function openUploadPopup() {
  console.log("Opening");
  document.getElementById("uploadPopup").style.display = "flex";
  document.getElementById("uploadOverlay").style.display = "block";
  console.log("Opened");
}

function closeUploadPopup() {
  document.getElementById("uploadPopup").style.display = "none";
  document.getElementById("uploadOverlay").style.display = "none";
}
//

// Function to upload file to Go backend
async function uploadFile(file) {
  // Validate PDF
  if (!file.name.toLowerCase().endsWith(".pdf") && file.type !== "application/pdf") {
    showUploadError("Invalid file type - PDF required");
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
  console.log("Upload success:", message);
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
});
