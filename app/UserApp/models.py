from django.db import models
from django.contrib.auth.models import AbstractUser

# Create your models here.
class User(models.Model):
    ID = models.AutoField(primary_key=True)
    Fname = models.CharField(max_length=255)
    Lname = models.CharField(max_length=255)
    Email = models.EmailField(unique=True, max_length=255)
    Pass = models.CharField(max_length=255)
    
class alet(AbstractUser):
    id = models.AutoField(primary_key=True)
    fname = models.CharField(max_length=255)
    lname = models.CharField(max_length=255)
    email = models.EmailField(unique=True, max_length=255)
    passw = models.CharField(max_length=255)

    # กำหนด related_name ใน ManyToManyField ที่เป็น groups
    groups = models.ManyToManyField('auth.Group', related_name='alet_groups')
    
    # กำหนด related_name ใน ManyToManyField ที่เป็น user_permissions
    user_permissions = models.ManyToManyField('auth.Permission', related_name='alet_user_permissions')