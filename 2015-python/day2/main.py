from typing import Tuple

# https://adventofcode.com/2015/day/2

def calculateSize(l: int, w: int, h:int) -> int:
    a = l*w
    b = w*h
    c = h*l
    return 2*a + 2*b + 2*c + min(a,b,c)

def calculateRibbon(l: int, w: int, h:int) -> int:
    peri = lambda a,b: 2*(a+b)
    x = min(peri(l,w), peri(w,h), peri(l,h))
    return x + l*w*h

def parseInput(content: str) -> Tuple[int,int,int]:
    vals = content.split('x')
    if len(vals) != 3:
        raise Exception(f'invalid input {content}')
    return int(vals[0]),int(vals[1]),int(vals[2])

assert calculateSize(l=2, w=3, h=4) == 58
assert calculateSize(l=1, w=1, h=10) == 43

assert calculateRibbon(l=2, w=3, h=4) == 34
assert calculateRibbon(l=1, w=1, h=10) == 14

assert parseInput('2x3x4') == (2,3,4)
assert parseInput('1x1x10') == (1,1,10)

with open('inputData.txt') as f:
    lines = f.readlines()
    total = 0
    totalRibbon = 0
    for line in lines:
        parsed = parseInput(line)
        total += calculateSize(*parsed)
        totalRibbon += calculateRibbon(*parsed)
    print(f'total result {total}, ribbon {totalRibbon}')
