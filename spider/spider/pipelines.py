# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
from itemadapter import ItemAdapter
import scrapy
from scrapy.pipelines.images import ImagesPipeline
from scrapy.pipelines.files import FilesPipeline
from scrapy.exceptions import DropItem

from .items import DownloadFileItem, DownloadImageItem
from . import settings
import os


class SpiderPipeline:
    def process_item(self, item, spider):
        return item


class MyImagesPipeline(ImagesPipeline):

    def get_media_requests(self, item, info):
        # if 'image_url' in item and item['image_url']:
        #     yield scrapy.Request(item['image_url'])
        for k, v in item.items():
            if isinstance(v, DownloadImageItem):
                yield scrapy.Request(v['image_url'])

    def item_completed(self, results, item, info):
        for ok, x in results:
            if ok:
                x['path'] = os.path.join(settings.IMAGES_STORE, x['path'])
                for k, v in item.items():
                    if isinstance(v, DownloadImageItem) and v['image_url'] == x['url']:
                        item[k] = x
                # item['image'] = os.path.join(settings.IMAGES_STORE, x['path'])
                # break
        return item


class MyFilesPipeline(FilesPipeline):
    def get_media_requests(self, item, info):
        for k, v in item.items():
            if isinstance(v, DownloadFileItem):
                yield scrapy.Request(v['file_url'])

    def item_completed(self, results, item, info):
        for ok, x in results:
            if ok:
                x['path'] = os.path.join(settings.FILES_STORE, x['path'])
                for k, v in item.items():
                    if isinstance(v, DownloadFileItem) and v['file_url'] == x['url']:
                        item[k] = x
        return item