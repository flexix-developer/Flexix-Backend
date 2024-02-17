const post = async () => {
  const atnameValue = document.getElementById("input-3").value;

  try {
    const response = await fetch("http://127.0.0.1:5000/api/author", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        atname: atnameValue,
      }),
    });
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    console.log("Data:", data);
  } catch (error) {
    console.error("Error fetching data:", error);
  }
};

const gost = async () => {
  const atnameValue = document.getElementById("input-3").value;

  try {
    const response = await fetch("http://127.0.0.1:5000/api/author", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        atname: atnameValue,
      }),
    });
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    console.log("Data:", data);
  } catch (error) {
    console.error("Error fetching data:", error);
  }
};
const bad = async () => {
  const atnameValue = document.getElementById("input-3").value;

  try {
    const response = await fetch("http://127.0.0.1:5000/api/author", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        atname: atnameValue,
      }),
    });
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    console.log("Data:", data);
  } catch (error) {
    console.error("Error fetching data:", error);
  }
};
const sssost = async () => {
  const atnameValue = document.getElementById("input-3").value;

  try {
    const response = await fetch("http://127.0.0.1:5000/api/author", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        atname: atnameValue,
      }),
    });
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    console.log("Data:", data);
  } catch (error) {
    console.error("Error fetching data:", error);
  }
};
const get_books = async () => {
  try {
    const response = await fetch("http://127.0.0.1:5000/token/api/book?id=1", {
      method: "GET",
      headers: {
        "Content-Type": "application/json", // กำหนด Content-Type header เป็น application/json
        Authorization:
          "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc3MzQyNzUsInVzZXJJZCI6MX0.vZCgeTAFCwyRsvRRy-cL7zTxUxokucy9gtLHAviwEkI",
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
