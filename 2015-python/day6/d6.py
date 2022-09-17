from typing import Tuple, List, Optional, Callable
import re
import datetime

# https://adventofcode.com/2015/day/6

class Coordinate:
    def __init__(self, x: int, y: int):
        self.x = x
        self.y = y
    def __str__(self) -> str:
        return f'({self.x},{self.y})'
    def __eq__(self, other) -> bool:
        return self.x==other.x and self.y == other.y

class Rectangle:
    def __init__(self, coordA: Coordinate, coordB: Coordinate):
        self.coordA = coordA
        self.coordB = coordB
        assert self.coordA.x <= self.coordB.x
        assert self.coordA.y <= self.coordB.y

    def size(self) -> int:
        difX = self.coordB.x - self.coordA.x +1
        difY = self.coordB.y - self.coordA.y +1
        return difX * difY
    def rows(self) -> Tuple[int,int]:
        return self.coordA.y, self.coordB.y
    def cols(self) -> Tuple[int,int]:
        return self.coordA.x, self.coordB.x
    def __str__(self) -> str:
        return f'{self.coordA} : {self.coordB}'
    def __repr__(self) -> str:
        return str(self)
    def __eq__(self, other) -> bool:
        return self.coordA == other.coordA and self.coordB == other.coordB

class Operation:
    TURN_ON = 0
    TURN_DOWN = 1
    TOGGLE = 2

class AbstractGridStrategy:
    def getNumLit(self) -> int:
        raise Exception('to implement')
    def handleOperation(self, operation, row, col):
        raise Exception('to implement')

class BrightnessStrategy(AbstractGridStrategy):
    def __init__(self):
        self.table = [[0 for j in range(1000)] for i in range(1000)]
        self.numOfLit = 0
    def handleOperation(self, operation, row, col):
        currentLight = self.table[row][col]
        operationIncrease = 0
        if operation == Operation.TURN_ON:
            operationIncrease = 1
        elif operation == Operation.TURN_DOWN:
            operationIncrease = -1
        elif operation == Operation.TOGGLE:
            operationIncrease = 2
            
        finalValue = max(0, currentLight + operationIncrease)
        self.numOfLit = self.numOfLit - currentLight + finalValue
        self.table[row][col] = finalValue

    def getNumLit(self) -> int:
        return self.numOfLit

class StateStrategy(AbstractGridStrategy):
    def __init__(self):
        self.table = [[False for j in range(1000)] for i in range(1000)]
        self.numOfLit = 0
    def handleOperation(self, operation, row, col):
        currentLight = self.table[row][col]
        if operation == Operation.TURN_ON and currentLight == False:
            self.table[row][col] = True
            self.numOfLit += 1
        elif operation == Operation.TURN_DOWN and currentLight == True:
            self.table[row][col] = False
            self.numOfLit -= 1
        elif operation == Operation.TOGGLE and currentLight == False:
            self.table[row][col] = True
            self.numOfLit += 1
        elif operation == Operation.TOGGLE and currentLight == True:
            self.table[row][col] = False
            self.numOfLit -= 1
    def getNumLit(self) -> int:
        return self.numOfLit

class Grid:
    def __init__(self):
        self.changeLog: List[Tuple[Operation, Rectangle]] = []
    def add(self, op: Operation, rect: Rectangle):
        self.changeLog.append((op, rect))

    def run(self, st: AbstractGridStrategy) -> int:
        start = datetime.datetime.now()
        for operation, rectangle in self.changeLog:
            rowStart, rowEnd = rectangle.rows()
            colStart, colEnd = rectangle.cols()

            for row in range(rowStart, rowEnd+1):
                for col in range(colStart, colEnd+1):
                    st.handleOperation(operation, row, col)
                    
        dif = datetime.datetime.now() - start
        print(f'took {dif}')
        return st.getNumLit()

    def applyChangeLog(self) -> int:
        return self.run(StateStrategy())

    def applyChangeLog2(self) -> int:
        return self.run(BrightnessStrategy())

def parseCommand(inputStr: str) -> Tuple[int, Rectangle]:
    if inputStr is None or len(inputStr) == 0:
        raise ValueError('Empty or None string provided')
    res = re.findall(r'(\w+\s?\w+?) (\d+,\d+) through (\d+,\d+)', inputStr)
    if res is None or len(res) != 1 or len(res[0]) != 3:
        raise ValueError(f'Cound not parse {inputStr}')

    res = res[0]

    op = findOperation(res[0])
    leftCoord = parseCoordinate(res[1])
    rightCoord = parseCoordinate(res[2])
    if op is None or leftCoord is None or rightCoord is None:
        raise ValueError(f'Error in parsing {inputStr}')

    return (op, Rectangle(leftCoord, rightCoord))

def parseCoordinate(inputStr: str) -> Optional[Coordinate]:
    parts = inputStr.split(',')
    if len(parts) != 2:
        return None
    
    return Coordinate(int(parts[0]), int(parts[1]))

def findOperation(inputStr: str) -> Optional[int]:
    if "turn on" in inputStr:
        return Operation.TURN_ON
    elif "toggle" in inputStr:
        return Operation.TOGGLE
    elif "turn off" in inputStr:
        return Operation.TURN_DOWN
    return None

if __name__ == "__main__":
    with open('input.txt') as f:
        lines = f.readlines()
        g = Grid()
        for i in lines:
            operation, rectangle = parseCommand(i)
            g.add(operation, rectangle)
        
        # todo: after abstracting strategy - much slower! 7s vs 14 and 12 vs 19

        # took 0:00:07.045773
        result1 = g.applyChangeLog()
        assert result1 == 569999
        print(f'p1 {result1}')
        
        # took 0:00:12.519982
        result2 = g.applyChangeLog2()
        print(f'p2: {result2}')
        assert result2 == 17836115
