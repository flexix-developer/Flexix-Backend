
const add_author = async () => {
  // รับค่าจาก input fields
  const atname = document.getElementById("input-1").value;


  // สร้าง object สำหรับส่งข้อมูล
  const dataBody = {
atname:atname,
  };

  try {
    const response = await fetch("http://127.0.0.1:5000/api/author", {
      method: "POST",
            headers: {
        "Content-Type": "application/json", // กำหนด Content-Type header เป็น application/json
        Authorization: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc3MzQyNzUsInVzZXJJZCI6MX0.vZCgeTAFCwyRsvRRy-cL7zTxUxokucy9gtLHAviwEkI'
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

