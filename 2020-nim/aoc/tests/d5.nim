import unittest
import strutils
import options
import math
import strformat
import sequtils
import algorithm

proc seatId(d: string): int = 
    var row = 0
    var column = 0
    for i,c in d[0..6].pairs:
        case c:
        of 'F': discard
        of 'B': 
            row = row + int(pow(2.0, float(6-i)))
        else: raise newException(ValueError, fmt"invalid value {c} in string {d}")

    for i, c in d[7..9].pairs:
        case c:
        of 'L': discard
        of 'R':
            column = column + int(pow(2.0, float(2-i)))
        else: raise newException(ValueError, fmt"invalid value {c} in string {d}")
    row * 8 + column

proc p1(d: string): Option[int] = 
    var res = none[int]()
    for line in d.splitLines:
        let seat = seatId(line)
        if res.isNone or res.get < seat:
            res = some(seat)
    res

proc p2(d: string): Option[int] = 
    let all = d.splitLines.mapIt(seatId(it)).sorted()
    for i in 0..<all.len-1:
        let this = all[i]
        let next = all[i+1]
        if next - this != 1:
            return some(this+1)

    none[int]()


suite "d5 seatid":
    let data = [
        ("FBFBBFFRLR", 357),
        ("BFFFBBFRRR", 567),
        ("FFFBBBFRRR", 119),
        ("BBFFBBFRLL", 820)
    ]
    for d in data:
        test d[0]: 
            check seatId(d[0]) == d[1]

    let file = readFile("resources/d5.txt")

    test "p1":
        check p1(file) == some(850)    
    test "p2":
        check p2(file) ==  some(599)