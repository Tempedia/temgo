# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy
from scrapy.loader import ItemLoader
from scrapy.loader.processors import TakeFirst


class DownloadImageItem(scrapy.Item):
    image_url = scrapy.Field()
    image = scrapy.Field()


class DownloadImagesItem(scrapy.Item):
    image_urls = scrapy.Field()
    images = scrapy.Field()


class TemtemImagesItem(scrapy.Item):
    image_urls = scrapy.Field()


class TemtemImageItem(scrapy.Item):
    image_url = scrapy.Field()


class DownloadFileItem(scrapy.Item):
    file_url = scrapy.Field()
    file = scrapy.Field()


class TypeItem(scrapy.Item):
    name = scrapy.Field()
    icon = scrapy.Field()
    comment = scrapy.Field()
    commentAuthor = scrapy.Field()
    effectiveAgainst = scrapy.Field()
    ineffectiveAgainst = scrapy.Field()
    resistantTo = scrapy.Field()
    weakTo = scrapy.Field()
    trivia = scrapy.Field()


class TemtemItem(scrapy.Item):
    no = scrapy.Field()
    name = scrapy.Field()
    type = scrapy.Field()
    genderRatio = scrapy.Field()
    catchRate = scrapy.Field()
    experienceYieldModifier = scrapy.Field()
    normalIcon = scrapy.Field()
    lumaIcon = scrapy.Field()
    traits = scrapy.Field()
    description = scrapy.Field()
    cry = scrapy.Field()
    locations = scrapy.Field()
    height = scrapy.Field()
    weight = scrapy.Field()
    tvYield = scrapy.Field()
    evolvesTo = scrapy.Field()
    stats = scrapy.Field()
    typeMatchup = scrapy.Field()
    techniques = scrapy.Field()
    trivia = scrapy.Field()
    gallery = scrapy.Field()
    renders = scrapy.Field()


class TechniqueItem(scrapy.Item):
    name = scrapy.Field()
    type = scrapy.Field()
    synergyType = scrapy.Field()
    clas = scrapy.Field()
    dmg = scrapy.Field()
    sta = scrapy.Field()
    hold = scrapy.Field()
    priority = scrapy.Field()
    targeting = scrapy.Field()
    video = scrapy.Field()

    effect = scrapy.Field()
    synergyEffect = scrapy.Field()
    desc = scrapy.Field()


class TraitItem(scrapy.Item):
    name = scrapy.Field()
    desc = scrapy.Field()
    impact = scrapy.Field()
    trigger = scrapy.Field()
    effect = scrapy.Field()


class LocationItem(scrapy.Item):
    name = scrapy.Field()
    connectedAreas = scrapy.Field()
    island = scrapy.Field()
    image = scrapy.Field()
    desc = scrapy.Field()
    comment = scrapy.Field()
    areas = scrapy.Field()
