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
          
 document.getElementById("P-IMG").src = data.Product_Image;          
document.getElementById("P-Name").textContent = data.Product_Name;          
document.getElementById("P-Author").textContent = data.Author_Name;          
document.getElementById("P-Price").textContent = data.Product_Price;          
document.getElementById("P-Detail").textContent = data.Product_Detail;          
document.getElementById("P-Cate").textContent = data.Category_Name;
  })
  .catch((error) => console.error("Error loading product details:", error));

});
