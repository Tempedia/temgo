#!/usr/env python3

import scrapy
from ..items import DownloadFileItem, DownloadImageItem, TemtemItem
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
            r'#ttw-temtem-luma > span:nth-child(1) > a:nth-child(1) > img:nth-child(1)::attr(src)').get().strip()
        lumaIcon = DownloadImageItem(image_url=response.urljoin(lumaIconSrc))

        description = {'Physical Appearance': '', 'Tempedia': ''}
        description['Physical Appearance'] = response.css(
            r'.mw-parser-output > h3:contains("Physical Appearance") + p::text').get()
        description['Tempedia'] = response.css(
            r'.mw-parser-output > h3:contains("Tempedia") + p >i::text').get()
        
        crySrc=response.css(r'tr.infobox-row th:contains("Cry") + td audio::attr(src)').get()
        cry=DownloadFileItem(file_url=crySrc)

        locations=[]
        for a in response.css(r'tr.infobox-row th:contains("Locations") + td a'):
            locations.append(a.css('::text').get())
        
        height=response.css(r'table.infobox-half-row tr:contains("Height") + tr td::text').get()
        weight=response.css(r'table.infobox-half-row tr:contains("Weight") + tr td::text').get()

        tvYield={}
        tvYield['HP']=response.css(r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(1)::text').get('').strip()
        tvYield['STA']=response.css(r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(2)::text').get('').strip()
        tvYield['SPD']=response.css(r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(3)::text').get('').strip()
        tvYield['ATK']=response.css(r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(4)::text').get('').strip()
        tvYield['DEF']=response.css(r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(5)::text').get('').strip()
        tvYield['SPATK']=response.css(r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(6)::text').get('').strip()
        tvYield['SPDEF']=response.css(r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(7)::text').get('').strip()


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
            cry=cry,
            locations=locations,
            height=height,
            weight=weight,
            tvYield=tvYield,
        )
