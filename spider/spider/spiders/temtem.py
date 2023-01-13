#!/usr/env python3

import scrapy

from .utils import parseStrList
from ..items import DownloadFileItem, DownloadImageItem, DownloadImagesItem, TemtemImagesItem, TemtemItem
from scrapy.exceptions import DropItem


class TypeSpider(scrapy.Spider):
    name = 'temtem'
    start_urls = ['https://temtem.wiki.gg/wiki/Temtem_(creatures)']

    def parse(self, response):
        for page in response.css('.wikitable > tbody > tr > td:nth-child(2) > a'):
            yield response.follow(page, self.parse_temtem)

    def parse_temtem(self, response):
        name = response.css('#firstHeading span::text').get().strip()
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
            r'tr.infobox-row th:contains("Exp Yield Modifier") + td::text').get().strip())

        traits = []
        for a in response.css(r'tr.infobox-row th:contains("Traits") + td > a'):
            traits.append(a.css('::text').get().strip())

        normalIconSrc = response.css(
            r'article.tabber__panel[data-title=Normal] img::attr(src)').get().strip()
        normalIcon = DownloadImageItem(
            image_url=response.urljoin(normalIconSrc))

        lumaIconSrc = response.css(
            r'article.tabber__panel[data-title=Luma] img::attr(src)').get().strip()
        lumaIcon = DownloadImageItem(image_url=response.urljoin(lumaIconSrc))

        description = {'Physical Appearance': '', 'Tempedia': ''}
        description['Physical Appearance'] = parseStrList(response.css(
            r'.mw-parser-output > h3:contains("Physical Appearance") ~ p').getall())
        appearance = []
        x = response.css(
            r'.mw-parser-output > h3:contains("Physical Appearance")')
        while True:
            x = x.css(r':scope + *')
            if x.xpath('name()').get() in ('h3', 'h2'):
                break
            appearance.append(x.get())
        description['Physical Appearance'] = parseStrList(appearance)
        description['Tempedia'] = response.css(
            r'.mw-parser-output > h3:contains("Tempedia") + p >i::text').get()

        crySrc = response.css(
            r'tr.infobox-row th:contains("Cry") + td audio::attr(src)').get()
        cry = DownloadFileItem(file_url=crySrc)

        locations = []
        for a in response.css(r'tr.infobox-row th:contains("Locations") + td a'):
            location = a.css('::text').get('')
            if not location or location == 'Evolution':
                continue
            locations.append(location)

        height = response.css(
            r'table.infobox-half-row tr:contains("Height") + tr td::text').get()
        weight = response.css(
            r'table.infobox-half-row tr:contains("Weight") + tr td::text').get()

        tvYield = {}
        tvYield['HP'] = response.css(
            r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(1)::text').get('').strip()
        tvYield['STA'] = response.css(
            r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(2)::text').get('').strip()
        tvYield['SPD'] = response.css(
            r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(3)::text').get('').strip()
        tvYield['ATK'] = response.css(
            r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(4)::text').get('').strip()
        tvYield['DEF'] = response.css(
            r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(5)::text').get('').strip()
        tvYield['SPATK'] = response.css(
            r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(6)::text').get('').strip()
        tvYield['SPDEF'] = response.css(
            r'.tv-table > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(7)::text').get('').strip()

        # 进化链
        evolvesTo = []
        for div in response.css(
                r'div.evobox-container > table.evobox:contains("'+name+r'") + div.evobox-evolution'):
            method = div.css(r'div.evolution-description::text').get('') + ' ' +\
                ' '.join(div.css(r'div.evolution-description a::text').getall())
            method = method.strip()
            totable = div.css(r':scope + table.evobox')
            if not totable:
                continue
            to = totable.css('tr.evobox-name > td > a::text').get('')
            evolvesTo.append({
                'method': method,
                'to': to,
            })

        # 基本属性
        stats = {
            'HP': {
                'base': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(3) > th:nth-child(1) > div:nth-child(2)::text').get(),
                '50': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(3) > small:nth-child(1) > b:nth-child(1)::text').get(),
                '100': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(4) > small:nth-child(1) > b:nth-child(1)::text').get(),
            },
            'STA': {
                'base': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(4) > th:nth-child(1) > div:nth-child(2)::text').get(),
                '50': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(4) > td:nth-child(3) > small:nth-child(1) > b:nth-child(1)::text').get(),
                '100': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(4) > td:nth-child(4) > small:nth-child(1) > b:nth-child(1)::text').get(),
            },
            'SPD': {
                'base': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(5) > th:nth-child(1) > div:nth-child(2)::text').get(),
                '50': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(5) > td:nth-child(3) > small:nth-child(1) > b:nth-child(1)::text').get(),
                '100': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(5) > td:nth-child(4) > small:nth-child(1) > b:nth-child(1)::text').get(),
            },
            'ATK': {
                'base': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(6) > th:nth-child(1) > div:nth-child(2)::text').get(),
                '50': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(6) > td:nth-child(3) > small:nth-child(1) > b:nth-child(1)::text').get(),
                '100': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(6) > td:nth-child(4) > small:nth-child(1) > b:nth-child(1)::text').get(),
            },
            'DEF': {
                'base': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(7) > th:nth-child(1) > div:nth-child(2)::text').get(),
                '50': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(7) > td:nth-child(3) > small:nth-child(1) > b:nth-child(1)::text').get(),
                '100': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(7) > td:nth-child(4) > small:nth-child(1) > b:nth-child(1)::text').get(),
            },
            'SPATK': {
                'base': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(8) > th:nth-child(1) > div:nth-child(2)::text').get(),
                '50': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(8) > td:nth-child(3) > small:nth-child(1) > b:nth-child(1)::text').get(),
                '100': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(8) > td:nth-child(4) > small:nth-child(1) > b:nth-child(1)::text').get(),
            },
            'SPDEF': {
                'base': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(9) > th:nth-child(1) > div:nth-child(2)::text').get(),
                '50': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(9) > td:nth-child(3) > small:nth-child(1) > b:nth-child(1)::text').get(),
                '100': response.css(r'.statbox > tbody:nth-child(1) > tr:nth-child(9) > td:nth-child(4) > small:nth-child(1) > b:nth-child(1)::text').get(),
            },
        }

        # 属性抵抗
        def extractMatchup(response, group):
            matchup = []
            tmap = {}
            title = ''
            for i, th in enumerate(response.css(
                    r'.type-table > tbody:nth-child(1) > tr:nth-child(1) > th')):
                t = th.css('a::attr(title)').get('').strip()
                if 'type' in t or 'Type' in t:
                    tmap[i] = t
                else:
                    title = th.css('::text').get('').strip()
            for tr in response.css(
                    r'.type-table > tbody:nth-child(1) > tr:not(:nth-child(1)):not(:last-child)'):
                m = {}
                for i, td in enumerate(tr.css('td')):
                    tdclass = td.css('::attr(class)').get('').strip()
                    if i in tmap and tdclass.startswith('resist'):
                        v = 1
                        if tdclass == 'resist--0':
                            v = 0
                        elif tdclass == 'resist--2':
                            v = 2
                        elif tdclass == 'resist--05':
                            v = 0.5
                        elif tdclass == 'resist--4':
                            v = 4
                        elif tdclass == 'resist--025':
                            v = 0.25
                        m[tmap[i]] = v
                    else:
                        m['name'] = td.css('::text').get().strip()
                        m['title'] = title
                    if group:
                        m['group'] = group
                matchup.append(m)
            return matchup

        typeMatchup = []
        tabs = response.css(
            r'.mw-parser-output > h2:contains("Type Matchup") + div.koish-tabs,.mw-parser-output > h2:contains("Type Matchup") + p + div.koish-tabs')
        if not tabs:  # 没有分组
            typeMatchup = extractMatchup(response, None)
        else:   # 分组
            for article in tabs.css('div.tabber article'):
                group = article.css('::attr(title)').get('').strip()
                typeMatchup.extend(extractMatchup(article, group))

        # 技能
        techniques = {
            'leveling_up': [],
            'course': [],
            'breeding': [],
        }
        # 升级技能

        def extractLevelingUpTechniques(response, group=None):
            rt = []
            for tr in response.css(r'tbody tr:not(:last-child)'):
                if not tr.css('td:nth-child(2)'):
                    continue
                levelStr = tr.css(r'td:first-child::text').get().strip()
                level = int(levelStr) if levelStr != '?' else 0
                stab = bool(tr.css(r'td:nth-child(2) > i'))
                technique = tr.css(r'td:nth-child(2) a::text').get('').strip()
                rt.append({
                    'level': level,
                    'stab': stab,
                    'technique': technique,
                    'group': group,
                })
            return rt
        table = response.css(
            r'.mw-parser-output > h3:contains("By Leveling up") + table.learnlist')
        if table:
            techniques['leveling_up'] = extractLevelingUpTechniques(table)
        else:
            tabs = response.css(
                r'.mw-parser-output > h3:contains("By Leveling up") + div.koish-tabs')
            if not tabs:
                table = response.css(
                    r'.mw-parser-output > h2:contains("Techniques") + table.learnlist')
                if not table:
                    raise Exception('no Leveling Up Techniques')
                techniques['leveling_up'] = extractLevelingUpTechniques(table)
            else:
                for tab in tabs.css(r'div.tabber article'):
                    group = tab.css(r'::attr(title)').get('').strip()
                    table = tab.css(r'table.learnlist')
                    rt = extractLevelingUpTechniques(table, group)
                    techniques['leveling_up'].extend(rt)
        # for tr in response.css(
        #         r'.mw-parser-output > h3:contains("By Leveling up") + table.learnlist tbody tr:not(:last-child)'):
        #     if not tr.css('td:nth-child(2)'):
        #         continue
        #     levelStr = tr.css(r'td:first-child::text').get().strip()
        #     level = int(levelStr) if levelStr != '?' else 0
        #     stab = bool(tr.css(r'td:nth-child(2) > i'))
        #     technique = tr.css(r'td:nth-child(2) a::text').get('').strip()
        #     techniques['leveling_up'].append({
        #         'level': level,
        #         'stab': stab,
        #         'technique': technique,
        #     })
        # 教程技能
        for tr in response.css(
                r'.mw-parser-output > h3:contains("By Technique Course") + table.learnlist tbody tr:not(:last-child)'):
            if not tr.css('td:nth-child(2)'):
                continue
            course = tr.css(r'td:first-child a::text').get('').strip()
            stab = bool(tr.css(r'td:nth-child(2) > i'))
            technique = tr.css(r'td:nth-child(2) a::text').get('').strip()
            techniques['course'].append({
                'course': course,
                'stab': stab,
                'technique': technique,
            })
        # 遗传技能
        breeding = response.css(
            r'.mw-parser-output > h3:contains("By Breeding")')
        if not breeding:
            breeding = response.css(
                r'.mw-parser-output > h3:contains("By breeding")')
        for tr in breeding.css(r':scope + table.learnlist tbody tr:not(:last-child)'):
            if not tr.css('td:nth-child(2)'):
                continue
            stab = bool(tr.css(r'td:nth-child(2) > i'))
            technique = tr.css(r'td:nth-child(2) a::text').get('').strip()
            parents = []
            for pdiv in tr.css(r'td:first-child div.parents div.parent'):
                hint = pdiv.css('::text').get('').strip()
                parent = pdiv.css('a::attr(title)').get()
                parents.append({
                    'name': parent,
                    'hint': hint,
                })
            # parents = tr.css(
            #     'td:first-child div.parents div.parent a::attr(title)').getall()
            techniques['breeding'].append({
                'parents': parents,
                'stab': stab,
                'technique': technique,
            })
        # 冷知识
        trivia = []
        for li in response.css(r'.mw-parser-output > h2:contains("Trivia") + ul li '):
            trivia.append(li.get())

        # 图库
        gallery = TemtemImagesItem(image_urls=[])
        for li in response.css(r'.mw-parser-output > h2:contains("Gallery") + ul.gallery li '):
            src = li.css(r'div.thumb a.image::attr(href)').get()
            if not src:
                continue
            gallery['image_urls'].append({
                'url': response.urljoin(src),
                'text': li.css(r'div.gallerytext p').get(),
            })

        # Renders
        renders = TemtemImagesItem(image_urls=[])
        ul = response.css(
            r'.mw-parser-output > h3:contains("Renders") + ul.gallery')
        if ul:
            for li in ul.css('li'):
                src = li.css(r'div.thumb a.image::attr(href)').get()
                if not src:
                    continue
                renders['image_urls'].append({
                    'url': response.urljoin(src),
                    'text': li.css(r'div.gallerytext p').get(),
                })
        else:
            tabs = response.css(
                r'.mw-parser-output > h3:contains("Renders") + div.koish-tabs')
            if tabs:
                for tab in tabs.css(r'div.tabber article'):
                    group = tab.css(r'::attr(title)').get('').strip()
                    for li in tab.css(r'ul.gallery li'):
                        src = li.css(r'div.thumb a.image::attr(href)').get()
                        if not src:
                            continue
                        renders['image_urls'].append({
                            'url': response.urljoin(src),
                            'text': li.css(r'div.gallerytext p').get(),
                            'group': group,
                        })

        subspecies = TemtemImagesItem(image_urls=[])
        tabs = response.css(
            r'.mw-parser-output > h3:contains("Subspecies Variations") + p + div.koish-tabs')
        if tabs:
            image_urls = []
            for tab in tabs.css(r'div.tabber article'):
                group = tab.css(r'::attr(title)').get('').strip()
                url = tab.css(
                    r'table.wikitable tbody tr:nth-child(2) td:nth-child(1) a::attr(href)').get('')
                text = 'normal'
                image_urls.append({
                    'url': response.urljoin(url),
                    'text': text,
                    'group': group,
                })
                url = tab.css(
                    r'table.wikitable tbody tr:nth-child(2) td:nth-child(2) a::attr(href)').get('')
                text = 'luma'
                image_urls.append({
                    'url': response.urljoin(url),
                    'text': text,
                    'group': group,
                })
            subspecies = TemtemImagesItem(image_urls=image_urls)

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
            evolvesTo=evolvesTo,
            stats=stats,
            typeMatchup=typeMatchup,
            techniques=techniques,
            trivia=trivia,
            gallery=gallery,
            renders=renders,
            subspecies=subspecies,
        )
