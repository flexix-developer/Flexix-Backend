package otp

import (
	"flexix_backend/app_golang/orm"
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
)

func sendOTPEmail(recipientEmail string) string {
	// สุ่มตัวเลข OTP
	randints := randOTP()

user := orm.User{
    OTP: randints, // กำหนดค่า OTP ที่คุณต้องการ
}

orm.Db.Model(&user).Where("email = ?", recipientEmail).Update("otp", user.OTP)




	// fmt.Println("randints",randints)
	if randints == "" {
		fmt.Println("Error generating OTP")
		return ""
	}
	fmt.Printf("randints: %v, type: %T\n", randints, randints)

	// ข้อมูล SMTP
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	senderEmail := "flexix.general@gmail.com"
	senderPassword := "gngwhqjznhefbnrk"

	// ข้อมูลข้อความอีเมล
	subject := "OTP Reset Password"
	// body := fmt.Sprintf(`
	// 	<!DOCTYPE html>
	// 	<html>
	// 	<head>
	// 		<link rel="stylesheet" type="text/css" hs-webfonts="true" href="https://fonts.googleapis.com/css?family=Lato|Lato:i,b,bi">
	// 		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
	// 		<meta name="viewport" content="width=device-width, initial-scale=1.0">
	// 		<style type="text/css">
	// 			h1{{font-size:56px}}
	// 			h2{{font-size:28px;font-weight:900}}
	// 			p{{font-weight:100}}
	// 			td{{vertical-align:top}}
	// 			#email{{margin:auto;width:600px;background-color:#fff}}
	// 			table{{justify-content:center;}}
	// 		</style>
	// 	</head>
	// 	<body bgcolor="#253C70" style="width: 100%; font-family:Lato, sans-serif; font-size:18px;">
	// 	<div id="email" bgcolor="#253C70">
	// 		<table role="presentation" width="100%" style="justify-content: center;">
	// 			<tr>
	// 				<td bgcolor="#253C70" align="center" style="color: white;">
	// 					<h1> OTP </h1>
	// 				</td>
	// 		</table>
	// 		<table role="presentation" border="0" cellpadding="0" cellspacing="10px" width="100%">
	// 			<tbody><tr>
	// 				<td align="center" bgcolor="#253C70">
					
	// 					<h2 style="color: white;">%s</h2>
	// 				</td>
	// 			</tr>
	// 		</tbody></table>
	// 	</div>
	// 	</body>
	// 	</html>
	// `, randints)

	body := randints

	// เชื่อมต่อกับเซิร์ฟเวอร์ SMTP
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	// สร้างข้อมูลสำหรับส่งอีเมล
	msg := []byte(fmt.Sprintf("To: %s\r\n", recipientEmail) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		body)

	// ส่งอีเมล
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, senderEmail, []string{recipientEmail}, msg)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return ""
	}

	fmt.Println("Email sent successfully!")
	return randints
}

func randOTP() string {
	// สุ่มตัวเลข OTP
	randInt := rand.Intn(900000) + 100000
	fmt.Println(randInt)
	return strconv.Itoa(randInt)
}
