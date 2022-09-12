#!/usr/env python3

import scrapy
from ..items import DownloadImageItem, TemtemItem
from scrapy.exceptions import DropItem


class TypeSpider(scrapy.Spider):
    name = 'temtem'
    start_urls = ['https://temtem.wiki.gg/wiki/Temtem_(creatures)']

    def parse(self, response):
        for page in response.css('.wikitable > tbody > tr > td:nth-child(2) > a'):
            yield response.follow(page, self.parse_temtem)

    def parse_temtem(self, response):
        name = response.css('#firstHeading::text').get().strip()
        no = response.css(
            r'tr.infobox-row th:contains("No.") + td::text').get().strip()
        if not no.startswith('#'):
            raise DropItem()
        no = int(no[1:])

        # 属性
        type = []
        for a in response.css(r'tr.infobox-row th:contains("Type") + td > a'):
            type.append(a.css('::attr(title)').get())

        # 性别比例
        genderRatio = {'male': 0, 'female': 0}
        maleStyle = response.css(
            r'div.gender-ratio > div.male::attr(style)').get()
        femaleStyle = response.css(
            r'div.gender-ratio > div.female::attr(style)').get()
        if maleStyle and femaleStyle:
            t = maleStyle.split(':')[1]
            genderRatio['male'] = int(t.split('%')[0])
            t = femaleStyle.split(':')[1]
            genderRatio['female'] = int(t.split('%')[0])

        catchRate = float(response.css(
            r'tr.infobox-row th:contains("Catch Rate") + td::text').get().strip())

        experienceYieldModifier = float(response.css(
            r'tr.infobox-row th:contains("Experience Yield Modifier") + td::text').get().strip())

        traits = []
        for a in response.css(r'tr.infobox-row th:contains("Traits") + td > a'):
            traits.append(a.css('::text').get().strip())

        normalIconSrc = response.css(
            r'#ttw-temtem > span:nth-child(1) > a:nth-child(1) > img:nth-child(1)::attr(src)').get().strip()
        normalIcon = DownloadImageItem(
            image_url=response.urljoin(normalIconSrc))

        lumaIconSrc = response.css(
            '#ttw-temtem-luma > span:nth-child(1) > a:nth-child(1) > img:nth-child(1)::attr(src)').get().strip()
        lumaIcon = DownloadImageItem(image_url=response.urljoin(lumaIconSrc))

        description = {'Physical Appearance': '', 'Tempedia': ''}
        description['Physical Appearance'] = response.css(
            '.mw-parser-output > h3:contains("Physical Appearance") + p::text').get()
        description['Tempedia'] = response.css(
            '.mw-parser-output > h3:contains("Tempedia") + p >i::text').get()

        yield TemtemItem(
            name=name,
            no=no,
            type=type,
            genderRatio=genderRatio,
            catchRate=catchRate,
            experienceYieldModifier=experienceYieldModifier,
            normalIcon=normalIcon,
            lumaIcon=lumaIcon,
            traits=traits,
            description=description,
        )
