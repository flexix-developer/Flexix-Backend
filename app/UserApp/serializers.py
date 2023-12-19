from rest_framework import serializers
from UserApp.models import User, alet
from rest_framework import serializers

class StudentSerializer(serializers.ModelSerializer):
    class Meta:
        # model = User
        model = alet
        # fields = '__all__'
        fields = ('id', 'fname', 'lname', 'email', 'passw')

        

class LoginSerializer(serializers.Serializer):
    email = serializers.EmailField()
    password = serializers.CharField(write_only=True)
    