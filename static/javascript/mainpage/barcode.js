document.addEventListener("DOMContentLoaded", function () {
  const statusEl = document.getElementById("barcodeLink");
  console.log(statusEl);

  // Automatically use ws or wss based on page protocol
  const protocol = window.location.protocol === "https:" ? "wss" : "ws";
  const wsUrl = `${protocol}://${window.location.host}/api/ws`;

  const socket = new WebSocket(wsUrl);

  // Set the link to the barcode reader to gree to show there is no problems with the websocket
  socket.onopen = () => {
    statusEl.style.color = "green";
  };

  socket.onmessage = (event) => {
    // Remove the first and last character of the barcode result before searching it
    val = event.data.slice(1, event.data.length - 2);

    // Check if the info overlay or the item overlay is active so that searchs don't take place in the background
    if (
      document.getElementById("infooverlay").style.display == "none" &&
      document.getElementById("itemsOverlay").style.display == "none"
    ) {
      // Set search feild to barcode reader value
      document.getElementById("searchInput").value = val;
      // Perform the search
      reloadData();

      // Check it the item overlay is active to automatically fill in the upc value with the barcode value
    } else if (
      document.getElementById("itemsOverlay").style.display != "none"
    ) {
      // fill in the upc value with the barcode reader result
      document.getElementById("popup_items_upc").value = val;
    }
  };

  // Set the link to the barcode reader to red to show there is a problem with the websocket
  socket.onclose = () => {
    statusEl.style.color = "red";
  };

  // Set the link to the barcode reader to red to show there is a problem with the websocket
  socket.onerror = () => {
    statusEl.style.color = "red";
  };
});
