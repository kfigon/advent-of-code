from typing import List, Optional, Dict
from inputData import getInputData

# find 3 entries that sum to 2020
# and multiply those 3 numbers together

def sumExpense(ar: List[int]) -> List[int]:
    out: List[int] = []
    d: Dict[int, int] = {}
    for i in ar:
        if i in d:
            d[i] += 1
        else:
            d[i] = 1

    for i in range(len(ar)):
        for j in range(i, len(ar)):
            toFind = 2020-ar[i]-ar[j]
            if toFind in d and d[toFind] > 0:
                d[toFind] -= 1
                d[ar[i]] -= 1
                d[ar[j]] -= 1
                # print(f'{i} + {toFind}')
                out.append(ar[i]*ar[j]*toFind)
    # print(out)
    
    # brute force
    # for i in range(len(ar)):
    #     for j in range(i,len(ar)):
    #         for k in range(j,len(ar)):
    #             if ar[j]+ar[i]+ar[k] == 2020:
    #                 out.append(ar[j]*ar[i]*ar[k])
    return out

assert sumExpense([1721, 979, 366, 299, 675, 1456]) == [241861950]

def parseInputData(data: str) -> List[int]:
    return list(map(lambda x: int(x) ,data.splitlines()))

res = sumExpense(parseInputData(getInputData()))
assert res == [144554112]
print(res)