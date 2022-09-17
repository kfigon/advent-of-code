from typing import List, Tuple

# Your plane lands with plenty of time to spare. The final leg of your journey is a ferry
#  that goes directly to the tropical island where you can finally start your vacation.
#   As you reach the waiting area to board the ferry, you realize you're so early, nobody else has even arrived yet!

# By modeling the process people use to choose (or abandon) their seat in the waiting 
# area, you're pretty sure you can predict the best place to sit. You make a 
# quick map of the seat layout (your puzzle input).

# The seat layout fits neatly on a grid. Each position is either floor (.), an 
# empty seat (L), or an occupied seat (#). For example, the initial seat layout might look like this:

# --- Part Two ---
# As soon as people start to arrive, you realize your mistake. People don't 
# just care about adjacent seats - they care about the first seat they can see in each of those eight directions!

# Now, instead of considering just the eight immediately adjacent seats, 
# consider the first seat in each of those eight directions. 
# For example, the empty seat below would see eight occupied seats:

# .......#.
# ...#.....
# .#.......
# .........
# ..#L....#
# ....#....
# .........
# #........
# ...#.....
# The leftmost empty seat below would only see one empty seat, but cannot see any of the occupied ones:

# .............
# .L.L.#.#.#.#.
# .............
# The empty seat below would see no occupied seats:

# .##.##.
# #.#.#.#
# ##...##
# ...L...
# ##...##
# #.#.#.#
# .##.##.
# Also, people seem to be more tolerant than you expected: it now takes five 
# or more visible occupied seats for an occupied seat to become empty 
# (rather than four or more from the previous rules). The other rules still apply: 
# empty seats that see no occupied seats become occupied, seats matching no rule don't change, and floor never changes.

# Given the same starting layout as above, these new rules cause 
# the seating area to shift around as follows:

# L.LL.LL.LL
# LLLLLLL.LL
# L.L.L..L..
# LLLL.LL.LL
# L.LL.LL.LL
# L.LLLLL.LL
# ..L.L.....
# LLLLLLLLLL
# L.LLLLLL.L
# L.LLLLL.LL
# #.##.##.##
# #######.##
# #.#.#..#..
# ####.##.##
# #.##.##.##
# #.#####.##
# ..#.#.....
# ##########
# #.######.#
# #.#####.##
# #.LL.LL.L#
# #LLLLLL.LL
# L.L.L..L..
# LLLL.LL.LL
# L.LL.LL.LL
# L.LLLLL.LL
# ..L.L.....
# LLLLLLLLL#
# #.LLLLLL.L
# #.LLLLL.L#
# #.L#.##.L#
# #L#####.LL
# L.#.#..#..
# ##L#.##.##
# #.##.#L.##
# #.#####.#L
# ..#.#.....
# LLL####LL#
# #.L#####.L
# #.L####.L#
# #.L#.L#.L#
# #LLLLLL.LL
# L.L.L..#..
# ##LL.LL.L#
# L.LL.LL.L#
# #.LLLLL.LL
# ..L.L.....
# LLLLLLLLL#
# #.LLLLL#.L
# #.L#LL#.L#
# #.L#.L#.L#
# #LLLLLL.LL
# L.L.L..#..
# ##L#.#L.L#
# L.L#.#L.L#
# #.L####.LL
# ..#.#.....
# LLL###LLL#
# #.LLLLL#.L
# #.L#LL#.L#
# #.L#.L#.L#
# #LLLLLL.LL
# L.L.L..#..
# ##L#.#L.L#
# L.L#.LL.L#
# #.LLLL#.LL
# ..#.L.....
# LLL###LLL#
# #.LLLLL#.L
# #.L#LL#.L#
# Again, at this point, people stop shifting around and the seating 
# area reaches equilibrium. Once this occurs, you count 26 occupied seats.

# Given the new visibility method and the rule change for occupied 
# seats becoming empty, once equilibrium is reached, how many seats end up occupied?

first='''L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL'''

