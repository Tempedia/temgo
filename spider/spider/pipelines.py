# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
from scrapy.utils.defer import maybe_deferred_to_future
from itemadapter import ItemAdapter
import scrapy
# from scrapy.pipelines.images import ImagesPipeline
from scrapy.pipelines.files import FilesPipeline
from scrapy.exceptions import DropItem

from .items import DownloadFileItem, DownloadImageItem, DownloadImagesItem, TemtemImageItem, TemtemImagesItem
from . import settings
import os
import requests
from urllib.parse import urljoin


class SpiderPipeline:
    def process_item(self, item, spider):
        return item


class MyImagePipeline(FilesPipeline):

    def get_media_requests(self, item, info):
        for k, v in item.items():
            if isinstance(v, DownloadImageItem):
                yield scrapy.Request(v['image_url'])
            elif isinstance(v, DownloadImagesItem):
                for url in v['image_urls']:
                    yield scrapy.Request(url)

    def item_completed(self, results, item, info):
        for ok, x in results:
            if ok:
                x['path'] = os.path.join(settings.FILES_STORE, x['path'])
                for k, v in item.items():
                    if isinstance(v, DownloadImageItem) and v['image_url'] == x['url']:
                        item[k] = x
        return item


class MyImagesPipeline(FilesPipeline):

    def get_media_requests(self, item, info):
        for k, v in item.items():
            if isinstance(v, DownloadImagesItem):
                for url in v['image_urls']:
                    yield scrapy.Request(url)

    def item_completed(self, results, item, info):
        m = {}
        for k, v in item.items():
            if isinstance(v, DownloadImagesItem):
                m[k] = []
        for ok, x in results:
            if ok:
                x['path'] = os.path.join(settings.FILES_STORE, x['path'])
                for k, v in item.items():
                    if isinstance(v, DownloadImagesItem) and x['url'] in v['image_urls']:
                        m[k].append(x)
        for k, v in m.items():
            item[k] = v
        return item


class TemtemImagesPipeline1(object):
    async def process_item(self, item, spider):
        for k, v in item.items():
            if isinstance(v, TemtemImagesItem):
                for url in v['image_urls']:
                    request = scrapy.Request(url['url'])
                    response = await maybe_deferred_to_future(spider.crawler.engine.download(request, spider))
                    src = response.css(r'#file > a::attr(href)').get()
                    url['url'] = response.urljoin(src)
        return item


class TemtemImagesPipeline2(FilesPipeline):

    def get_media_requests(self, item, info):
        for k, v in item.items():
            if isinstance(v, TemtemImagesItem):
                for url in v['image_urls']:
                    yield scrapy.Request(url['url'])

    def item_completed(self, results, item, info):
        m = {}
        for k, v in item.items():
            if isinstance(v, TemtemImagesItem):
                m[k] = []
        for ok, x in results:
            if ok:
                x['path'] = os.path.join(settings.FILES_STORE, x['path'])
                for k, v in item.items():
                    if isinstance(v, TemtemImagesItem):
                        for url in v['image_urls']:
                            if x['url'] == url['url']:
                                if 'text' in url:
                                    x['text'] = url['text']
                                if 'group' in url:
                                    x['group'] = url['group']
                                m[k].append(x)
        for k, v in m.items():
            item[k] = v
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


class TemtemImagePipeline1(object):
    async def process_item(self, item, spider):
        for k, v in item.items():
            if isinstance(v, TemtemImageItem):
                url = v['image_url']
                request = scrapy.Request(url['url'])
                response = await maybe_deferred_to_future(spider.crawler.engine.download(request, spider))
                src = response.css(r'#file > a::attr(href)').get()
                url['url'] = response.urljoin(src)
        return item


class TemtemImagePipeline2(FilesPipeline):

    def get_media_requests(self, item, info):
        for k, v in item.items():
            if isinstance(v, TemtemImageItem):
                url = v['image_url']
                yield scrapy.Request(url['url'])

    def item_completed(self, results, item, info):
        m = {}
        for k, v in item.items():
            if isinstance(v, TemtemImageItem):
                m[k] = None
        for ok, x in results:
            if ok:
                x['path'] = os.path.join(settings.FILES_STORE, x['path'])
                for k, v in item.items():
                    if isinstance(v, TemtemImageItem):
                        url = v['image_url']
                        if x['url'] == url['url']:
                            if 'text' in url:
                                x['text'] = url['text']
                            m[k] = x
                            break
        for k, v in m.items():
            item[k] = v
        return item
