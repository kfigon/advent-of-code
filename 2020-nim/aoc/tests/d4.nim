import unittest
import options
import strutils
import tables
import sequtils
import sets

const example = """ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in"""

# byr (Birth Year) - four digits; at least 1920 and at most 2002.
# iyr (Issue Year) - four digits; at least 2010 and at most 2020.
# eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
# hgt (Height) - a number followed by either cm or in:
# If cm, the number must be at least 150 and at most 193.
# If in, the number must be at least 59 and at most 76.
# hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
# ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
# pid (Passport ID) - a nine-digit number, including leading zeroes.

proc withinRange(s: string, min: int, max: int): bool = 
    let v = parseInt(s)
    v >= min and v <= max

proc byr(s: string): bool = withinRange(s, 1920, 2002)
proc iyr(s: string): bool = withinRange(s, 2010, 2020)

proc eyr(s: string): bool = withinRange(s, 2020, 2030)

proc hgt(s: string): bool = 
    if s.endsWith("cm"):
        return withinRange(s[0..<s.len-2], 150, 193)
    elif s.endsWith("in"):
        return withinRange(s[0..<s.len-2], 59, 76)
    false
        
proc hcl(s: string): bool = 
    if s.len() != 7: return false
    if s[0] != '#': return false
    let allowed = toHashSet(['a','b','c','d','e','f','0','1','2', '3', '4', '5', '6', '7', '8', '9'])
    for c in s[1..^1]:
        if not allowed.contains(c):
            return false
    true

proc ecl(s: string): bool = 
    let allowed = toHashSet(["amb", "blu", "brn", "gry", "grn", "hzl", "oth"])
    allowed.contains(s)

proc pid(s: string): bool = 
    if s.len() != 9: return false
    try:
        discard parseInt(s)
        return true
    except:
        return false

type ValidatorFn = proc(s: string): bool

let byrFn: ValidatorFn = byr
let iyrFn: ValidatorFn = iyr
let eyrFn: ValidatorFn = eyr
let hgtFn: ValidatorFn = hgt
let hclFn: ValidatorFn = hcl
let eclFn: ValidatorFn = ecl
let pidFn: ValidatorFn = pid

var validations: Table[string, ValidatorFn] = {
    "byr": byrFn,
    "iyr": iyrFn,
    "eyr": eyrFn,
    "hgt": hgtFn,
    "hcl": hclFn,
    "ecl": eclFn,
    "pid": pidFn,
}.toTable

proc parse(data: string): seq[Table[string,string]] = 
    var res: seq[Table[string,string]] = @[]
    var t = initTable[string,string]()
    for line in data.strip.splitLines():
        if line == "":
            res.add(t)
            t = initTable[string,string]()
        for v in line.splitWhitespace():
            let pair = v.split({':'})
            t[pair[0]] = pair[1]
    res

proc validP1(t: Table[string,string]): bool = 
    for k in validations.keys():
        if not t.contains(k):
            return false
    true

proc validP2(t: Table[string,string]): bool = 
    for k in validations.keys():
        if not t.contains(k):
            return false
        if not validations[k](t[k]):
            return false
    true

proc solve(data: string, validator: proc(t: Table[string,string]): bool): int =
    let got = data.parse()
    var cnt = 0
    for e in got:
        if validator(e):
            inc cnt
    cnt

proc p1(data: string): int = solve(data, validP1)
proc p2(data: string): int = solve(data, validP2)
        
suite "d4":
    test "p1 example": check p1(example) == 2
    test "p1": check p1(readFile("resources/d4.txt")) == 256
    test "p2 example": check p2(example) == 2
    test "p2": check p2(readFile("resources/d4.txt")) == 198