const add_bbok = async () => {
  // รับค่าจาก input fields
  const name = document.getElementById("input-1").value;
  const price = document.getElementById("input-2").value;
  const detail = document.getElementById("input-3").value;
  const in_stock = document.getElementById("input-4").value;
  const image = document.getElementById("input-5").value;
  const author_name = document.getElementById("input-6").value;
  const category_id = document.getElementById("input-7").value;


  // สร้าง object สำหรับส่งข้อมูล
  const dataBody = {
name:name,
  price:price,
  detail:detail,
  in_stock:in_stock,
  image:image,
  author_name:author_name,
  category_id:category_id,
  };

  try {
    const response = await fetch("http://ceproject.thddns.net:3326/book/add", {
      method: "POST",
            headers: {
        "Content-Type": "application/json", // กำหนด Content-Type header เป็น application/json
        Authorization: 'Bearer bf86eb551a1caaa206d0436c8cce47b9'
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

const delete_book = async () => {
  try {
    const id = document.getElementById("input-8").value;
    const response = await fetch(`http://ceproject.thddns.net:3326/book/deleteBook?id=${id}`, {
      method: "DELETE",
            headers: {
        "Content-Type": "application/json", // กำหนด Content-Type header เป็น application/json
        Authorization: 'Bearer bf86eb551a1caaa206d0436c8cce47b9'
      },
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

const update_book = async () => {
  // รับค่าจาก input fields
  const id = document.getElementById("input-9").value;
  const name = document.getElementById("input-10").value;
  const price = document.getElementById("input-11").value;
  const image = document.getElementById("input-12").value;


  // สร้าง object สำหรับส่งข้อมูล
  const dataBody = {
id:id,
  name:name,
  price:price,
  image:image,
  };

  try {
    const response = await fetch("http://ceproject.thddns.net:3326/book/update", {
      method: "PUT",
            headers: {
        "Content-Type": "application/json", // กำหนด Content-Type header เป็น application/json
        Authorization: 'Bearer bf86eb551a1caaa206d0436c8cce47b9'
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

