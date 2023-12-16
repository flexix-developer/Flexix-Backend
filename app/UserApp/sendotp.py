import smtplib
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
import random

def rand_otp():
    random_number = random.randint(100000, 999999)
    return random_number

def send_otp_email(recipient_email):
    sender_email = "flexix.general@gmail.com"
    sender_password = "gngwhqjznhefbnrk"
    # ข้อมูลข้อความอีเมล
    subject = 'OTP Reset Password'
    randints = str(rand_otp())

    # สร้างอ็อบเจ็กต์ MIMEMultipart
    message = MIMEMultipart()
    message['From'] = sender_email
    message['To'] = recipient_email
    message['Subject'] = subject

    # เพิ่มเนื้อหาข้อความ
    html_part = MIMEText('''
    <!DOCTYPE html>
    <html>
    <head>
        <link rel="stylesheet" type="text/css" hs-webfonts="true" href="https://fonts.googleapis.com/css?family=Lato|Lato:i,b,bi">
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <style type="text/css">
            h1{{font-size:56px}}
            h2{{font-size:28px;font-weight:900}}
            p{{font-weight:100}}
            td{{vertical-align:top}}
            #email{{margin:auto;width:600px;background-color:#fff}}
            table{{justify-content:center;}}
        </style>
    </head>
    <body bgcolor="#253C70" style="width: 100%; font-family:Lato, sans-serif; font-size:18px;">
    <div id="email" bgcolor="#253C70">
        <table role="presentation" width="100%" style="justify-content: center;">
            <tr>
                <td bgcolor="#253C70" align="center" style="color: white;">
                    <h1> OTP </h1>
                </td>
        </table>
        <table role="presentation" border="0" cellpadding="0" cellspacing="10px" width="100%">
            <tbody><tr>
                <td align="center" bgcolor="#253C70">
                
                    <h2 style="color: white;">{}</h2>
                </td>
            </tr>
        </tbody></table>
    </div>
    </body>
    </html>
    '''.format(randints), 'html')
    message.attach(html_part)

    # เชื่อมต่อกับเซิร์ฟเวอร์ SMTP
    with smtplib.SMTP('smtp.gmail.com', 587) as server:
        server.starttls()
        server.login(sender_email, sender_password)

        # ส่งอีเมล
        server.sendmail(sender_email, recipient_email, message.as_string())

    print('Email sent successfully!')
    
    return randints

# # ใช้ฟังก์ชันเพื่อส่ง OTP ไปยังอีเมล
# send_otp_email('flexixtest@gmail.com')

