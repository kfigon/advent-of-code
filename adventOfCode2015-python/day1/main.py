from typing import List, Tuple

# https://adventofcode.com/2015/day/1

# Santa is trying to deliver presents in a large apartment building, but he can't
# find the right floor - the directions he got are a little confusing.
# He starts on the ground floor (floor 0) and then follows the instructions one character at a time.
#
# An opening parenthesis, (, means he should go up one floor, and a closing parenthesis, ),
# means he should go down one floor.
#
# The apartment building is very tall, and the basement is very deep;
# he will never find the top or bottom floors.
#
# For example:
#
# (()) and ()() both result in floor 0.
# ((( and (()(()( both result in floor 3.
# ))((((( also results in floor 3.
# ()) and ))( both result in floor -1 (the first basement level).
# ))) and )())()) both result in floor -3.
# To what floor do the instructions take Santa?

def floorFinder(data: str) -> int:
    floor = 0
    for c in data:
        floor += parseFloor(c)
    return floor

def parseFloor(c : str) -> int:
    if c == '(':
        return 1
    elif c == ')':
        return -1
    else:
        raise Exception(f'invalid character {c}')

def findFirstBasementFloor(data: str) -> int:
    currentFloor = 0
    for i,c in enumerate(data):
        currentFloor += parseFloor(c)
        if currentFloor == -1:
            return i+1

    raise Exception(f'-1 not found!')

testCases: List[Tuple[str,int]] = [
    ('(())', 0),
    ('()()', 0),
    ('(((', 3),
    ('(()(()(', 3),
    ('))(((((', 3),
    (')())())', -3),
]
for t in testCases:
    assert floorFinder(t[0]) == t[1]

with open('inputData.txt') as f:
    inputData = f.read()
    assert floorFinder(inputData) == 280
    print(findFirstBasementFloor(inputData))
    assert findFirstBasementFloor(inputData) == 1797