
from dataclasses import dataclass
from typing import List
import re

@dataclass
class Coord:
    id: str
    x: int
    y: int
    width: int
    height: int

def parse() -> List[Coord]:
    with open('input.txt') as f:
        return [parseLine(line) for line in f.readlines()]

def parseLine(line: str) -> Coord:
    data = re.findall('#(\d+) @ (\d+),(\d+): (\d+)x(\d+)', line)[0]
    return Coord(id=data[0], 
        x=int(data[1]),
        y=int(data[2]),
        width=int(data[3]),
        height=int(data[4]))

class Map:
    def __init__(self) -> None:
        self.d = [ [0] * 1000 for _ in range(1000)]

    def process(self, coords: List[Coord]):
        for c in coords:
            for row in range(c.y, c.y+c.height):
                for col in range(c.x, c.x+c.width):
                    self.d[row][col] += 1
    
    def calc(self) -> int:
        cnt = 0
        for row in self.d:
            for el in row:
                if el > 1:
                    cnt+=1
        return cnt
    
    def findNonOverlapping(self, coords: List[Coord]) -> str:
        for c in coords:
            allOnes = True
            for row in range(c.y, c.y+c.height):
                for col in range(c.x, c.x+c.width):
                    if self.d[row][col] != 1:
                        allOnes = False
            if allOnes:
                return c.id

        return ''

coords = parse()
m = Map()
m.process(coords)

assert m.calc() == 116920
assert m.findNonOverlapping(coords) == '382'