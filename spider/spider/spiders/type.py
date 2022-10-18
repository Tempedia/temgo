#!/usr/env python3

from curses import meta
import scrapy
from ..items import DownloadImageItem, TypeItem


class TypeSpider(scrapy.Spider):
    name = 'type'
    start_urls = ['https://temtem.wiki.gg/wiki/Temtem_types']

    def parse(self, response):
        i = 0
        for li in response.css(r'.mw-parser-output > h2:contains("Types") + p + ul > li'):
            page = li.css(r':scope > a:nth-child(1)')[0]
            color = li.css(r':scope > ul li b::text').get('')
            yield response.follow(page, self.parse_type, meta={'sort': i, 'color': color})
            i += 1

    def parse_type(self, response):
        name = response.css('#firstHeading::text').get().strip()
        comment = response.css(
            '.mw-parser-output > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(2) > i:nth-child(1)::text').get().strip()
        commentAuthor = response.css(
            '.mw-parser-output > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(2) > small:nth-child(1) > b:nth-child(1)::text').get().strip()
        iconSrc = response.css('.image > img:nth-child(1)::attr(src)').get()

        # 属性相克
        effectiveAgainst = []
        for a in response.css('tr.infobox-row:nth-child(4) > td:nth-child(2) > a'):
            effectiveAgainst.append(a.css('::attr(title)').get())
        ineffectiveAgainst = []
        for a in response.css('tr.infobox-row:nth-child(5) > td:nth-child(2) > a'):
            ineffectiveAgainst.append(a.css('::attr(title)').get())
        resistantTo = []
        for a in response.css('tr.infobox-row:nth-child(6) > td:nth-child(2) > a'):
            resistantTo.append(a.css('::attr(title)').get())
        weakTo = []
        for a in response.css('tr.infobox-row:nth-child(7) > td:nth-child(2) > a'):
            weakTo.append(a.css('::attr(title)').get())
        trivia = []
        for li in response.css('.mw-parser-output > ul:nth-child(15) > li'):
            trivia.append(li.get())

        icon = DownloadImageItem(image_url=response.urljoin(iconSrc))

        yield TypeItem(
            icon=icon,
            name=name,
            comment=comment,
            commentAuthor=commentAuthor,
            effectiveAgainst=effectiveAgainst,
            ineffectiveAgainst=ineffectiveAgainst,
            resistantTo=resistantTo,
            weakTo=weakTo,
            trivia=trivia,
            sort=response.meta['sort'],
            color=response.meta['color'],
        )
