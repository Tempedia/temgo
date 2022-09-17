from pyexpat import model
from django.db import models
from django.contrib.postgres.fields import ArrayField
# Create your models here.


class Type(models.Model):
    name = models.CharField(primary_key=True, max_length=64, db_index=True)
    icon = models.CharField(max_length=1024)
    comment = models.TextField()
    trivia = ArrayField(models.TextField())
    effective_against = ArrayField(models.CharField(max_length=64))
    ineffective_against = ArrayField(models.CharField(max_length=64))
    resistant_to = ArrayField(models.CharField(max_length=64))
    weak_to = ArrayField(models.CharField(max_length=64))

    sort = models.PositiveSmallIntegerField('sort', db_index=True)

    class Meta:
        db_table = "temtem_type"
