











const Delete_Book = async () => {
  try {
    const id = document.getElementById("DeleteBookIDInput").value;
    const response = await fetch(`http://ceproject.thddns.net:3323/api/delete_book?id=${id}`, {
      method: "DELETE",
            headers: {
        "Content-Type": "application/json", // กำหนด Content-Type header เป็น application/json
        
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

const Update_Book = async () => {
  // รับค่าจาก input fields
  const name = document.getElementById("UpdateBookNameInput").value;
  const price = document.getElementById("UpdateBookPriceInput").value;
  const image = document.getElementById("UpdateBookImageInput").value;
  const id = document.getElementById("UpdateBookIDInput").value;


  // สร้าง object สำหรับส่งข้อมูล
  const dataBody = {
name:name,
  price:price,
  image:image,
  id:id,
  };

  try {
    const response = await fetch("http://ceproject.thddns.net:3323/api/update_book", {
      method: "PUT",
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

const Add_Book = async () => {
  // รับค่าจาก input fields
  const name = document.getElementById("BookNameInput").value;
  const price = document.getElementById("PriceInput").value;
  const detail = document.getElementById("DetailInput").value;
  const in_stock = document.getElementById("InStockInput").value;
  const image = document.getElementById("ImageInput").value;
  const author_id = document.getElementById("AuthorIDInput").value;
  const cate_id = document.getElementById("CateInput").value;


  // สร้าง object สำหรับส่งข้อมูล
  const dataBody = {
name:name,
  price:price,
  detail:detail,
  in_stock:in_stock,
  image:image,
  author_id:author_id,
  cate_id:cate_id,
  };

  try {
    const response = await fetch("http://ceproject.thddns.net:3323/api/addbook", {
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

