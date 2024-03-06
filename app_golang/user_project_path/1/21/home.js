window.onload = function () {
  fetch("http://ceproject.thddns.net:3323/token/getbestsellerbook", {
    method: "GET", // สามารถเปลี่ยนเป็น 'POST', 'PUT', ถ้ามีความจำเป็น
    headers: {
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc3MzQyNzUsInVzZXJJZCI6MX0.vZCgeTAFCwyRsvRRy-cL7zTxUxokucy9gtLHAviwEkI",
    },
  })
    .then((response) => response.json())
    .then((data) => {
      const APIData = data.product;
      console.log(APIData);
      const sourceElement = document.getElementById("row-4");
      const container = sourceElement.parentNode;
      container.innerHTML = ""; // Clear the container to prepare for new elements
      APIData.forEach((item, i) => {
        const clonedElement = sourceElement.cloneNode(true);
        clonedElement.id = `${sourceElement.id}`; // No need to escape backticks here
        // Customize the clonedElement as necessary
        clonedElement.querySelectorAll("*").forEach((child, index) => {
          const newId = `${child.id}`; // No need to escape backticks here
          child.id = newId; // Set the new id
          // Check and change src for images

          // // Check and change src for P-IMG
          if (child.tagName === "IMG" && child.id.includes("P-IMG")) {
            child.src = item.Product_Image; // Set the new src
          }
          // // Check and change src for BT-Detail
          if (child.id.includes("BT-Detail")) {
            child.addEventListener("click", function () {
              window.location.href = `detail.html?id=${item.ID}`;
              console.log(item.ID);
            });
          }
          // Modify text for P-Name
          if (child.id.includes("P-Name")) {
            child.textContent = item.Product_Name; // Set the new text
          }
          // // Check and change src for BT-Detail
          if (child.id.includes("BT-Detail")) {
            child.addEventListener("click", function () {
              window.location.href = `detail.html?id=${item.ID}`;
              console.log(item.ID);
            });
          }
          // Modify text for P-Price
          if (child.id.includes("P-Price")) {
            child.textContent = item.Product_Price; // Set the new text
          }
          // // Check and change src for BT-Detail
          if (child.id.includes("BT-Detail")) {
            child.addEventListener("click", function () {
              window.location.href = `detail.html?id=${item.ID}`;
              console.log(item.ID);
            });
          }

          // // Check and change src for BT-Detail
          if (child.tagName === "BUTTON" && child.id.includes("BT-Detail")) {
            child.value = item.ID; // Set the new src
            child.dataset.id = item.ID;
          }
          // // Check and change src for BT-Detail
          if (child.id.includes("BT-Detail")) {
            child.addEventListener("click", function () {
              window.location.href = `detail.html?id=${item.ID}`;
              console.log(item.ID);
            });
          }
        });
        container.appendChild(clonedElement); // Add the clonedElement to the container
      });
    })
    .catch((error) => console.error("Error:", error));
};
