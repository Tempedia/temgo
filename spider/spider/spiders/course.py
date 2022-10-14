#!/usr/env python3

import scrapy

from .utils import parseStrList
from ..items import TechniqueCourseItem


class TechniqueCourseSpider(scrapy.Spider):
    name = 'technique-course'
    start_urls = ['https://temtem.wiki.gg/wiki/Items']

    def parse(self, response):
        for tr in response.css(
                r'.mw-parser-output > h3:contains("Technique Courses") + p + table tbody tr'):
            t = tr.css(r'td:nth-child(1) a:nth-child(2)::attr(title)').get()
            if not t:
                continue
            names = t.split(':')
            no = names[0].strip()
            technique = names[1].strip()
            source = tr.css(r'td:nth-child(2)').get()
            yield TechniqueCourseItem(
                no=no,
                technique=technique,
                source=source,
            )
