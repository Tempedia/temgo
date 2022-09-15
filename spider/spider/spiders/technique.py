#!/usr/env python3

import scrapy

from ..items import DownloadFileItem, TechniqueItem
from .utils import parseStrList, parseInt


class TypeSpider(scrapy.Spider):
    name = 'technique'
    start_urls = ['https://temtem.wiki.gg/wiki/Techniques']

    def parse(self, response):
        for tr in response.css(r'.mw-parser-output > h3:contains("Techniques without Synergy") + table.wikitable tbody tr'):
            a = tr.css(r'td:first-child a')
            if not a:
                continue
            name = a.css(r'::text').get()
            type = tr.css(r'td:nth-child(2) a::attr(title)').get('').strip()
            clas = tr.css(r'td:nth-child(3) a::attr(title)').get('').strip()
            dmg = parseInt(tr.css(r'td:nth-child(4)::text').get('').strip())
            sta = parseInt(tr.css(r'td:nth-child(5)::text').get('').strip())
            hold = parseInt(tr.css(r'td:nth-child(6)::text').get('').strip())
            priority = parseInt(
                tr.css(r'td:nth-child(7) a::attr(title)').get('').strip().split('_')[0])
            targeting = tr.css(r'td:nth-child(8)::text').get('').strip()
            item = TechniqueItem(
                name=name,
                type=type,
                clas=clas,
                dmg=dmg,
                sta=sta,
                hold=hold,
                priority=priority,
                targeting=targeting,
            )
            yield response.follow(a[0], self.parse_technique, meta={'item': item})
        # 协同
        for tr in response.css(r'.mw-parser-output > h3:contains("Techniques with Synergy") + table.wikitable tbody tr'):
            a = tr.css(r'td:first-child a')
            if not a:
                continue
            name = a.css(r'::text').get()
            type = tr.css(
                r'td:nth-child(2) a:first-child::attr(title)').get('').strip()
            synergyType = tr.css(
                r'td:nth-child(2) a:nth-child(2)::attr(title)').get('').strip()
            clas = tr.css(r'td:nth-child(3) a::attr(title)').get('').strip()
            dmg = parseInt(tr.css(r'td:nth-child(4)::text').get('').strip())
            sta = parseInt(tr.css(r'td:nth-child(5)::text').get('').strip())
            hold = parseInt(tr.css(r'td:nth-child(6)::text').get('').strip())
            priority = parseInt(
                tr.css(r'td:nth-child(7) a::attr(title)').get('').strip().split('_')[0])
            targeting = tr.css(r'td:nth-child(8)::text').get('').strip()
            item = TechniqueItem(
                name=name,
                type=type,
                clas=clas,
                dmg=dmg,
                sta=sta,
                hold=hold,
                priority=priority,
                targeting=targeting,
                synergyType=synergyType,
            )
            yield response.follow(a[0], self.parse_technique, meta={'item': item})

    def parse_technique(self, response):
        item = response.meta['item']

        i = response.css(
            r'.infobox-table > tbody > tr:nth-child(6) > td:nth-child(1) > i')
        item['desc'] = i.get()

        item['effect'] = response.css(
            r'.mw-parser-output > h2:contains("Effect") + p').get()
        item['synergyEffect'] = response.css(
            r'.mw-parser-output > h2:contains("Effect") + p + p').get()

        videoSrc = response.css(
            r'.infobox-table > tbody > tr > td > video::attr(src)').get('')
        if videoSrc:
            item['video'] = DownloadFileItem(
                file_url=response.urljoin(videoSrc))

        yield item
