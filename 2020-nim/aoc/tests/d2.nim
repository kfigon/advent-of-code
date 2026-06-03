import unittest
import strutils

const example = """1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc"""

type
    Spec = object
        min: int
        max: int
        c: char

proc validPasswordP1(spec: Spec, line: string): bool = 
    let cnt = line.count(spec.c)
    cnt >= spec.min and cnt <= spec.max

proc validPasswordP2(spec: Spec, line: string): bool = 
    let first = line[spec.min-1] == spec.c
    let second = line[spec.max-1] == spec.c
    first xor second

proc solve(data: string, validator: proc(spec: Spec, l: string): bool): int = 
  var valid = 0
  for line in data.strip().splitLines():
    let got = line.split({'-',' ', ':'})
    let s = Spec(
      min: got[0].parseInt,
      max: got[1].parseInt,
      c: got[2][0]
    )
    if validator(s, got[4]):
      inc valid
        
  return valid

proc p1(data: string): int = solve(data, validPasswordP1)
proc p2(data: string): int = solve(data, validPasswordP2)
  
suite "d2":
  let data = [
    (name: "p1 ex", fn: p1, input: example, exp: 2),
    (name: "p1", fn: p1, input: readFile("resources/d2.txt"), exp: 500),
    (name: "p2 ex", fn: p2, input: example, exp: 1),
    (name: "p2", fn: p2, input: readFile("resources/d2.txt"), exp: 313),
  ]
  for (name, fn, input, exp) in data:
    test name:
        check fn(input) == exp