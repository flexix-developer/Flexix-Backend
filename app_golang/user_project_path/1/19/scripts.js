เลือกทุก < a > ในเอกสาร;
var allLinks = document.querySelectorAll("a");

// วนลูปผ่านทุก <a> เพื่อลบคำสั่ง onClick
allLinks.forEach(function (link) {
  // ลบคำสั่ง onClick ออก
  link.removeAttribute("onClick");
});
