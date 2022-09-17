from typing import List

# map - # tree, . empty space
# pattern repeates infinitely to right
# start at left-top corner, idx 0
# using given steps - right/down, count how many trees we'll hit

# ==================================================
# part 2
# Time to check the rest of the slopes - you need to minimize the probability of a sudden arboreal stop, after all.
# Determine the number of trees you would encounter if, for each of the following slopes, you start at the top-left corner and traverse the map all the way to the bottom:

# Right 1, down 1.
# Right 3, down 1. (This is the slope you already checked.)
# Right 5, down 1.
# Right 7, down 1.
# Right 1, down 2.
# In the above example, these slopes would find 2, 7, 3, 4, and 2 tree(s) respectively; multiplied together, these produce the answer 336.

# What do you get if you multiply together the number of trees encountered on each of the listed slopes?

pattern: str ='''..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#'''

def isTree(line: str, idx: int) -> bool:
    return line[idx % len(line)] == '#'

def countTrees(pat: str, rightStep: int, downStep: int) -> int:
    print(f'parsing pattern, args: r {rightStep}, d {downStep}')
    lines: List[str] = pat.splitlines()
    numberOfTrees = 0
    colIdx = rightStep
    for rowIdx in range(downStep, len(lines), downStep):
        if isTree(lines[rowIdx], colIdx):
            print(f'tree at: {rowIdx}, {colIdx}')
            numberOfTrees += 1
        colIdx += rightStep

    print(f'result: {numberOfTrees}\n')
    return numberOfTrees

assert countTrees(pattern, rightStep=3, downStep=1) == 7
assert countTrees(pattern, rightStep=2, downStep=2) == 1
assert countTrees(pattern, rightStep=1, downStep=2) == 2

with open('inputData.txt', 'r') as f:
    content = f.read()
    assert countTrees(content, 3,1) == 167
    inputs = [(1,1), (3,1), (5,1), (7,1), (1,2)]
    result=1
    for right, down in inputs:
        result *= countTrees(content, right, down)
    assert result == 736527114
    print(f'result for multiple slopes: {result}')
    