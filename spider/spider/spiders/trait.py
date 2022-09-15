#!/usr/env python3

import scrapy

from ..items import TraitItem


class TraitSpider(scrapy.Spider):
    name = 'trait'
    start_urls = ['https://temtem.wiki.gg/wiki/Traits']

    def parse(self, response):
        for tr in response.css(r'.mw-parser-output > h2:contains("Available Traits") + table.wikitable tbody tr'):
            name = tr.css(r'td:first-child a::text').get('').strip()
            if not name:
                continue
            desc = tr.css(r'td:nth-child(2)').get()
            impact = tr.css(r'td:nth-child(3)::text').get('').strip()
            trigger = tr.css(r'td:nth-child(4)::text').get('').strip()
            effect = tr.css(r'td:nth-child(5)::text').get('').strip()
            yield TraitItem(
                name=name,
                desc=desc,
                impact=impact,
                trigger=trigger,
                effect=effect,
            )
