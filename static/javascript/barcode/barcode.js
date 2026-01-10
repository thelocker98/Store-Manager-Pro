document.addEventListener("DOMContentLoaded", function () {
  const statusEl = document.getElementById("status");
  const messagesEl = document.getElementById("messages");

  // Create overlay and popup elements
  const overlay = document.createElement("div");
  overlay.id = "connectionOverlay";
  overlay.style.cssText = `
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.3);
    display: none;
    z-index: 999;
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
    z-index: 1000;
    min-width: 300px;
  `;

  popup.innerHTML = `
    <h2 style="margin: 0 0 15px 0; color: #d32f2f;">Disconnected</h2>
    <p style="margin: 0 0 20px 0; color: #666;">Lost connection to the server</p>
    <button id="recheckBtn" style="
      background-color: #1976d2;
      color: white;
      border: none;
      padding: 10px 20px;
      border-radius: 4px;
      cursor: pointer;
      font-size: 16px;
    ">Recheck Connection</button>
  `;

  document.body.appendChild(overlay);
  document.body.appendChild(popup);

  function showDisconnectPopup() {
    overlay.style.display = "block";
    popup.style.display = "block";
  }

  function hideDisconnectPopup() {
    overlay.style.display = "none";
    popup.style.display = "none";
  }

  // Automatically use ws or wss based on page protocol
  const protocol = window.location.protocol === "https:" ? "wss" : "ws";
  const wsUrl = `${protocol}://${window.location.host}/api/ws`;
  let socket = new WebSocket(wsUrl);

  function connectWebSocket() {
    socket = new WebSocket(wsUrl);

    socket.onopen = () => {
      statusEl.textContent = "Connected";
      statusEl.style.color = "green";
      hideDisconnectPopup();
    };

    socket.onmessage = (event) => {
      const div = document.createElement("div");
      div.className = "message";
      div.textContent = event.data;
      messagesEl.appendChild(div);
      messagesEl.scrollTop = messagesEl.scrollHeight;
    };

    socket.onclose = () => {
      statusEl.textContent = "Disconnected";
      statusEl.style.color = "red";
      showDisconnectPopup();
    };

    socket.onerror = () => {
      statusEl.textContent = "Error";
      statusEl.style.color = "red";
      showDisconnectPopup();
    };
  }

  connectWebSocket();

  // Recheck connection button handler
  document.getElementById("recheckBtn").addEventListener("click", () => {
    if (
      socket.readyState === WebSocket.CLOSED ||
      socket.readyState === WebSocket.CLOSING
    ) {
      connectWebSocket();
    }
  });
});
