#!/usr/env python3

def parseInt(i):
    if not i or i == '-':
        return -1
    return int(i)


def parseStrList(l):
    texts = []
    for i in l:
        i = i.strip()
        if i:
            texts.append(i)
    return ' '.join(texts)
