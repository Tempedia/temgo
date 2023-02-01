# TemGO


This project contains 3 parts.

## TemGO

**TemGO** is a web service created with Go. It provides http api for client.

## Spider

**Spider** is a web crawler created with Python Scrapy. It's used to fetch Temtem data from [their offical website](https://temtem.wiki.gg).

**Spider** save Temtem data as JSON format, also it will download web assets like image and video.

## Loader

**Loader** is a Python [Django](https://www.djangoproject.com/) project. It doesn't provide any http api, but has a script used to load JSON data to database.

Be careful, every time you run **Loader**, will overwrite previous data.

*WHY NOT USE DJANGO TO PROVIDER HTTP API?*

*I READLY DON'T KNOW....*

Maybe we should remove **TemGO**, use Django to make this project simpler.


## Flow

1. **Spider** fetches Temtem data and save them as JSON.
2. **Loader** reads JSON data and insert them into database like PostgreSQL.
3. **TemGO** reads PostgreSQL and provides http api for client.
