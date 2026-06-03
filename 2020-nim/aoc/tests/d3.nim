import unittest
import options
import strutils

const example = """..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#"""

type 
    Grid = object
        rows: seq[string]
    Field = enum
        Empty, Tree  

proc dims(g: Grid): tuple[height: int, width: int] = 
    let height = g.rows.len()
    let width = g.rows[0].len()
    (height: height, width: width)

proc get(g: Grid, col: int, row: int): Option[Field] = 
    let (height, width) = g.dims
    
    case g.rows[row mod height][col mod width]:
        of '.': some[Field](Empty)
        of '#': some[Field](Tree)
        else: none[Field]()
        
proc parse(map: string): Grid = 
    let rows = map.strip.splitLines()
    if rows.len() == 0:
        raise newException(ValueError, "empty rows provided")
    Grid(rows: rows)

proc p1(map: string): int = 
    let right = 3
    let down = 1
    let g = parse(map)

    let height = g.dims().height
    var count = 0
    var c = 0
    var r = 0
    while r < height:
        let el = g.get(c,r).get()
        if el == Tree: inc count
    
        c = c + right
        r = r + down

    count

suite "d3":
    test "p1 example": check p1(example) == 7
    test "p1": check p1(readFile("resources/d3.txt")) == 167