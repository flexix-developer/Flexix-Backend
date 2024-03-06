// detail.js
document.addEventListener("DOMContentLoaded", function () {
  const urlParams = new URLSearchParams(window.location.search);
  const param = urlParams.get("id"); // ดึงค่า ID
  console.log(param); // แสดงค่า ID ใน console
fetch(`http://ceproject.thddns.net:3323/token/api/book?id=${param}`, {
      method: "GET", // สามารถเปลี่ยนเป็น 'POST', 'PUT', ถ้ามีความจำเป็น
      headers: {'Authorization':'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc3MzQyNzUsInVzZXJJZCI6MX0.vZCgeTAFCwyRsvRRy-cL7zTxUxokucy9gtLHAviwEkI'},
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
document.getElementById("BT-BUY").value = data.ID;          
document.getElementById("BT-ADD").value = data.ID;
  })
  .catch((error) => console.error("Error loading product details:", error));

});
const add_book_in_basket = async () => {
  // รับค่าจาก input fields
  const Book_ID = document.getElementById("BT-ADD").value;
  const Quantity = document.getElementById("Quantity").value;


  // สร้าง object สำหรับส่งข้อมูล
  const dataBody = {
Book_ID:Book_ID,
  Quantity:Quantity,
  };

  try {
    const response = await fetch("http://ceproject.thddns.net:3323/api/addbusbooks", {
      method: "POST",
            headers: {
        "Content-Type": "application/json", // กำหนด Content-Type header เป็น application/json
        
      },
      body: JSON.stringify(dataBody), // แปลง object เป็นสตริง JSON
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json(); // อ่าน JSON จาก response body
    console.log("Data:", data);
  } catch (error) {
    console.error("Error fetching data:", error);
  }
};

