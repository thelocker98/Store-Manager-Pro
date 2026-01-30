// Values
var BarcodeValue = "";

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
    BarcodeValue = event.data;

    // Check if the info overlay or the item overlay is active so that searchs don't take place in the background
    if (
      document.getElementById("infooverlay").style.display == "none" &&
      document.getElementById("itemsOverlay").style.display == "none"
    ) {
      // Set search feild to barcode reader value
      // If the item begans with a 2 then ignore the last 5 digits
      if (BarcodeValue.slice(0, 1) == "2") {
        document.getElementById("searchInput").value =
          BarcodeValue.slice(1, BarcodeValue.length - 6) + "00000";
      } else {
        document.getElementById("searchInput").value = BarcodeValue.slice(
          1,
          BarcodeValue.length - 1,
        );
      }

      reloadData();

      // Check it the item overlay is active to automatically fill in the upc value with the barcode value
    } else if (
      document.getElementById("itemsOverlay").style.display != "none"
    ) {
      // fill in the upc value with the barcode reader result
      // If the item begans with a 2 then ignore the last 5 digits
      if (BarcodeValue.slice(0, 1) == "2") {
        document.getElementById("popup_items_upc").value =
          BarcodeValue.slice(0, BarcodeValue.length - 6) + "00000";
      } else {
        document.getElementById("popup_items_upc").value = BarcodeValue.slice(
          0,
          BarcodeValue.length - 1,
        );
      }
    }
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
