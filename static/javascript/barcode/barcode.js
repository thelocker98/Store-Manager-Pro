document.addEventListener("DOMContentLoaded", function () {
  const statusEl = document.getElementById("status");
  const messagesEl = document.getElementById("messages");

  // Automatically use ws or wss based on page protocol
  const protocol = window.location.protocol === "https:" ? "wss" : "ws";
  const wsUrl = `${protocol}://${window.location.host}/api/ws`;

  const socket = new WebSocket(wsUrl);

  socket.onopen = () => {
    statusEl.textContent = "Connected";
    statusEl.style.color = "green";
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
  };

  socket.onerror = () => {
    statusEl.textContent = "Error";
    statusEl.style.color = "red";
  };
});
