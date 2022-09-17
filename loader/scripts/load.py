#!/usr/bin/env python3

from genericpath import isfile
import os
import json
from django.db import transaction

from temtem.models import Type
import shutil
from pathlib import Path
from bs4 import BeautifulSoup

# Build paths inside the project like this: BASE_DIR / 'subdir'.
BASE_DIR = Path(__file__).resolve().parent
filesfolder = os.path.join(BASE_DIR, 'files')

# 创建文件储存目录
if not os.path.exists(filesfolder):
    os.mkdir(filesfolder)


def copyfile(src, dst):
    if os.path.exists(os.path.join(dst, os.path.basename(src))):
        return
    shutil.copy2(src, dst)


def updateHTML(html):
    '''更新HTML文本，给图片添加前缀'''
    s = BeautifulSoup(html, 'html.parser')
    for a in s.find_all('a'):
        if not a['href'].startswith('http'):
            a['href'] = 'https://temtem.wiki.gg'+a['href']
            a['target'] = '_blank'
    for img in s.find_all('img'):
        if not img['src'].startswith('http'):
            img['src'] = 'https://temtem.wiki.gg'+img['src']
    return s


@transaction.atomic
def loadType(path):
    Type.objects.all().delete()
    types = json.load(open(path))
    sort = 0
    for t in types:
        icon = os.path.basename(t['icon']['path'])
        copyfile(t['icon']['path'], filesfolder)

        trivia = []
        for tt in t['trivia']:
            s = updateHTML(tt)
            trivia.append(s.li.encode_contents().decode())
        tt = Type(
            name=t['name'],
            icon=icon,
            comment=t['comment'],
            trivia=trivia,
            effective_against=t['effectiveAgainst'],
            ineffective_against=t['ineffectiveAgainst'],
            resistant_to=t['resistantTo'],
            weak_to=t['weakTo'],
            sort=sort,
        )
        tt.save()
        sort += 1


def run(*args):
    if len(args) != 1:
        print('load <json data folder>')
        return
    print('文件保存在: %s' % filesfolder)
    folder = args[0]
    typeJsonFile = os.path.join(folder, 'type.json')
    loadType(typeJsonFile)
