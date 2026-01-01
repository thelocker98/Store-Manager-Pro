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
