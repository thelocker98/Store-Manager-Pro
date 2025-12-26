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
