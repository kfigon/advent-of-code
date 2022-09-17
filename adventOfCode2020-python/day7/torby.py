from typing import List, Tuple, Dict, Set
import re


# Due to recent aviation regulations, many rules (your puzzle input) are being enforced about bags and their contents; 
# bags must be color-coded and must contain specific quantities of other color-coded bags. 
# Apparently, nobody responsible for these regulations considered how long they would take to enforce!

# For example, consider the following rules:

# light red bags contain 1 bright white bag, 2 muted yellow bags.
# dark orange bags contain 3 bright white bags, 4 muted yellow bags.
# bright white bags contain 1 shiny gold bag.
# muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
# shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
# dark olive bags contain 3 faded blue bags, 4 dotted black bags.
# vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
# faded blue bags contain no other bags.
# dotted black bags contain no other bags.
# These rules specify the required contents for 9 bag types. 

# In this example, every faded blue bag is empty, every vibrant plum 
# bag contains 11 bags (5 faded blue and 6 dotted black), and so on.

# You have a shiny gold bag. If you wanted to carry it in at least 
# one other bag, how many different bag colors would be valid for the outermost bag? 
# (In other words: how many colors can, eventually, contain at least one shiny gold bag?)

# In the above rules, the following options would be available to you:

# A bright white bag, which can hold your shiny gold bag directly.
# A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
# A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
# A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.

# So, in this example, the number of bag colors that can eventually contain at least one shiny gold bag is 4.

# How many bag colors can eventually contain at least one shiny gold bag? 
# (The list of rules is quite long; make sure you get all of it.)

# =============================--- Part Two ---====================================
# It's getting pretty expensive to fly these days - not because of ticket prices, but because of the ridiculous number of bags you need to buy!

# Consider again your shiny gold bag and the rules from the above example:

# faded blue bags contain 0 other bags.
# dotted black bags contain 0 other bags.
# vibrant plum bags contain 11 other bags: 5 faded blue bags and 6 dotted black bags.
# dark olive bags contain 7 other bags: 3 faded blue bags and 4 dotted black bags.

# So, a single shiny gold bag must contain 1 dark olive bag (and the 7 bags within it) 
# plus 2 vibrant plum bags (and the 11 bags within each of those): 1 + 1*7 + 2 + 2*11 = 32 bags!

# Of course, the actual rules have a small chance of going several levels deeper than this example; 
# be sure to count all of the bags, even if the nesting becomes topologically impractical!

# Here's another example:

# shiny gold bags contain 2 dark red bags.
# dark red bags contain 2 dark orange bags.
# dark orange bags contain 2 dark yellow bags.
# dark yellow bags contain 2 dark green bags.
# dark green bags contain 2 dark blue bags.
# dark blue bags contain 2 dark violet bags.
# dark violet bags contain no other bags.

# In this example, a single shiny gold bag must contain 126 other bags.

# How many individual bags are required inside your single shiny gold bag?

def parseLine(line: str) -> Tuple[str, Dict[str, int]]:
    rules = line.split('contain')
    bagRule = rules[0].replace('bags','').strip()
    if 'no other bags' in rules[1]: 
        return bagRule, {}

    containedBags = rules[1].split(',')
    d: Dict[str, int] = {}
    for bag in containedBags:
        cleanedString = bag.replace('bags', '').replace('bag','').replace('.','').strip()
        result = re.findall(r'(\d+) (\w+ \w+)', cleanedString)

        assert result is not None
        assert len(result) == 1
        d[result[0][1]] = int(result[0][0])
    
    return bagRule, d

def parseData(content: str) -> Dict[str, Dict[str, int]]:
    d: Dict[str, Dict[str, int]] = {}
    rules = content.splitlines()
    for rule in rules:
        parsed = parseLine(rule)
        assert parsed[0] not in d

        d[parsed[0]] = parsed[1]
    return d

def findWhereBagIsLocated(rules: Dict[str, Dict[str, int]], bagColorToFind: str) -> int:
    cnt: int = 0
    visited: Set[str] = set()

    def traverse(node: str, toFind: str):
        bags = rules[node]
        if toFind not in bags or node in visited or len(bags) == 0:
            return
        
        visited.add(node)
        nonlocal cnt
        cnt += 1
        for r in rules:
            traverse(r, node)

    for r in rules:
        traverse(r, bagColorToFind)

    print(f'{bagColorToFind} can be reached through {cnt} bags')
    return cnt


def findHowManyBagsInside(rules: Dict[str, Dict[str, int]], bagToAnalyze: str) -> int:
    def traverse(node: str) -> int:
        childrenNodes = rules[node]
        if len(childrenNodes) == 0:
            return 0
        
        c = 0
        for r in childrenNodes:
            c += childrenNodes[r] + childrenNodes[r]*traverse(r)
        return c

    cnt = traverse(bagToAnalyze)

    print(f'you need {cnt} bags in single {bagToAnalyze}')
    return cnt

inputData = '''light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.'''

inputData2='''shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.'''

rules = inputData.splitlines()
parsedRules = list(map(parseLine, rules))

assert parsedRules[0] == ('light red', {'bright white': 1, 'muted yellow':2})
assert parsedRules[2] == ('bright white', {'shiny gold': 1})
assert parsedRules[7] == ('faded blue', {})
assert parseLine('shiny green bags contain 2 bright lavender bags, 3 shiny olive bags, 4 mirrored violet bags, 5 posh white bags.') == ('shiny green', {'bright lavender': 2, 'shiny olive': 3, 'mirrored violet':4, 'posh white':5})

# d = parseData(inputData)
# for k in d:
#     print(f'{k} -> {d[k]}')

assert findWhereBagIsLocated(parseData(inputData), 'shiny gold') == 4
assert findWhereBagIsLocated(parseData(inputData), 'light red') == 0

assert findHowManyBagsInside(parseData(inputData), 'shiny gold') == 32
assert findHowManyBagsInside(parseData(inputData2), 'shiny gold') == 126

with open('inputData.txt') as f:
    fileRules = parseData(f.read())
    assert len(fileRules) == 594
    assert findWhereBagIsLocated(fileRules, 'shiny gold') == 124
    assert findHowManyBagsInside(fileRules, 'shiny gold') == 34862