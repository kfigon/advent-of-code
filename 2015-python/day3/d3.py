from typing import Tuple, Set

# https://adventofcode.com/2015/day/3

class Traverser:
    def __init__(self):
        self.lastCoordinate: Tuple[int,int] = (0,0)
        self.visitedNodes: Set[Tuple[int,int]] = set()
        self.visitedNodes.add(self.lastCoordinate)
    
    def go(self, c: str):
        x,y = self.lastCoordinate
        if c == "^":
            y-=1
        elif c == "v":
            y+=1
        elif c == ">":
            x+=1
        elif c == "<":
            x-=1
        else:
            raise Exception("invalid char!")
        newCoord = (x,y)

        self.visitedNodes.add(newCoord)
        self.lastCoordinate = newCoord

    def numberOfUniqueHouses(self) -> int:
        return len(self.visitedNodes)

def isDataCorrect(inputData: str) -> bool:
    return inputData is not None and inputData != ''

def traversePath(inputData: str) -> int:
    if not isDataCorrect(inputData):
        return 1
    t = Traverser()
    for c in inputData:
        t.go(c)

    return t.numberOfUniqueHouses()

def traverseParallel(inputData: str) -> int:
    if not isDataCorrect(inputData):
        return 1

    santa = Traverser()
    robot = Traverser()

    for i,c in enumerate(inputData):
        if i % 2 == 0:
            santa.go(c)
        else:
            robot.go(c)

    # out = len(santa.visitedNodes)
    # for i in robot.visitedNodes:
    #     if i not in santa.visitedNodes:
    #         out+=1
    # return out
    return len(santa.visitedNodes.union(robot.visitedNodes))

assert traversePath('>') == 2
assert traversePath('^>v<') == 4
assert traversePath('^v^v^v^v^v') == 2

assert traverseParallel('^v') == 3
assert traverseParallel('^>v<') == 3
assert traverseParallel('^v^v^v^v^v') == 11

with open('input.txt') as f:
    content = f.read()
    p1 = traversePath(content)
    p2 = traverseParallel(content)

    assert p1 == 2592
    assert p2 == 2360
