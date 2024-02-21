window.onload = function () {
    fetch("http://127.0.0.1:5000/api"
    , {
        method: "GET", // สามารถเปลี่ยนเป็น 'POST', 'PUT', ถ้ามีความจำเป็น
        headers: {},
      })
      .then((response) => response.json())
      .then((data) => {
        const sourceElement = document.getElementById("col-1");
        const container = sourceElement.parentNode;
        container.innerHTML = ""; // Clear the container to prepare for new elements
        data.forEach((item, i) => {
          const clonedElement = sourceElement.cloneNode(true);
          clonedElement.id = `${sourceElement.id}`; // No need to escape backticks here
          // Customize the clonedElement as necessary
          clonedElement.querySelectorAll("*").forEach((child, index) => {
            
            const newId = `${child.id}`; // No need to escape backticks here
            child.id = newId; // Set the new id
            // Check and change src for images
          
// // Check and change src for image-1
    if (child.tagName === "IMG" && child.id.includes("image-1")) {
    child.src = item.Product_Image; // Set the new src
  }          
// // Check and change src for button-1
                   if (child.id.includes("button-1")) {
                    child.addEventListener("click", function () {
                      window.location.href = `Detail.html?id=${item.ID}`;
                      console.log(item.ID);
                    });
                  }


            
          });
          container.appendChild(clonedElement); // Add the clonedElement to the container
        });
      })
      .catch((error) => console.error("Error:", error));
  };