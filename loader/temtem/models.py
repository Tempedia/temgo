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


class Temtem(models.Model):
    no = models.PositiveIntegerField(primary_key=True)
    name = models.CharField(max_length=64)
    type = ArrayField(models.CharField(max_length=64))
    catch_rate = models.FloatField()
    gender_ratio = models.JSONField()
    experience_yield_modifier = models.FloatField()
    icon = models.CharField(max_length=1024)  # 图标
    luma_icon = models.CharField(max_length=1024)  # 闪光图标
    traits = ArrayField(models.CharField(max_length=64))
    description = models.JSONField()
    cry = models.CharField(max_length=1024)  # 叫声
    height = models.FloatField()  # 身高，cm
    weight = models.FloatField()  # 体重, kg
    tv_yield = models.JSONField()
    evolves_to = models.JSONField()  # 进化到
    stats = models.JSONField()
    type_matchup = models.JSONField()
    techniques = models.JSONField()
    trivia = ArrayField(models.TextField())
    gallery = models.JSONField()
    renders = models.JSONField()

    class Meta:
        db_table = "temtem"
