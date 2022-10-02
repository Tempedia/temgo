
import scrapy


class TypeSpider(scrapy.Spider):
    name = 'temtem'
    start_urls = ['https://temtem.wiki.gg/wiki/Status_Conditions']

    def parse(self, response):
        pass
