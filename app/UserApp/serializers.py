from rest_framework import serializers
from UserApp.models import User

class StudentSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = '__all__'