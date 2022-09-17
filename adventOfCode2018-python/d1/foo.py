# https://adventofcode.com/2018/day/1

from typing import List


def read() -> List[int]:
    with open('input.txt') as f:
        return [int(line) for line in f]

def p1() -> int:
    return sum(read())

def p2() -> int:
    currentFreq = 0
    freq = set([currentFreq])
    i = 0
    data = read()
    while True:
        currentFreq += data[i]
        if currentFreq in freq:
            return currentFreq
        freq.add(currentFreq)
        i = (i + 1) % len(data)
    
    return -1

assert p1() == 454
assert p2() == 566
