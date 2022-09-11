# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
from itemadapter import ItemAdapter
import scrapy
from scrapy.pipelines.images import ImagesPipeline
from scrapy.exceptions import DropItem
from . import settings
import os

class SpiderPipeline:
    def process_item(self, item, spider):
        return item


class MyImagesPipeline(ImagesPipeline):

    def get_media_requests(self, item, info):
        if 'icon_url' in item and item['icon_url']:
            yield scrapy.Request(item['icon_url'])

    def item_completed(self, results, item, info):
        for ok,x in results:
            if ok:
                item['icon']=os.path.join(settings.IMAGES_STORE, x['path'])
                break
        return item
