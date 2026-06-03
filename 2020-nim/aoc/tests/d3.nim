import unittest
import options
import strutils
import sequtils

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

proc solve(g: Grid, steps: tuple[right: int, down: int]): int =
    let (right, down) = steps

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

proc p1(map: string): int = 
    let g = parse(map)
    solve(g, (3,1))

proc p2(map: string): int = 
    let g = parse(map)
    let steps = [
        (1,1),
        (3,1),
        (5,1),
        (7,1),
        (1,2),
    ]
    steps.mapIt(solve(g, it)).foldl(a*b)

suite "d3":
    test "p1 example": check p1(example) == 7
    test "p1": check p1(readFile("resources/d3.txt")) == 167
    test "p2 example": check p2(example) == 336
    test "p2": check p2(readFile("resources/d3.txt")) == 736527114