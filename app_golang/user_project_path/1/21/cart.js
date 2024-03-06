window.onload = function () {
    fetch("http://ceproject.thddns.net:3323/get_inbucket"
    , {
        method: "GET", // สามารถเปลี่ยนเป็น 'POST', 'PUT', ถ้ามีความจำเป็น
        headers: {},
      })
      .then((response) => response.json())
      .then((data) => {
        const APIData =  data.product
        const sourceElement = document.getElementById("row-5");
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
          
// // Check and change src for Pimg
    if (child.tagName === "IMG" && child.id.includes("Pimg")) {
    child.src = item.Product_Image; // Set the new src
  }          
// Modify text for Pprice
  if (child.id.includes("Pprice")) {
    child.textContent = item.Product_Price; // Set the new text
  }          
// Modify text for Pname
  if (child.id.includes("Pname")) {
    child.textContent = item.Product_Name; // Set the new text
  }          
// Modify text for Pprice
  if (child.id.includes("Pprice")) {
    child.textContent = item.Product_Price; // Set the new text
  }          
// Modify text for PQuantity
  if (child.id.includes("PQuantity")) {
    child.textContent = item.Quntity; // Set the new text
  }          
// Modify text for PTotal
  if (child.id.includes("PTotal")) {
    child.textContent = item.Total_Price; // Set the new text
  }
        
// // Check and change src for BTDelete
    if (child.tagName === "BUTTON" && child.id.includes("BTDelete")) {
    child.value = item.Book_ID; // Set the new src
    child.dataset.id = item.Book_ID;
    }


            
          });
          container.appendChild(clonedElement); // Add the clonedElement to the container
        });
      })
      .catch((error) => console.error("Error:", error));
  };