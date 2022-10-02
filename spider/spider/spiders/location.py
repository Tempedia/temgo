#!/usr/env python3

from os import sep
import scrapy

from .utils import parseStrList

from ..items import LocationItem, TemtemImageItem


class LocationSpider(scrapy.Spider):
    name = 'location'
    start_urls = ['https://temtem.wiki.gg/wiki/Temtem_(creatures)']

    def parse(self, response):
        for page in response.css('.wikitable > tbody > tr > td:nth-child(2) > a'):
            yield response.follow(page, self.parse_temtem)

    def parse_temtem(self, response):
        for page in response.css(r'tr.infobox-row th:contains("Locations") + td a'):
            if page.css('::text').get('') == 'Evolution':
                continue
            yield response.follow(page, self.parse_location)

    def parse_location(self, response):
        name = response.css(r'#firstHeading::text').get()
        if not name:
            return
        comment = response.css(r'.infobox + table').get()
        desc = ''
        if not comment:
            desc = parseStrList(response.css(
                r'.mw-parser-output > .infobox ~ p').getall())
        else:
            desc = parseStrList(response.css(
                r'.mw-parser-output > .infobox + table ~ p').getall())
        connectedAreas = []
        for a in response.css(r'tr.infobox-row th:contains("Connected Areas") + td a'):
            connectedAreas.append(a.css('::attr(title)').get())
        island = response.css(
            r'tr.infobox-row th:contains("Island") + td a::attr(title)').get()
        imgSrc = response.css(
            r'.infobox-table > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(1) > a:nth-child(1)::attr(href)').get('')
        imgSrc = response.urljoin(imgSrc)
        imgText = response.css(
            r'.infobox-table > tbody:nth-child(1) > tr:nth-child(6) > td:nth-child(1) > i:nth-child(1)::text').get('')

        areas = []

        for table in response.css(r'.encounterbox article.tabber__panel table.encounterbox-table'):
            title = table.css(r'tr td.areaName b::text').get('')
            if not title:
                continue
            imgSrc = table.css(r'tr td.map a.image::attr(href)').get('')
            if imgSrc:
                imgSrc = response.urljoin(imgSrc)
            temtems = []
            for t in table.css(r'tr td.encounters table.encounterbox-temtem'):
                tname = t.css(r'tr td.temtemName a::attr(title)').get('')
                if not tname:
                    continue
                odds = t.css(r'tr:nth-child(4) td::text').get('').strip()
                levels = t.css(r'tr:nth-child(5) td::text').get('').strip()
                egg = False
                fromm = 0
                to = 0
                if levels == 'Egg':
                    egg = True
                else:
                    seps = levels.split('-')
                    if len(seps) == 1:
                        fromm = int(seps[0])
                        to = fromm
                    else:
                        fromm = int(seps[0])
                        to = int(seps[1])
                temtems.append({
                    'name': tname,
                    'odds': odds,
                    'level': {
                        'from': fromm,
                        'to': to,
                        'egg': egg,
                    }
                })
            area = {
                'title': title,
                'image': imgSrc,
                'temtems': temtems,
            }
            areas.append(area)

        yield LocationItem(
            name=name,
            desc=desc,
            connectedAreas=connectedAreas,
            island=island,
            image=TemtemImageItem(image_url={'url': imgSrc, 'text': imgText}),
            comment=comment,
            areas=areas,
        )
