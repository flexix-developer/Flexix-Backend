from django.shortcuts import render

# Create your views here.
from django.views.decorators.csrf import csrf_exempt
from rest_framework.parsers import JSONParser
from django.http.response import JsonResponse
from UserApp.serializers import StudentSerializer
from UserApp.models import User

@csrf_exempt
def studentApi(request,id=0):
    if request.method=='GET':
        student = User.objects.all()
        student_serializer=StudentSerializer(student,many=True)
        return JsonResponse(student_serializer.data,safe=False)
    elif request.method=='POST':
        student_data=JSONParser().parse(request)
        student_serializer=StudentSerializer(data=student_data)
        if student_serializer.is_valid():
            student_serializer.save()
            return JsonResponse("Added Successfully",safe=False)
        return JsonResponse({"error": "Failed to Add"}, status=400)
    elif request.method=='PUT':
        student_data=JSONParser().parse(request)
        student=User.objects.get(ID=id)
        student_serializer=StudentSerializer(student,data=student_data)
        if student_serializer.is_valid():
            student_serializer.save()
            return JsonResponse("Updated Successfully",safe=False)
        return JsonResponse("Failed to Update")
    elif request.method=='DELETE':
        student=User.objects.get(ID=id)
        student.delete()
        return JsonResponse("Deleted Successfully",safe=False)
