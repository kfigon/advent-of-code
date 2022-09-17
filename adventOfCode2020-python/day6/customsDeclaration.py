from typing import List, Dict

# you filled customs declaration form:
# 26 yes/no questions marked a-z. Answers "yes" are written, rest are "no"

# people fill that in groups and e.g. in one:
# abcx
# abcy
# abcz

# In this group, there are 6 questions to which anyone answered "yes": a, b, c, x, y, and z. 
# (Duplicate answers to the same question don't count extra; each question counts at most once.)

# Each group's answers are separated by a blank line, and within each group, each person's answers are on a single line. For example:

# abc

# a
# b
# c

# ab
# ac

# a
# a
# a
# a

# b

# This list represents answers from five groups:

# The first group contains one person who answered "yes" to 3 questions: a, b, and c.
# The second group contains three people; combined, they answered "yes" to 3 questions: a, b, and c.
# The third group contains two people; combined, they answered "yes" to 3 questions: a, b, and c.
# The fourth group contains four people; combined, they answered "yes" to only 1 question, a.
# The last group contains one person who answered "yes" to only 1 question, b.
# In this example, the sum of these counts is 3 + 3 + 3 + 1 + 1 = 11.

# For each group, count the number of questions to which anyone answered "yes". What is the sum of those counts?

# ---------------------------------------------
# v2
# you notice that you misread one word in the instructions
# You don't need to identify the questions to which anyone answered "yes"; 
# you need to identify the questions to which everyone answered "yes"!

# This list represents answers from five groups:

# In the first group, everyone (all 1 person) answered "yes" to 3 questions: a, b, and c.
# In the second group, there is no question to which everyone answered "yes".
# In the third group, everyone answered yes to only 1 question, a. Since some people did not answer "yes" to b or c, they don't count.
# In the fourth group, everyone answered yes to only 1 question, a.
# In the fifth group, everyone (all 1 person) answered "yes" to 1 question, b.
# In this example, the sum of these counts is 3 + 0 + 1 + 1 + 1 = 6.

# For each group, count the number of questions to which everyone answered "yes". What is the sum of those counts?

inputData = '''abc

a
b
c

ab
ac

a
a
a
a

b'''

def parseData(content: str) -> List[str]:
    groups: List[str] = content.split('\n\n')
    allAnswers: List[str] = []

    for group in groups:
        peopleAnswers = group.splitlines()
        groupAsnwer = ''
        for personAns in peopleAnswers:
            for a in personAns:
                groupAsnwer += a
        allAnswers.append(groupAsnwer)
    return allAnswers

def countAnswers(answersInGroup: str) -> int:
    return len(set(answersInGroup))

def sumAnswers(answersInGroups: List[str]) -> int:
    return sum(map(countAnswers, answersInGroups))

def parseDataPerPerson(content: str) -> List[List[str]]:
    groups: List[str] = content.split('\n\n')
    allAnswers: List[List[str]] = []

    for group in groups:
        peopleAnswers: List[str] = group.splitlines()
        allAnswers.append(peopleAnswers)
    return allAnswers

def countAnswersPerPerson(answersPerPerson: List[str]) -> int:
    d : Dict[str, int] = {}
    for person in answersPerPerson:
        for a in person:
            if a in d:
                d[a]+=1
            else:
                d[a]=1
    allSame = 0
    numberOfPeople = len(answersPerPerson)
    for key in d:
        if d[key] == numberOfPeople:
            allSame += 1
    return allSame

def sumAnswersPerPerson(answersPerGroupPerPerson: List[List[str]]) -> int:
    return sum(map(countAnswersPerPerson, answersPerGroupPerPerson))

assert parseData(inputData) == ['abc','abc','abac','aaaa','b']
assert sumAnswers(parseData(inputData)) == 11

assert parseDataPerPerson(inputData) == [['abc'],['a','b','c'],['ab','ac'],['a','a','a','a'],['b']]
assert countAnswersPerPerson(['abc']) == 3
assert countAnswersPerPerson(['a','b','c']) == 0
assert countAnswersPerPerson(['ab','ac']) == 1
assert countAnswersPerPerson(['a','a','a','a','a']) == 1
assert countAnswersPerPerson(['b']) == 1

with open('inputData.txt') as f:
    content = f.read()
    parsedAnswers = parseData(content)
    answerNum = sumAnswers(parsedAnswers)
    assert answerNum == 6335

    answersPerPerson = parseDataPerPerson(content)
    assert sumAnswersPerPerson(answersPerPerson) == 3392