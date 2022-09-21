#!/usr/env python3

import scrapy

from ..items import DownloadFileItem, TechniqueItem
from .utils import parseStrList, parseInt


class TechniqueSpider(scrapy.Spider):
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

        # item['effect'] = response.css(
        #     r'.mw-parser-output > h2:contains("Effect") + p').get()
        # item['synergyEffect'] = response.css(
        #     r'.mw-parser-output > h2:contains("Effect") + p + p').get()

        videoSrc = response.css(
            r'.infobox-table > tbody > tr:nth-child(3) > td:nth-child(1) > video::attr(src)').get('')
        if videoSrc:
            item['video'] = DownloadFileItem(
                file_url=response.urljoin(videoSrc),
            )

        if 'synergyType' in item and item['synergyType']:
            synergy = response.css(
                r'.infobox-table > tbody> tr:contains("Synergy Details") ~ *')
            item['synergyVideo'] = DownloadFileItem(
                file_url=synergy.css('video::attr(src)').get(''),
            )

            priority = synergy.css(
                'th:contains("Priority") + td a::attr(title)').get('')
            if priority:
                item['synergyPriority'] = int(priority.split('_')[0])
            item['synergyDamage'] = int(synergy.css(
                r'th:contains("Damage") + td::text').get('-1'))
            item['synergyEffects'] = synergy.css(
                r'th:contains("Effects") + td').get('')
            item['synergySta'] = int(synergy.css(
                r'th:contains("STA Cost") + td::text').get('-1'))
            item['synergyTargeting'] = synergy.css(
                r'th:contains("Targeting") + td a::text').get('').strip()
            item['synergyDesc'] = synergy.css(
                r'tr td.infobox-centered[colspan="2"] i').get('')

        yield item
