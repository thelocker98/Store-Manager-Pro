// Automatically use ws or wss based on page protocol
const protocol = window.location.protocol === "https:" ? "wss" : "ws";
const wsUrl = `${protocol}://${window.location.host}/api/ws`;
let socket = new WebSocket(wsUrl);

function connectWebSocket() {
  socket = new WebSocket(wsUrl);

  // Set the link to the barcode reader to gree to show there is no problems with the websocket
  socket.onopen = () => {
    hideDisconnectPopup();
  };

  // Set the link to the barcode reader to red to show there is a problem with the websocket
  socket.onerror = () => {
    showDisconnectPopup();
  };

  socket.onmessage = (event) => {
    // Remove the first and last character of the barcode result before searching it
    val = event.data.slice(1, event.data.length - 2);

    // Set search feild to barcode reader value
    document.getElementById("searchInput").value = val;

    // Perform the search
    reloadData();
  };

  // Set the link to the barcode reader to red to show there is a problem with the websocket
  socket.onclose = () => {
    showDisconnectPopup();
  };

  // Set the link to the barcode reader to red to show there is a problem with the websocket
  socket.onerror = () => {
    showDisconnectPopup();
  };
}

document.addEventListener("DOMContentLoaded", function () {
  connectWebSocket();
});
