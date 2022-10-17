
import scrapy

from ..items import ItemItem, TemtemImageItem


class ItemSpider(scrapy.Spider):
    name = 'item'
    start_urls = ['https://temtem.wiki.gg/wiki/Items']

    sort = 0

    def parseTable(self, response, table, category, subcategory, tradable):
        items = []
        for tr in table.css(r'tbody tr'):
            td = tr.css(r'td:nth-child(1)')
            if not td:
                continue
            iconsrc = td.css(r'a:nth-child(1)::attr(href)').get('')
            name = td.css(r'a:nth-child(2)::attr(title)').get('')
            if not name or not iconsrc:
                continue
            icon = TemtemImageItem(
                image_url={'url': response.urljoin(iconsrc)},
            )
            description = tr.css(r'td:nth-child(2)').get()
            buyPrice = tr.css(r'td:nth-child(3)').get()
            sellPrice = tr.css(r'td:nth-child(4)').get()

            self.sort += 1
            items.append(ItemItem(
                icon=icon,
                name=name,
                description=description,
                buyPrice=buyPrice,
                sellPrice=sellPrice,
                category=category,
                subcategory=subcategory,
                tradable=tradable,
                sort=self.sort,
            ))
        return items

    def parseTableNoBuyPrice(self, response, table, category, subcategory, tradable):
        items = []
        for tr in table.css(r'tbody tr'):
            td = tr.css(r'td:nth-child(1)')
            if not td:
                continue
            iconsrc = td.css(r'a:nth-child(1)::attr(href)').get('')
            name = td.css(r'a:nth-child(2)::attr(title)').get('')
            if not name or not iconsrc:
                continue
            icon = TemtemImageItem(
                image_url={'url': response.urljoin(iconsrc)},
            )
            description = tr.css(r'td:nth-child(2)').get()
            # buyPrice = tr.css(r'td:nth-child(3)').get()
            sellPrice = tr.css(r'td:nth-child(3)').get()
            self.sort += 1

            items.append(ItemItem(
                icon=icon,
                name=name,
                description=description,
                buyPrice='<td></td>',
                sellPrice=sellPrice,
                category=category,
                subcategory=subcategory,
                tradable=tradable,
                sort=self.sort,
            ))
        return items

    def parseTableCapture(self, response, table, category, subcategory, tradable):
        items = []
        for tr in table.css(r'tbody tr'):
            td = tr.css(r'td:nth-child(1)')
            if not td:
                continue
            iconsrc = td.css(r'a:nth-child(1)::attr(href)').get('')
            name = td.css(r'a:nth-child(2)::attr(title)').get('')
            if not name or not iconsrc:
                continue
            icon = TemtemImageItem(
                image_url={'url': response.urljoin(iconsrc)},
            )
            description = tr.css(r'td:nth-child(2)').get()
            bonus = tr.css(r'td:nth-child(3)').get()
            buyPrice = tr.css(r'td:nth-child(4)').get()
            sellPrice = tr.css(r'td:nth-child(5)').get()
            self.sort += 1

            items.append(ItemItem(
                icon=icon,
                name=name,
                description=description,
                buyPrice=buyPrice,
                sellPrice=sellPrice,
                category=category,
                subcategory=subcategory,
                tradable=tradable,
                extra={'Capture Bonus': bonus, },
                sort=self.sort,
            ))
        return items

    def parseTableSource(self, response, table, category, subcategory, tradable):
        items = []
        for tr in table.css(r'tbody tr'):
            td = tr.css(r'td:nth-child(1)')
            if not td:
                continue
            iconsrc = td.css(r'a:nth-child(1)::attr(href)').get('')
            name = td.css(r'a:nth-child(2)::attr(title)').get('')
            if not name or not iconsrc:
                continue
            icon = TemtemImageItem(
                image_url={'url': response.urljoin(iconsrc)},
            )
            description = tr.css(r'td:nth-child(2)').get()
            source = tr.css(r'td:nth-child(3)').get()
            self.sort += 1

            items.append(ItemItem(
                icon=icon,
                name=name,
                description=description,
                category=category,
                subcategory=subcategory,
                tradable=tradable,
                extra={'Source': source, },
                sort=self.sort,
            ))
        return items

    def parseTableLocation(self, response, table, category, subcategory, tradable):
        items = []
        for tr in table.css(r'tbody tr'):
            td = tr.css(r'td:nth-child(1)')
            if not td:
                continue
            iconsrc = td.css(r'a:nth-child(1)::attr(href)').get('')
            name = td.css(r'a:nth-child(2)::attr(title)').get('')
            if not name or not iconsrc:
                continue
            icon = TemtemImageItem(
                image_url={'url': response.urljoin(iconsrc)},
            )
            description = tr.css(r'td:nth-child(2)').get()
            location = tr.css(r'td:nth-child(3)').get()
            buyPrice = tr.css(r'td:nth-child(4)').get()
            self.sort += 1

            items.append(ItemItem(
                icon=icon,
                name=name,
                description=description,
                category=category,
                subcategory=subcategory,
                tradable=tradable,
                extra={'Location': location, },
                buyPrice=buyPrice,
                sort=self.sort,
            ))
        return items

    def parseTableKey(self, response, table, category, subcategory, tradable):
        items = []
        for tr in table.css(r'tbody tr'):
            td = tr.css(r'td:nth-child(1)')
            if not td:
                continue
            iconsrc = td.css(r'a:nth-child(1)::attr(href)').get('')
            name = td.css(r'a:nth-child(2)::attr(title)').get('')
            if not name or not iconsrc:
                continue
            icon = TemtemImageItem(
                image_url={'url': response.urljoin(iconsrc)},
            )
            description = tr.css(r'td:nth-child(2)').get()
            # location = tr.css(r'td:nth-child(3)').get()
            # buyPrice = tr.css(r'td:nth-child(4)').get()
            self.sort += 1

            items.append(ItemItem(
                icon=icon,
                name=name,
                description=description,
                category=category,
                subcategory=subcategory,
                tradable=tradable,
                sort=self.sort,
                # extra={'Location': location, },
                # buyPrice=buyPrice,
            ))
        return items

    def parseTableQuest(self, response, table, category, subcategory, tradable):
        items = []
        for tr in table.css(r'tbody tr'):
            td = tr.css(r'td:nth-child(1)')
            if not td:
                continue
            iconsrc = td.css(r'a:nth-child(1)::attr(href)').get('')
            name = td.css(r'a:nth-child(2)::attr(title)').get('')
            if not name or not iconsrc:
                continue
            icon = TemtemImageItem(
                image_url={'url': response.urljoin(iconsrc)},
            )
            description = tr.css(r'td:nth-child(2)').get()
            quest = tr.css(r'td:nth-child(3)').get()
            # buyPrice = tr.css(r'td:nth-child(4)').get()
            self.sort += 1

            items.append(ItemItem(
                icon=icon,
                name=name,
                description=description,
                category=category,
                subcategory=subcategory,
                tradable=tradable,
                extra={'Quest': quest, },
                sort=self.sort,
                # buyPrice=buyPrice,
            ))
        return items

    def parse(self, response):
        subcategories = ['Scents', 'Valuable',
                         'Encounter Rate Boosters', 'Other']
        for subcategory in subcategories:
            table = response.css(
                r'.mw-parser-output > h3:contains("%s") + p + table.wikitable' % (subcategory,))
            for item in self.parseTable(response, table, 'General', subcategory, True):
                yield item
        subcategories = ['Experience']
        for subcategory in subcategories:
            table = response.css(
                r'.mw-parser-output > h3:contains("%s") + p + table.wikitable' % (subcategory,))
            for item in self.parseTableNoBuyPrice(response, table, 'General', subcategory, False):
                yield item
        table = response.css(
            r'.mw-parser-output > h2:contains("Capture") + p + table.wikitable')
        for item in self.parseTableCapture(response, table, 'Capture', 'TemCard', False):
            yield item
        subcategories = ['Combat Medicine']
        for subcategory in subcategories:
            table = response.css(
                r'.mw-parser-output > h3:contains("%s") + p + table.wikitable' % (subcategory,))
            for item in self.parseTable(response, table, 'Medicine', subcategory, True):
                yield item
        subcategories = ['TV Fruits', 'TV Candies',
                         'TV Smoothies', 'TV Essences', 'Telomere Hacks', 'Telomere Hotfixes', 'Telomere Bugs']
        for subcategory in subcategories:
            table = response.css(
                r'.mw-parser-output > h3:contains("%s") + p + table.wikitable' % (subcategory,))
            for item in self.parseTable(response, table, 'Performance', subcategory, False):
                yield item
        subcategories = ['Level', 'Evolution']
        for subcategory in subcategories:
            table = response.css(
                r'.mw-parser-output > h3:contains("%s") + p + table.wikitable' % (subcategory,))
            for item in self.parseTableNoBuyPrice(response, table, 'Performance', subcategory, True):
                yield item

        subcategories = ['PvE Gear', 'PvP Gear']
        for subcategory in subcategories:
            table = response.css(
                r'.mw-parser-output > h3:contains("%s") + p + table.wikitable' % (subcategory,))
            for item in self.parseTableSource(response, table, 'Gear', subcategory, True):
                yield item
        subcategories = ['Breeding Gear']
        for subcategory in subcategories:
            table = response.css(
                r'.mw-parser-output > h3:contains("%s") + p + table.wikitable' % (subcategory,))
            for item in self.parseTable(response, table, 'Gear', subcategory, True):
                yield item
        subcategories = ['TV Gear']
        for subcategory in subcategories:
            table = response.css(
                r'.mw-parser-output > h3:contains("%s") + p + table.wikitable' % (subcategory,))
            for item in self.parseTableLocation(response, table, 'Gear', subcategory, True):
                yield item

        table = response.css(
            r'.mw-parser-output > h2:contains("%s") + p + table.wikitable' % ('Key',))
        for item in self.parseTableKey(response, table, 'Key', 'Regular', True):
            yield item
        subcategories = ['Quest Items']
        for subcategory in subcategories:
            table = response.css(
                r'.mw-parser-output > h3:contains("%s") + p + table.wikitable' % (subcategory,))
            for item in self.parseTableQuest(response, table, 'Key', subcategory, True):
                yield item
