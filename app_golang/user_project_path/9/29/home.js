window.onload = function () {
    fetch("http://127.0.0.1:5000/news"
    , {
        method: "GET", // สามารถเปลี่ยนเป็น 'POST', 'PUT', ถ้ามีความจำเป็น
        headers: {},
      })
      .then((response) => response.json())
      .then((data) => {
        const APIData =  data
        const sourceElement = document.getElementById("row-16");
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
          
// Modify text for name
  if (child.id.includes("name")) {
    child.textContent = item.Name; // Set the new text
  }          
// Modify text for DATE
  if (child.id.includes("DATE")) {
    child.textContent = item.Date; // Set the new text
  }


            
          });
          container.appendChild(clonedElement); // Add the clonedElement to the container
        });
      })
      .catch((error) => console.error("Error:", error));
  };