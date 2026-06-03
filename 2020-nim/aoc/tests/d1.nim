import unittest
import strutils
import sequtils
import sets
import options

const example = """1721
979
366
299
675
1456"""

proc parse(data: string): seq[int] = data.strip().splitLines().map(parseInt)

proc p1(data: string): Option[int] =
  let nums = parse(data)
  let s = nums.toHashSet()
  let target = 2020
  for i in nums:
    let toFind = target - i
    if toFind in s:
      return some(toFind * i)

  return none[int]()

# 3 sum = nested 2 sum
proc p2(data: string): Option[int] = 
  let nums = parse(data)
  let target = 2020
  for i in 0..<nums.len:
   var hash = initHashSet[int]() 

   for j in i+1..<nums.len:
    let toFind = target - nums[i]-nums[j]
    if toFind in hash:
      return some[int](nums[i]*nums[j]*toFind)

    hash.incl(nums[j])

  return none[int]()

suite "d1":
  let data = [
    (name:"p1 ex", file:example, fn: p1, exp:514579),
    (name:"p1", file:readFile("resources/d1.txt"), fn: p1, exp:719796),
    (name: "p2 ex", file: example, fn: p2, exp: 241861950),
    (name:"p2", file:readFile("resources/d1.txt"), fn: p2, exp:144554112),
  ]
  for (name, file, fn, exp) in data:
    test name: check fn(file) == some(exp)
