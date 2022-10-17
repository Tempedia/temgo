from email.policy import default
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
    name = models.CharField(max_length=64, unique=True)
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
    # techniques = models.JSONField()
    trivia = ArrayField(models.TextField())
    gallery = models.JSONField()
    renders = models.JSONField()

    subspecies = models.JSONField(default=list)

    class Meta:
        db_table = "temtem"


class TemtemTrait(models.Model):
    name = models.CharField(max_length=64, primary_key=True)
    description = models.TextField()
    impact = models.CharField(max_length=64)
    trigger = models.CharField(max_length=256)
    effect = models.CharField(max_length=256)

    class Meta:
        db_table = "temtem_trait"


TECHNIQUE_CLASS_CHOICES = (
    ("Physical", "Physical"),
    ("Special", "Special"),
    ("Status", "Status"),
)

TECHNIQUE_TARGETING_CHOICES = (
    ("Self", "Self"),
    ("Single Target", "Single Target"),
    ("Single other Target", "Single other Target"),
    ("Single Team", "Single Team"),
    ("Other Team or Ally", "Other Team or Ally"),
    ("All", "All"),
    ("All Other Temtem", "All Other Temtem"),
)


class TemtemTechnique(models.Model):
    name = models.CharField(max_length=64, primary_key=True)
    type = models.CharField(max_length=64)
    cls = models.CharField(db_column='class', max_length=32,
                           choices=TECHNIQUE_CLASS_CHOICES)
    damage = models.SmallIntegerField()
    sta_cost = models.SmallIntegerField()
    hold = models.SmallIntegerField()
    priority = models.SmallIntegerField()
    targeting = models.CharField(
        max_length=64, choices=TECHNIQUE_TARGETING_CHOICES)
    description = models.TextField(default='')
    video = models.CharField(max_length=1024, default='')

    synergy_description = models.TextField(default='')
    synergy_type = models.CharField(max_length=32)
    synergy_effects = models.TextField(default='')
    synergy_damage = models.SmallIntegerField(default=0)
    synergy_sta_cost = models.SmallIntegerField(default=0)
    synergy_priority = models.SmallIntegerField(default=0)
    synergy_targeting = models.CharField(
        max_length=64, choices=TECHNIQUE_TARGETING_CHOICES, default='')
    synergy_video = models.CharField(max_length=1024, default='')

    class Meta:
        db_table = "temtem_technique"


class TemtemLevelingUpTechnique(models.Model):
    temtem = models.CharField(max_length=64)
    level = models.SmallIntegerField()
    technique_name = models.CharField(max_length=64)
    stab = models.BooleanField(default=False)
    group = models.CharField(max_length=32, default='')

    class Meta:
        db_table = "temtem_leveling_up_technique"


class TemtemCourseTechnique(models.Model):
    temtem = models.CharField(max_length=64)
    course = models.CharField(max_length=32)
    technique_name = models.CharField(max_length=64)
    stab = models.BooleanField(default=False)

    class Meta:
        db_table = "temtem_course_technique"


class TemtemBreedingTechnique(models.Model):
    temtem = models.CharField(max_length=64)
    parents = models.JSONField()
    technique_name = models.CharField(max_length=64)
    stab = models.BooleanField(default=False)

    class Meta:
        db_table = "temtem_breeding_technique"


class TemtemLocation(models.Model):
    name = models.CharField(max_length=128, primary_key=True)
    description = models.TextField()
    island = models.CharField(
        max_length=64, blank=True, null=False, default='')
    image = models.CharField(max_length=1024)
    comment = models.TextField()
    connected_locations = ArrayField(models.CharField(max_length=128))

    class Meta:
        db_table = "temtem_location"


class TemtemLocationArea(models.Model):
    location = models.CharField(max_length=128)
    name = models.CharField(max_length=256)
    image = models.CharField(max_length=1024)
    temtems = models.JSONField()

    class Meta:
        db_table = "temtem_location_area"
        unique_together = ('location', 'name')


class TemtemStatusCondition(models.Model):
    name = models.CharField(max_length=64, primary_key=True)
    icon = models.CharField(max_length=1024)
    description = models.TextField()
    group = models.CharField(max_length=64)
    techniques = ArrayField(models.CharField(max_length=64), default=list)
    traits = ArrayField(models.CharField(max_length=64), default=list)

    class Meta:
        db_table = "temtem_status_condition"


class TemtemCourseItem(models.Model):
    no = models.CharField(max_length=64, primary_key=True)
    technique = models.CharField(max_length=64)
    source = models.TextField()

    class Meta:
        db_table = "temtem_course_item"


class TemtemItemCategory(models.Model):
    name = models.CharField(max_length=64, primary_key=True)
    parent = models.CharField(max_length=64)
    sort = models.SmallIntegerField()

    class Meta:
        db_table = "temtem_item_category"


class TemtemItem(models.Model):
    name = models.CharField(max_length=128, primary_key=True)
    icon = models.CharField(max_length=1024)
    description = models.TextField()
    tradable = models.BooleanField(default=False)
    buy_price = models.TextField()
    sell_price = models.TextField()

    category = models.CharField(max_length=64)
    extra = models.JSONField(default=dict)
    sort = models.SmallIntegerField()

    class Meta:
        db_table = "temtem_item"
