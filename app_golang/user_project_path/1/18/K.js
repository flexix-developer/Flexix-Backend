window.onload = function () {
    fetch("https://pokeapi.co/api/v2/pokemon"
    , {
        method: "GET", // สามารถเปลี่ยนเป็น 'POST', 'PUT', ถ้ามีความจำเป็น
        headers: {},
      })
      .then((response) => response.json())
      .then((data) => {
        const APIData =  data.results
        const sourceElement = document.getElementById("row-2");
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
          
// Modify text for text-1
  if (child.id.includes("text-1")) {
    child.textContent = item.name; // Set the new text
  }          
// // Check and change src for image-1
    if (child.tagName === "IMG" && child.id.includes("image-1")) {
    child.src = item.url; // Set the new src
  }


            
          });
          container.appendChild(clonedElement); // Add the clonedElement to the container
        });
      })
      .catch((error) => console.error("Error:", error));
  };