# Now, you just need to model the people who will be arriving shortly. Fortunately, people are entirely
#  predictable and always follow a simple set  of rules. All decisions are based on the number of occupied
#   seats adjacent to a given seat (one of the eight positions immediately up, down, left, right, or 
#   diagonal from the seat). The following rules are applied to every seat simultaneously:

# If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
# If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
# Otherwise, the seat's state does not change.
# Floor (.) never changes; seats don't move, and nobody sits on the floor.

# After one round of these rules, every seat in the example layout becomes occupied:

second='''#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##'''
# After a second round, the seats with four or more occupied adjacent seats become empty again:

third='''#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##'''
# This process continues for three more rounds:

fourth='''#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##'''

fifth='''#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##'''

sixth='''#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##'''

# At this point, something interesting happens: the chaos stabilizes and further applications of these rules cause 
# no seats to change state! Once people stop moving around, you count 37 occupied seats.

# Simulate your seating area by applying the seating rules repeatedly until no seats change state. How many seats end up occupied?

class SeatMatrix:
    def __init__(self, lines: str):
        self.state: List[str] = lines.splitlines()

    def isEmpty(self, row: int, col: int) -> bool:
        return self.state[row][col] == 'L'
    def isFloor(self, row: int, col: int) -> bool:
        return self.state[row][col] == '.'
    def isOccupied(self, row: int, col: int) -> bool:
        return self.state[row][col] == '#'

    def rows(self) -> int:
        return len(self.state)

    def cols(self) -> int:
        return len(self.state[0])

    def validate(self, row: int, col: int) -> bool:
        return (row >= 0 and row < len(self.state)) and (col >= 0 and col < len(self.state[0]))

    def takeSeat(self, row: int, col: int):
        if not self.isFloor(row, col):
            s = self.state[row]
            self.state[row] = mutateString(s, col, '#')
    
    def freeSeat(self, row: int, col: int):
        if not self.isFloor(row, col):
            s = self.state[row]
            self.state[row] = mutateString(s, col, 'L')

    def numberOfOccupiedNei(self, row: int, col: int) -> int:
        if not self.validate(row, col):
            raise Exception(f'invalid idx {row}; {col}')

        idx = self.adjecentIdxs(row, col)
        out = 0

        for i in idx:
            if self.isOccupied(i[0],i[1]):
                out += 1
        return out

    def numberOfOccupiedSeats(self) -> int:
        out = 0
        for r in range(self.rows()):
            for c in range(self.cols()):
                if self.isOccupied(r,c):
                    out += 1
        return out

    def adjecentIdxs(self, row: int, col: int) -> List[Tuple[int,int]]:
        return adjecentIdxs(row, col, self.rows(), self.cols())
        
    def __eq__(self, other) -> bool:
        if other is None or type(other) != SeatMatrix:
            return False
        if len(self.state) != len(other.state):
            return False

        for i in range(len(self.state)):
            if self.state[i] != other.state[i]:
                return False
        return True
    
    def __hash__(self) -> int:
        h = 0
        for i in self.state:
            h += hash(i)
        return h

def validateInput(data: str) -> bool:
    rows: List[str] = data.splitlines()
    numberOfCols = set(map(lambda x: len(x), rows))
    return len(numberOfCols) == 1

def mutateString(s: str, idx: int, newChar: str) -> str:
    assert len(newChar) == 1
    assert idx >=0 and idx < len(s)
    return s[0:idx] + newChar + s[idx+1:]

def adjecentIdxs(row: int, col: int, rowLen: int, colLen: int) -> List[Tuple[int,int]]:
    candidates = []
    for r in [-1,0,1]:
        for c in [-1,0,1]:
            rNei = row + r
            cNei = col + c
            if (r != 0 or c != 0) and (rNei >=0 and rNei < rowLen) and (cNei >= 0 and cNei < colLen):
                candidates.append((rNei, cNei))
    return candidates


