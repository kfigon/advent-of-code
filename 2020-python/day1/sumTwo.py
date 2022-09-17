from typing import List, Optional, Dict
from inputData import getInputData

# find 2 entries that sum to 2020
# and multiply those 2 numbers together


# O(2n)
def sumExpense(ar: List[int]) -> List[int]:
    out: List[int] = []
    d: Dict[int, int] = {}
    for i in ar:
        if i in d:
            d[i] += 1
        else:
            d[i] = 1

    for i in ar:
        toFind = 2020-i
        if toFind != i and toFind in d and d[toFind] > 0:
            d[toFind] -= 1
            d[i] -= 1
            # print(f'{i} + {toFind}')
            out.append(i*toFind)
        elif toFind == i and toFind in d and d[toFind] > 1:
            d[toFind] -= 2
            # print(f'{i} + {toFind}')
            out.append(i*toFind)
    # print(out)
    return out
         
assert sumExpense([1721, 979, 366, 299, 675, 1456]) == [514579]
assert sumExpense([1721, 979, 2020, 0, 675, 1456]) == [0]
assert sumExpense([1721, 979, 1010, 1, 675, 1456]) == []
assert sumExpense([1721, 979, 1010, 1, 675, 1010]) == [1020100]

def parseInputData(data: str) -> List[int]:
    return list(map(lambda x: int(x) ,data.splitlines()))

res = sumExpense(parseInputData(getInputData()))
print(res)