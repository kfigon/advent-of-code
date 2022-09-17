from typing import List, Tuple


def read() -> List[str]:
    with open('input.txt') as f:
        return [line.strip() for line in f.readlines()]

def scanLine(line: str) -> Tuple[bool, bool]:
    freq = {}
    for c in line:
        if c in freq:
            freq[c] += 1
        else:
            freq[c] = 1
    occurences = set(map(lambda x: x[1], freq.items()))
    return (2 in occurences, 3 in occurences)

    
def p1() -> int:
    twos = 0
    thirds = 0
    for line in read():
        d = scanLine(line)
        if d[0]:
            twos+=1
        if d[1]:
            thirds+=1
    return twos*thirds

def p2() -> str:
    data = read()
    for i in range(len(data)):
        for j in range(len(data)):
            if diff(data[i], data[j]) == 1:
                return diffStr(data[i], data[j])
    return ''

def diff(a: str, b: str) -> int:
    assert len(a) == len(b)
    diff = 0
    for i in range(len(a)):
        if a[i] != b[i]:
            diff += 1
    return diff

def diffStr(a: str, b: str) -> str:
    assert len(a) == len(b)
    same = ''
    for i in range(len(a)):
        if a[i] == b[i]:
            same += a[i]
    return same

assert p1() == 6150
assert p2() == 'rteotyxzbodglnpkudawhijsc'
