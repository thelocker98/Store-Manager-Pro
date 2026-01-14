function loadListName() {
  const display = document.getElementById("listName");

  // Load List Name from Database
  axios
    .get(`/api/lists/${ListID}`)
    .then((res) => {
      display.textContent = res.data.list_name;
    })
    .catch((err) => {
      console.log("error updating list name in database");
    });
}

function editListName() {
  const display = document.getElementById("listName");
  const input = document.getElementById("listNameInput");
  // Load List Name from Database
  axios
    .get(`/api/lists/${ListID}`)
    .then((res) => {
      input.value = res.data.list_name;
    })
    .catch((err) => {
      return;
    });

  display.style.display = "none";
  input.style.display = "block";
  input.focus();
  input.select();

  input.onblur = () => saveListName();

  input.addEventListener("keypress", function (e) {
    if (e.key === "Enter") saveListName();
  });
}

function saveListName() {
  const display = document.getElementById("listName");
  const input = document.getElementById("listNameInput");

  // quit if name is empty
  if (input.value == "") {
    display.style.display = "block";
    input.style.display = "none";
    return;
  }

  // Save to Database
  axios
    .put(`/api/lists/${ListID}`, { list_name: input.value })
    .then(() => {
      display.textContent = input.value;
    })
    .catch((err) => {
      console.log("error updating list name in database");
    });

  // Exit title editor
  display.style.display = "block";
  input.style.display = "none";
}

document.addEventListener("DOMContentLoaded", function () {
  loadListName();
});
