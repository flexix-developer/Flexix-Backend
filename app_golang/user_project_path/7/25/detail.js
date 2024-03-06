// detail.js
document.addEventListener("DOMContentLoaded", function () {
  const urlParams = new URLSearchParams(window.location.search);
  const param = urlParams.get("id"); // ดึงค่า ID
  console.log(param); // แสดงค่า ID ใน console
fetch(`http://ceproject.thddns.net:3326/book/getByID?id=${param}`, {
      method: "GET", // สามารถเปลี่ยนเป็น 'POST', 'PUT', ถ้ามีความจำเป็น
      headers: {'Authorization':'Bearer bf86eb551a1caaa206d0436c8cce47b9'},
    })
  .then((response) => response.json())
  .then((data) => {
    console.log(data);
    // อัปเดต UI ตามข้อมูลที่ได้
          
 document.getElementById("P-IMG").src = data.product_Image;
  })
  .catch((error) => console.error("Error loading product details:", error));

});
