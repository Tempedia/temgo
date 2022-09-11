#!/usr/env python3

import scrapy
from ..items import DownloadImageItem

class TypeSpider(scrapy.Spider):
    name = 'type'
    start_urls = ['https://temtem.wiki.gg/wiki/Temtem_types']

    def parse(self, response):
        for page in response.css('.mw-parser-output > ul:nth-child(5) > li > a:nth-child(1)'):
            yield response.follow(page,self.parse_type)
            # yield {'type-name': li.css('a:nth-child(1)::attr(href)').get()}

    def parse_type(self,response):
        name = response.css('#firstHeading::text').get().strip()
        comment=response.css('.mw-parser-output > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(2) > i:nth-child(1)::text').get().strip()
        commentAuthor=response.css('.mw-parser-output > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(2) > small:nth-child(1) > b:nth-child(1)::text').get().strip()
        iconSrc=response.css('.image > img:nth-child(1)::attr(src)').get()

        icon= DownloadImageItem(icon_url=response.urljoin(iconSrc),name=name,comment=comment,commentAuthor=commentAuthor)
        yield icon
