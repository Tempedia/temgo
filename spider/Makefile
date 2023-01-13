
.PHONY: all

OUTPUT=./output

all: temtem condition course item location technique trait type

condition:
	scrapy crawl condition -O $(OUTPUT)/condition.json -L INFO

temtem:
	scrapy crawl temtem -O $(OUTPUT)/temtem.json -L INFO

course:
	scrapy crawl course -O $(OUTPUT)/course.json -L INFO

item:
	scrapy crawl item -O $(OUTPUT)/item.json -L INFO

location:
	scrapy crawl location -O $(OUTPUT)/location.json  -L INFO

technique:
	scrapy crawl technique -O $(OUTPUT)/technique.json -L INFO

trait:
	scrapy crawl trait -O $(OUTPUT)/trait.json  -L INFO

type:
	scrapy crawl type -O $(OUTPUT)/type.json  -L INFO


