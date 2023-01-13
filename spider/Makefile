
.PHONY: all

OUTPUT=./output

all: temtem condition technique-course item location technique trait type

condition:
	scrapy crawl condition -O $(OUTPUT)/condition.json

temtem:
	scrapy crawl temtem -O $(OUTPUT)/temtem.json

technique-course:
	scrapy crawl technique-course -O $(OUTPUT)/technique-course.json

item:
	scrapy crawl item -O $(OUTPUT)/item.json

location:
	scrapy crawl location -O .$(OUTPUT)/location.json

technique:
	scrapy crawl technique -O $(OUTPUT)/technique.json

trait:
	scrapy crawl trait -O $(OUTPUT)/trait.json

type:
	scrapy crawl type -O $(OUTPUT)/type.json


