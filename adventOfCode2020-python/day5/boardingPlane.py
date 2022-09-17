from typing import List, Tuple, Callable, Optional

# input - boarding spaces in a plane
# this airline uses binary space partitioning to seat people. 
# A seat might be specified like FBFBBFFRLR, 
# where F means "front", B means "back", L means "left", and R means "right"

# The first 7 characters will either be F or B; 
# these specify exactly one of the 128 rows on the plane (numbered 0 through 127). 
# Each letter tells you which half of a region the given seat is in. Start with the whole list of rows; 
# the first letter indicates whether the seat is in the front (0 through 63) or the back (64 through 127). 
# The next letter indicates which half of that region the seat is in, and so on until you're left with exactly one row.

# For example, consider just the first seven characters of FBFBBFFRLR:

# Start by considering the whole range, rows 0 through 127.
# F means to take the lower half, keeping rows 0 through 63.
# B means to take the upper half, keeping rows 32 through 63.
# F means to take the lower half, keeping rows 32 through 47.
# B means to take the upper half, keeping rows 40 through 47.
# B keeps rows 44 through 47.
# F keeps rows 44 through 45.
# The final F keeps the lower of the two, row 44.
# The last three characters will be either L or R; 
# these specify exactly one of the 8 columns of seats on the plane (numbered 0 through 7). 
# The same process as above proceeds again, this time with only three steps. 
# L means to keep the lower half, while R means to keep the upper half.

# For example, consider just the last 3 characters of FBFBBFFRLR:

# Start by considering the whole range, columns 0 through 7.
# R means to take the upper half, keeping columns 4 through 7.
# L means to take the lower half, keeping columns 4 through 5.
# The final R keeps the upper of the two, column 5.
# So, decoding FBFBBFFRLR reveals that it is the seat at row 44, column 5.

# Every seat also has a unique seat ID: multiply the row by 8, then add the column. In this example, the seat has ID 44 * 8 + 5 = 357.

# Here are some other boarding passes:

# BFFFBBFRRR: row 70, column 7, seat ID 567.
# FFFBBBFRRR: row 14, column 7, seat ID 119.
# BBFFBBFRLL: row 102, column 4, seat ID 820.
# As a sanity check, look through your list of boarding passes. What is the highest seat ID on a boarding pass?


# ========================================================
# v2
# It's a completely full flight, so your seat should be the only missing boarding pass in your list. 

# However, there's a catch: some of the seats at the very front and back of the
# plane don't exist on this aircraft, so they'll 
# be missing from your list as well.

# Your seat wasn't at the very front or back, though; 
# the seats with IDs +1 and -1 from yours will be in your list.

def getFirstPart(content: str) -> str:
    return content[:7]

def getSecondPart(content: str) -> str:
    return content[7:]

def validate(content: str):
    if len(content) != 10:
        raise Exception(f'invalid length of {content} ({len(content)} != 10)')

    for c in getFirstPart(content):
        if c!='F' and c!='B':
            raise Exception(f'unknown character {c} in first part of {content}')

    for c in getSecondPart(content):
        if c!='R' and c!='L':
            raise Exception(f'unknown character {c} in second part of {content}')

def getCoordinates(content: str) -> Tuple[int, int]:
    validate(content)
    mappingFB: Callable[[str], int] = lambda c: 1 if c=='B' else 0
    mappingRL: Callable[[str], int] = lambda c: 1 if c=='R' else 0
    
    sumaFB = 0
    for i,c in enumerate(reversed(getFirstPart(content))):
        sumaFB += (2**i) * mappingFB(c)

    sumaRL = 0
    for i,c in enumerate(reversed(getSecondPart(content))):
        sumaRL += (2**i) * mappingRL(c)

    # print(f'{content} -> {sumaFB}, {sumaRL}')
    return sumaFB, sumaRL

def countSeatId(seatCoordinates: Tuple[int, int]) -> int:
    return seatCoordinates[0]*8 + seatCoordinates[1]

def findPlaceWithSkip(listOfSeats: List[int]) -> Optional[int]:
    listOfSeats.sort()

    for i in range(len(listOfSeats)-1):
        currentSeat = listOfSeats[i]
        nextSeat = listOfSeats[i+1]

        if (currentSeat+1) != nextSeat:
            print(f'my seat {currentSeat+1}')
            return currentSeat+1
    return None

assert getCoordinates('BFFFBBFRRR') == (70,7)
assert getCoordinates('FFFBBBFRRR') == (14,7)
assert getCoordinates('BBFFBBFRLL') == (102,4)

assert countSeatId((44,5)) == 357
assert countSeatId((70,7)) == 567
assert countSeatId((14,7)) == 119
assert countSeatId((102,4)) == 820

with open('inputData.txt') as f:
    lines = f.readlines()
    maxSeatId = None
    listOfSeats: List[int] = []
    for line in lines:
        currentSeatId = countSeatId(getCoordinates(line.strip()))
        if maxSeatId is None or currentSeatId > maxSeatId:
            maxSeatId = currentSeatId
        
        listOfSeats.append(currentSeatId)

    assert maxSeatId == 850
    assert findPlaceWithSkip(listOfSeats) == 599
