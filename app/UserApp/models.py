from django.db import models

# Create your models here.
class User(models.Model):
    ID = models.AutoField(primary_key=True)
    Fname = models.CharField(max_length=255)
    Lname = models.CharField(max_length=255)
    Email = models.EmailField(unique=True, max_length=255)
    Pass = models.CharField(max_length=255)