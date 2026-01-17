async function printList(list_id) {
  try {
    // Fetch the PDF
    const response = await fetch("/api/export/list/" + list_id);
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