assert mutateString('abcd', 0, 'x') == 'xbcd'
assert mutateString('abcd', 1, 'x') == 'axcd'
assert mutateString('abcd', 2, 'x') == 'abxd'
assert mutateString('abcd', 3, 'x') == 'abcx'

# 012
# 345
# 678
assert adjecentIdxs(0,0,3,3) == [(0,1),(1,0),(1,1)]
assert adjecentIdxs(0,1,3,3) == [(0,0),(0,2),(1,0),(1,1),(1,2)]
assert adjecentIdxs(2,2,3,3) == [(1,1),(1,2),(2,1)]
assert adjecentIdxs(1,1,3,3) == [(0,0),(0,1),(0,2),(1,0),(1,2),(2,0),(2,1),(2,2)]

def doCopy(content: str) -> str:
    out = ''
    for i in content:
        out += i
    return out

def simulate(seats: SeatMatrix) -> Tuple[SeatMatrix, bool]:
    toOccupy: List[Tuple[int,int]] = []
    toFree: List[Tuple[int,int]] = []

    for row in range(seats.rows()):
        for col in range(seats.cols()):
            if seats.isFloor(row, col):
                continue

            numOfNei = seats.numberOfOccupiedNei(row, col)
            if seats.isEmpty(row,col) and numOfNei == 0:
                toOccupy.append((row, col))
            elif seats.isOccupied(row, col) and numOfNei >= 4:
                toFree.append((row, col))
    
    doesChange = False
    for r,c in toOccupy:
        seats.takeSeat(r,c)
    for r,c in toFree:
        seats.freeSeat(r,c)
    
    doesChange = len(toOccupy) != 0 or len(toFree) != 0
    return seats, doesChange

def simulateAndCount(content: str) -> int:
    if not validateInput(content):
        raise Exception('received content is not valid')

    seats = SeatMatrix(content)
    steps = 0
    MAX_STEPS = 500
    while steps < MAX_STEPS:
        seats, doesChange = simulate(seats)
        
        if not doesChange:
            print(f'running stepps {steps}')
            return seats.numberOfOccupiedSeats()
        steps += 1

    raise Exception(f'invalid state!')


inputData = [first, second, third, fourth, fifth, sixth]
for i in inputData:
    assert validateInput(i)

seats = SeatMatrix(doCopy(first))
assert seats == SeatMatrix(doCopy(first))
assert simulate(seats) == (SeatMatrix(doCopy(second)), True)
assert simulate(seats) == (SeatMatrix(doCopy(third)), True)
assert simulate(seats) == (SeatMatrix(doCopy(fourth)), True)
assert simulate(seats) == (SeatMatrix(doCopy(fifth)), True)
assert simulate(seats) == (SeatMatrix(doCopy(sixth)), True)

assert simulate(seats) == (SeatMatrix(doCopy(sixth)), False)
assert simulate(seats) == (SeatMatrix(doCopy(sixth)), False)
assert simulate(SeatMatrix(doCopy(sixth))) == (SeatMatrix(doCopy(sixth)), False)

assert hash(SeatMatrix(first)) == hash(SeatMatrix(first))
assert hash(SeatMatrix(second)) == hash(SeatMatrix(second))
assert hash(SeatMatrix(third)) == hash(SeatMatrix(third))

assert SeatMatrix(first) == SeatMatrix(first)
assert SeatMatrix(second) == SeatMatrix(second)
assert SeatMatrix(third) == SeatMatrix(third)

assert hash(SeatMatrix(first)) != hash(SeatMatrix(fourth))
assert hash(SeatMatrix(first)) != hash(SeatMatrix(second))
assert hash(SeatMatrix(first)) != hash(SeatMatrix(third))

assert SeatMatrix(first) != SeatMatrix(fourth)
assert SeatMatrix(first) != SeatMatrix(second)
assert SeatMatrix(first) != SeatMatrix(third)

assert simulateAndCount(doCopy(first)) == 37

with open('inputData.txt') as f:
    res1 = simulateAndCount(f.read())
    assert res1 == 2270
    print(f'first result -> {res1}')