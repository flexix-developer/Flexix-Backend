from django.shortcuts import render

# Create your views here.
from django.views.decorators.csrf import csrf_exempt
from rest_framework.parsers import JSONParser
from django.http.response import JsonResponse
from UserApp.serializers import StudentSerializer
from UserApp.models import User
from UserApp.sendotp import send_otp_email
from django.core.exceptions import ObjectDoesNotExist
from django.contrib.auth import authenticate, login, logout


@csrf_exempt
def UserApi(request,id=0):
    if request.method=='GET':
        student = User.objects.all()
        student_serializer=StudentSerializer(student,many=True)
        return JsonResponse(student_serializer.data,safe=False)
    elif request.method=='POST':
        student_data=JSONParser().parse(request)
        # print("student_data",student_data)
        student_serializer=StudentSerializer(data=student_data)
        print("student_serializer",student_serializer)
        if student_serializer.is_valid():
            user = student_serializer.save()
            # print(user)
            # login(request, user)
            return JsonResponse("Added Successfully",safe=False)
        return JsonResponse({"error": "Failed to Add"}, status=400)
    elif request.method=='PUT':
        student_data=JSONParser().parse(request)
        student=User.objects.get(ID=id)
        student_serializer=StudentSerializer(student,data=student_data)
        print("sdsdsds",student_serializer)
        if student_serializer.is_valid():
            student_serializer.save()
            return JsonResponse("Updated Successfully",safe=False)
        return JsonResponse("Failed to Update")
    elif request.method=='DELETE':
        student=User.objects.get(ID=id)
        student.delete()
        return JsonResponse("Deleted Successfully",safe=False)


otp_num = None  # ประกาศตัวแปร otp_num ไว้นอกฟังก์ชัน
entered_email = None

@csrf_exempt
def ForgotApi(request):
    global otp_num  # ให้ otp_num เป็นตัวแปร global
    if request.method == 'POST':
        # Assuming that the email is sent in the request body as 'Email'
        student_data = JSONParser().parse(request)
        email_to_check = student_data.get('Email', '')

        # Check if the email exists in the database
        user_with_email = User.objects.filter(Email=email_to_check).first()
        reqEmail = email_to_check
    
        if user_with_email:
            # Email exists, you can perform further actions here
            # ใช้ฟังก์ชันเพื่อส่ง OTP ไปยังอีเมล
            otp_num = send_otp_email(reqEmail)
            # print("otp_num",otp_num)

            # เพิ่มข้อมูลที่ต้องการส่งกลับใน JSON response
            response_data = {
                'message': 'Email exists in the system.',
                'otp_num': otp_num,
                'additional_info': 'Additional information for the frontend.'
            }

            return JsonResponse(response_data, status=200)
        else:
            # Email does not exist
            return JsonResponse({'message': 'Email not found in the system.'}, status=400)

@csrf_exempt
def VerifyOTPApi(request):
    global otp_num
    global entered_email
    # ให้ otp_num เป็นตัวแปร global
    print("otp_num",otp_num)
    try:
        # Assuming that the OTP is sent in the request body as 'otp_code'
        otp_data = JSONParser().parse(request)
        entered_otp = otp_data.get('otp_code', '')
        entered_email = otp_data.get('email', '')

        # ตรวจสอบค่า OTP ว่าตรงกับค่าที่ถูกส่งไปหรือไม่
        if entered_otp == otp_num:
            return JsonResponse({'message': 'OTP is valid.'}, status=200)
        else:
            return JsonResponse({'message': 'Invalid OTP.'}, status=400)

    except ObjectDoesNotExist:
        return JsonResponse({'message': 'OTP not found.'}, status=404)
    except Exception as e:
        return JsonResponse({'message': f'Error: {str(e)}'}, status=500)

@csrf_exempt   
def RePassApi(request):
    if request.method=='PUT':
        student_data = JSONParser().parse(request)
        email_to_check = student_data.get('Email', '')
        newPassword = student_data.get('Pass', '')
        user_with_email = User.objects.filter(Email=email_to_check).first()
        # print("user_with_email",user_with_email.ID)
        student_serializer=StudentSerializer(user_with_email,data={'Fname':user_with_email.Fname, 'Lname': user_with_email.Lname, 
                                                                   'Email': user_with_email.Email, 'Pass': newPassword})
        print(student_serializer)
        if student_serializer.is_valid():
            student_serializer.save()
            return JsonResponse("Updated Successfully",safe=False)
        return JsonResponse("Failed to Update")


from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from django.contrib.auth import authenticate, login
from .serializers import LoginSerializer

class LoginAPIView(APIView):
    def post(self, request, *args, **kwargs):
        serializer = LoginSerializer(data=request.data)
        serializer = LoginSerializer(data=request.data)
        print("Is serializer valid?", serializer.is_valid())
        print("Serializer errors:", serializer.errors)

        if serializer.is_valid():
            email = serializer.validated_data['email']
            password = serializer.validated_data['password']
            print(email,password)
            

            user = authenticate(request, username=email, password=password)
            print(user)

            if user:
                login(request, user)
                return Response({'message': 'Login successful'}, status=status.HTTP_200_OK)
            else:
                return Response({'message': 'Invalid credentials'}, status=status.HTTP_401_UNAUTHORIZED)
        else:
            return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)