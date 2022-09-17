from typing import List
import hashlib


# https://adventofcode.com/2015/day/4


# brute force, can it be done better?
def findLowestNumProducingHash(secretKey: str, prefixToFind: str) -> int:
    for i in range(10000000):
        hashInput = secretKey + str(i)
        calculated = hashlib.md5(hashInput.encode('utf-8')).hexdigest()
        if calculated.startswith(prefixToFind):
            return i
    raise Exception(f'not found for {secretKey}')

part1Target = '00000'
part2Target = '000000'
assert findLowestNumProducingHash('abcdef',part1Target) == 609043
assert findLowestNumProducingHash('pqrstuv', part1Target) == 1048970
assert findLowestNumProducingHash('bgvyzdsv', part1Target) == 254575

p1 = findLowestNumProducingHash('bgvyzdsv', part1Target)
p2 = findLowestNumProducingHash('bgvyzdsv', part2Target)

assert p1 == 254575
assert p2 == 1038736
print(f'results: p1: {p1}, p2: {p2}')

