const add_news = async () => {
  // รับค่าจาก input fields
  const Name = document.getElementById("Name-Input").value;
  const Date = document.getElementById("Date-Input").value;
  const news = document.getElementById("News-Input").value;


  // สร้าง object สำหรับส่งข้อมูล
  const dataBody = {
Name:Name,
  Date:Date,
  news:news,
  };

  try {
    const response = await fetch("http://127.0.0.1:5000/add_news", {
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

