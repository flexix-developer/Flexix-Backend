// detail.js
document.addEventListener("DOMContentLoaded", function () {
  const urlParams = new URLSearchParams(window.location.search);
  const param = urlParams.get("id"); // ดึงค่า ID
  console.log(param); // แสดงค่า ID ใน console
  fetch(`http://127.0.0.1:5000/api/book?id=${param}`, {
    method: "GET", // สามารถเปลี่ยนเป็น 'POST', 'PUT', ถ้ามีความจำเป็น
    headers: {},
  })
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      // อัปเดต UI ตามข้อมูลที่ได้

      document.getElementById("image-6").src = data.Product_Image;
      document.getElementById("text-6").textContent = data.Product_Name;
    })
    .catch((error) => console.error("Error loading product details:", error));
});
