#!/usr/bin/env python3

import os
import json
from unicodedata import name
from django.db import transaction

from temtem.models import Temtem, Type
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
        return os.path.basename(src)
    shutil.copy2(src, dst)
    return os.path.basename(src)


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
    for t in types:
        icon = copyfile(t['icon']['path'], filesfolder)

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
            sort=t['sort'],
        )
        tt.save()


@transaction.atomic
def loadTemtem(path):
    Temtem.objects.all().delete()
    temtems = json.load(open(path))
    for t in temtems:
        icon = copyfile(t['normalIcon']['path'], filesfolder)
        lumaIcon = copyfile(t['lumaIcon']['path'], filesfolder)
        cry = copyfile(t['cry']['path'], filesfolder)
        height = float(t['height'].split('cm')[0])
        weight = float(t['weight'].split('kg')[0])
        tvYield = t['tvYield']
        for k in t['tvYield']:
            if t['tvYield'][k]:
                tvYield[k]=int(t['tvYield'][k])
            else:
                tvYield[k]=0
        evolvesTo = []
        for e in t['evolvesTo']:
            m = {
                'to': e['to'],
            }
            if '+' in e['method'] and ('level' in e['method'] or 'Level' in e['method']):
                m['method'] = 'levelplus'
                m['level'] = int(e['method'].split('+')[1].split(' ')[0])
                if 'Female' in e['method']:
                    m['gender'] = 'female'
                elif 'Male' in e['method']:
                    m['gender'] = 'male'
            elif e['method'] == 'Trade':
                m['method'] = 'trade'
            elif 'TVs' in e['method']:
                m['method'] = 'tv'
                m['tv'] = int(e['method'].split(' ')[0])
            elif 'at' in e['method']:
                m['method'] = 'place'
                m['place'] = e['method']
            else:
                raise Exception('unknown evolution method %s' % e['method'])
            evolvesTo.append(m)
        stats = {}
        for k in t['stats']:
            stats[k] = {}
            s = t['stats'][k]
            stats[k]['base'] = int(s['base'])
            s50 = s['50'].split('-')
            stats[k]['50'] = {
                'from': int(s50[0]),
                'to': int(s50[1]),
            }
            s100 = s['100'].split('-')
            stats[k]['100'] = {
                'from': int(s100[0]),
                'to': int(s100[1]),
            }
        gallery = []
        for g in t['gallery']:
            gallery.append({
                'fileid': copyfile(g['path'], filesfolder),
                'text': g['text']
            })
        renders = []
        for g in t['renders']:
            renders.append({
                'fileid': copyfile(g['path'], filesfolder),
                'text': g['text']
            })
        trivia = []
        for tt in t['trivia']:
            trivia.append(updateHTML(tt).li.encode_contents().decode())
        description = {}
        for k in t['description']:
            description[k] = updateHTML(
                t['description'][k]).encode_contents().decode()
        tem = Temtem(
            no=t['no'],
            name=t['name'],
            type=t['type'],
            gender_ratio=t['genderRatio'],
            catch_rate=t['catchRate'],
            experience_yield_modifier=t['experienceYieldModifier'],
            icon=icon,
            luma_icon=lumaIcon,
            traits=t['traits'],
            description=description,
            cry=cry,
            height=height,
            weight=weight,
            tv_yield=tvYield,
            evolves_to=evolvesTo,
            stats=stats,
            type_matchup=t['typeMatchup'],
            techniques=t['techniques'],
            trivia=trivia,
            gallery=gallery,
            renders=renders,
        )
        tem.save()


def run(*args):
    if len(args) != 1:
        print('load <json data folder>')
        return
    print('文件保存在: %s' % filesfolder)
    folder = args[0]
    loadType(os.path.join(folder, 'type.json'))
    loadTemtem(os.path.join(folder, 'temtem.json'))
