#!/usr/bin/env python3

import os
import json
from unicodedata import name
from django.db import transaction

from temtem.models import Temtem, TemtemBreedingTechnique, TemtemCourseTechnique, TemtemLevelingUpTechnique, TemtemTechnique, TemtemTrait, Type
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
    if not html:
        return ''
    s = BeautifulSoup(html, 'html.parser')
    for a in s.find_all('a'):
        a.unwrap()
        # if not a['href'].startswith('http'):
        #     a['href'] = 'https://temtem.wiki.gg'+a['href']
        #     a['target'] = '_blank'
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
    TemtemLevelingUpTechnique.objects.all().delete()
    TemtemBreedingTechnique.objects.all().delete()
    TemtemCourseTechnique.objects.all().delete()
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
                tvYield[k] = int(t['tvYield'][k])
            else:
                tvYield[k] = 0
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
            html = updateHTML(t['description'][k])
            description[k] = html.encode_contents().decode() if html else None
        subspecies = []
        for s in t['subspecies']:
            flag = False
            for ss in subspecies:
                if s['group'] == ss['type']:
                    if s['text'] == 'normal':
                        ss['icon'] = copyfile(s['path'], filesfolder)
                    elif s['text'] == 'luma':
                        ss['luma_icon'] = copyfile(s['path'], filesfolder)
                    flag = True
                    break
            if not flag:
                tt = {
                    'type': s['group'],
                }
                if s['text'] == 'normal':
                    tt['icon'] = copyfile(s['path'], filesfolder)
                elif s['text'] == 'luma':
                    tt['luma_icon'] = copyfile(s['path'], filesfolder)
                subspecies.append(tt)
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
            # techniques=t['techniques'],
            trivia=trivia,
            gallery=gallery,
            renders=renders,
            subspecies=subspecies,
        )
        tem.save()
        for tech in t['techniques']['leveling_up']:
            technique = TemtemLevelingUpTechnique(
                temtem=tem.name,
                level=tech['level'],
                technique_name=tech['technique'],
                stab=tech['stab'],
                group=tech.get('group', '') or '',
            )
            technique.save()
        for tech in t['techniques']['course']:
            technique = TemtemCourseTechnique(
                temtem=tem.name,
                course=tech['course'],
                technique_name=tech['technique'],
                stab=tech['stab'],
            )
            technique.save()
        for tech in t['techniques']['breeding']:
            technique = TemtemBreedingTechnique(
                temtem=tem.name,
                parents=tech['parents'],
                technique_name=tech['technique'],
                stab=tech['stab'],
            )
            technique.save()


@transaction.atomic
def loadTemtemTrait(path):
    TemtemTrait.objects.all().delete()
    traits = json.load(open(path))
    for t in traits:
        desc = updateHTML(t['desc']).td.encode_contents().decode()
        trait = TemtemTrait(
            name=t['name'],
            description=desc,
            impact=t['impact'],
            trigger=t['trigger'],
            effect=t['effect'],
        )
        trait.save()


@transaction.atomic
def loadTemtemTechnique(path):
    TemtemTechnique.objects.all().delete()
    techniques = json.load(open(path))
    for t in techniques:
        desc = ''
        i = updateHTML(t['desc'])
        if i:
            desc = i.i.encode_contents().decode()
        technique = TemtemTechnique(
            name=t['name'],
            type=t['type'],
            cls=t['clas'],
            damage=t['dmg'],
            sta_cost=t['sta'],
            hold=t['hold'],
            priority=t['priority'],
            targeting=t['targeting'],
            description=desc,
            video=copyfile(t['video']['path'], filesfolder),
        )

        if t.get('synergyType'):
            synergyDesc = ''
            synergyEffects = ''
            i = updateHTML(t.get('synergyDesc'))
            if i:
                synergyDesc = i.i.encode_contents().decode()
            i = updateHTML(t.get('synergyEffects'))
            if i:
                synergyEffects = i.td.encode_contents().decode()
            technique.synergy_description = synergyDesc
            technique.synergy_type = t['synergyType']
            technique.synergy_effects = synergyEffects
            technique.synergy_damage = t.get('synergyDamage', 0)
            technique.synergy_sta_cost = t.get('synergySta', 0)
            technique.synergy_priority = t.get('synergyPriority', 0)
            technique.synergy_targeting = t.get('synergyTargeting', '')
            technique.synergy_video = copyfile(
                t['synergyVideo']['path'], filesfolder),
        technique.save()


def run(*args):
    if len(args) != 1:
        print('load <json data folder>')
        return
    print('文件保存在: %s' % filesfolder)
    folder = args[0]
    loadType(os.path.join(folder, 'type.json'))
    loadTemtemTrait(os.path.join(folder, 'trait.json'))
    loadTemtemTechnique(os.path.join(folder, 'technique.json'))
    loadTemtem(os.path.join(folder, 'temtem.json'))
