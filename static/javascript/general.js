function formatPrice(price, weighed) {
  if (weighed) {
    return "$" + price.toFixed(2) + "/lb";
  } else {
    return "$" + price.toFixed(2);
  }
}
function formatDate(value) {
  const d = new Date(value);
  return d.toLocaleDateString("en-US", {
    month: "2-digit",
    day: "2-digit",
    year: "2-digit",
  });
}

function formatDateForInput(value) {
  if (!value) return "";
  const d = new Date(value);

  const year = d.getFullYear(); // 4-digit year
  const month = String(d.getMonth() + 1).padStart(2, "0"); // 01–12
  const day = String(d.getDate()).padStart(2, "0"); // 01–31

  return `${year}-${month}-${day}`; // "YYYY-MM-DD"
}

function formatDateForSQL(dateValue) {
  if (!dateValue) return ""; // handle empty input
  const [year, month, day] = dateValue.split("-");

  // Create a Date object at noon UTC
  const d = new Date(Date.UTC(year, month - 1, day, 12, 0, 0));

  return d.toISOString(); // returns "2026-01-01T12:00:00.000Z"
}

// Create overlay and popup elements
const overlay = document.createElement("div");
overlay.id = "connectionOverlay";
overlay.style.cssText = `
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: none;
    z-index: 1999;
  `;

const popup = document.createElement("div");
popup.id = "connectionPopup";
popup.style.cssText = `
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background: white;
    padding: 30px;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    text-align: center;
    display: none;
    z-index: 2000;
    min-width: 300px;
  `;

popup.innerHTML = `
    <h2 style="margin: 0 0 15px 0; color: #d32f2f;">Disconnected</h2>
    <p style="margin: 0 0 20px 0; color: #666;">Lost connection to the server</p>
    <button id="recheckBtn" onclick="connectWebSocket()" style="
      background-color: #1976d2;
      color: white;
      border: none;
      padding: 10px 20px;
      border-radius: 4px;
      cursor: pointer;
      font-size: 16px;
    ">Recheck Connection</button>
  `;

function showDisconnectPopup() {
  overlay.style.display = "block";
  popup.style.display = "block";
}

function hideDisconnectPopup() {
  overlay.style.display = "none";
  popup.style.display = "none";
}

document.addEventListener("DOMContentLoaded", function () {
  document.body.appendChild(overlay);
  document.body.appendChild(popup);

  // Also monitor online/offline events
  window.addEventListener("offline", () => {
    showDisconnectPopup();
  });

  window.addEventListener("online", () => {
    hideDisconnectPopup();
  });
});
