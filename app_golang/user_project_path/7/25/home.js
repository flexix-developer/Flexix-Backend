window.onload = function () {
    fetch("http://ceproject.thddns.net:3326/book/getTopBook"
    , {
        method: "GET", // สามารถเปลี่ยนเป็น 'POST', 'PUT', ถ้ามีความจำเป็น
        headers: {'Authorization':'Bearer bf86eb551a1caaa206d0436c8cce47b9'},
      })
      .then((response) => response.json())
      .then((data) => {
        const APIData =  data.product
        const sourceElement = document.getElementById("row-7");
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
          
// // Check and change src for image-3
    if (child.tagName === "IMG" && child.id.includes("image-3")) {
    child.src = item.product_Image; // Set the new src
  }          
// // Check and change src for button-1
                   if (child.id.includes("button-1")) {
                    child.addEventListener("click", function () {
                      window.location.href = `detail.html?id=${item.product_Image}`;
                      console.log(item.ID);
                    });
                  }
        
// // Check and change src for button-1
    if (child.tagName === "BUTTON" && child.id.includes("button-1")) {
    child.value = item.id; // Set the new src
    child.dataset.id = item.id;
    }          
// // Check and change src for button-1
                   if (child.id.includes("button-1")) {
                    child.addEventListener("click", function () {
                      window.location.href = `detail.html?id=${item.id}`;
                      console.log(item.ID);
                    });
                  }


            
          });
          container.appendChild(clonedElement); // Add the clonedElement to the container
        });
      })
      .catch((error) => console.error("Error:", error));
  };