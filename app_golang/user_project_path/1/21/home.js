window.onload = function () {
    fetch("http://ceproject.thddns.net:3326/book/getTopBook"
    , {
        method: "GET", // สามารถเปลี่ยนเป็น 'POST', 'PUT', ถ้ามีความจำเป็น
        headers: {'Authorization':'Bearer bf86eb551a1caaa206d0436c8cce47b9'},
      })
      .then((response) => response.json())
      .then((data) => {
        const APIData =  data.product
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
    child.src = item.product_Image; // Set the new src
  }          
// Modify text for P-Name
  if (child.id.includes("P-Name")) {
    child.textContent = item.Author_Name; // Set the new text
  }          
// Modify text for P-Price
  if (child.id.includes("P-Price")) {
    child.textContent = item.Product_Price; // Set the new text
  }


            
          });
          container.appendChild(clonedElement); // Add the clonedElement to the container
        });
      })
      .catch((error) => console.error("Error:", error));
  };