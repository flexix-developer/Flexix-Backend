window.onload = function () {
  fetch("http://ceproject.thddns.net:3323/allbookforbag", {
    method: "GET", // สามารถเปลี่ยนเป็น 'POST', 'PUT', ถ้ามีความจำเป็น
    headers: {},
  })
    .then((response) => response.json())
    .then((data) => {
      const APIData = data.product;
      const sourceElement = document.getElementById("col-2");
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

          // Modify text for ID
          if (child.id.includes("ID")) {
            child.textContent = item.ID; // Set the new text
          }
          // // Check and change src for image-1
          if (child.tagName === "IMG" && child.id.includes("image-1")) {
            child.src = item.Product_Image; // Set the new src
          }
          // Modify text for BookName
          if (child.id.includes("BookName")) {
            child.textContent = item.Product_Name; // Set the new text
          }
          // Modify text for ProductPrice
          if (child.id.includes("ProductPrice")) {
            child.textContent = item.Product_Price; // Set the new text
          }
          // Modify text for AuthorName
          if (child.id.includes("AuthorName")) {
            child.textContent = item.Author_Name; // Set the new text
          }
          // Modify text for CateName
          if (child.id.includes("CateName")) {
            child.textContent = item.Category_Name; // Set the new text
          }
          // Modify text for InStock
          if (child.id.includes("InStock")) {
            child.textContent = item.Product_In_Stock; // Set the new text
          }

          // // Check and change src for button-1
          if (child.tagName === "BUTTON" && child.id.includes("button-1")) {
            child.value = item.ID; // Set the new src
            child.dataset.id = item.ID;
          }
        });
        container.appendChild(clonedElement); // Add the clonedElement to the container
      });
    })
    .catch((error) => console.error("Error:", error));
};
