
from curses import meta
import scrapy

from ..items import DownloadImageItem, StatusConditionItem, TemtemImageItem


class ConditionSpider(scrapy.Spider):
    name = 'condition'
    start_urls = ['https://temtem.wiki.gg/wiki/Status_Conditions']

    def parseGroup(self, response, title, group):
        x = response.css(r'.mw-parser-output > h2:contains("%s")' % (title,))
        items = []
        while x:
            x = x.css(r':scope + *')
            tag = x.xpath('name()').get()
            if tag not in ('p', 'h3', 'ul'):
                break
            if tag == 'h3':
                h3 = x
                name = h3.css(':scope >span::text').get()
                desc = h3.css(r':scope + p').get()
                img = h3.css(r':scope + p img')[0]
                icon = response.urljoin(img.css('::attr(src)').get())
                items.append(StatusConditionItem(
                    name=name,
                    desc=desc,
                    icon=DownloadImageItem(image_url=icon),
                    group=group,
                ))
        return items

    def parse(self, response):
        items = []
        items.extend(self.parseGroup(
            response, 'Negative Status Conditions', 'Negative'))
        items.extend(self.parseGroup(
            response, 'Positive Status Conditions', 'Positive'))
        items.extend(self.parseGroup(
            response, 'Neutral Status Conditions', 'Neutral'))
        items.extend(self.parseGroup(
            response, 'Other Status Conditions', 'Other'))
        for item in items:
            techniques = response.css(r'a:contains("List of all %s Techniques")' %
                                      (item['name'],))
            traits = response.css(r'a:contains("List of all %s Traits")' %
                                  (item['name'],))
            if techniques:
                yield response.follow(techniques[0], self.parse_technique, meta={'item': item, 'trait': traits[0] if traits else None})

    def parse_technique(self, response):
        item = response.meta['item']
        item['techniques'] = []
        item['traits'] = []
        trait = response.meta['trait']

        for title in response.css(r'div.mw-content-ltr a::attr(title)').getall():
            item['techniques'].append(title)

        if not trait:
            yield item
            return
        yield response.follow(trait, self.parse_trait, meta={'item': item})

    def parse_trait(self, response):
        item = response.meta['item']
        item['traits'] = []
        for title in response.css(r'div.mw-content-ltr a::attr(title)').getall():
            item['traits'].append(title)
        yield item
