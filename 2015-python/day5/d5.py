from typing import List, Set

# https://adventofcode.com/2015/day/5

def isNiceString(line: str) -> bool:
    assert line is not None
    assert len(line) != 0

    bannedSubstrings = ['ab','cd','pq','xy']
    for banned in bannedSubstrings:
        if banned in line:
            return False

    numberOfVowels = 0
    hasDoubledCharacter = False
    allowedVowels: Set[str] = {'a', 'e', 'i', 'o', 'u'}

    for idx in range(len(line)):
        c = line[idx]
        if c in allowedVowels:
            numberOfVowels += 1
        if idx+1 < len(line) and c == line[idx+1]:
            hasDoubledCharacter = True

    # print(f'{line} -> {numberOfVowels}, {hasDoubledCharacter}')
    return numberOfVowels >= 3 and hasDoubledCharacter

def isNiceString2(line: str) -> bool:
    assert line is not None
    assert len(line) != 0

    # todo: first part - KMP algorithm would be better
    nonOverlappingRepetitionFound = False
    repeatedLetterWithOneBetweenFound = False

    for i in range(len(line)):
        if nonOverlappingRepetitionFound and repeatedLetterWithOneBetweenFound:
            break
        
        if not nonOverlappingRepetitionFound:
            for j in range(i+2, len(line)):
                if i+1 >= len(line) or j+1 >= len(line):
                    break
                if line[i] == line[j] and line[i+1] ==line[j+1]:
                    nonOverlappingRepetitionFound = True
                    break

        if not repeatedLetterWithOneBetweenFound:
            boundaryOk = i + 2 < len(line)
            if boundaryOk and line[i] == line[i+2]:
                repeatedLetterWithOneBetweenFound = True

    return repeatedLetterWithOneBetweenFound and nonOverlappingRepetitionFound

# nice words
for i in ['ugknbfddgicrmopn', 'aaa']:
    assert isNiceString(i), f'{i} should be true'

# naughty words
for i in ['ab', 'cd', 'pq', 'xy', 'jchzalrnumimnmhp', 'haegwjzuvuyypxyu', 'dvszwmarrgswjxmb']:
    assert not isNiceString(i), f'{i} should be false'

for i in ['qjhvhtzxzqqjkmpb', 'xxyxx']:
    assert isNiceString2(i), f'{i} should be true'

for i in ['uurcxstgmygtbstg', 'ieodomkazucvgmuy']:
    assert not isNiceString2(i), f'{i} should be false'

with open('input.txt') as f:
    lines: List[str] = f.readlines()
    numberOfNiceOld = 0
    numberOfNiceNew = 0
    for l in lines:
        if isNiceString(l):
            numberOfNiceOld+=1
        if isNiceString2(l):
            numberOfNiceNew += 1
    
    assert numberOfNiceOld == 258
    assert numberOfNiceNew == 53

    print(f'nice {numberOfNiceOld}, niceNew: {numberOfNiceNew}')