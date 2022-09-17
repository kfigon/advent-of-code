from typing import List, Dict, Iterator, Callable

# The automatic passport scanners are slow because they're having trouble detecting which passports have all required fields. The expected fields are as follows:

# byr (Birth Year)
# iyr (Issue Year)
# eyr (Expiration Year)
# hgt (Height)
# hcl (Hair Color)
# ecl (Eye Color)
# pid (Passport ID)
# cid (Country ID)

# Passport data is validated in batch files (your puzzle input). 
# Each passport is represented as a sequence of key:value pairs separated by spaces or newlines. 
# Passports are separated by blank lines.

# The first passport is valid - all eight fields are present. 
# 
# The second passport is invalid - it is missing hgt (the Height field).
# 
# The third passport is interesting; the only missing field is cid, so it looks 
# like data from North Pole Credentials, not a passport at all! 
# Surely, nobody would mind if you made the system temporarily ignore missing cid fields. Treat this "passport" as valid.
# 
# The fourth passport is missing two fields, cid and byr. Missing cid is fine, but missing any other field is not, so this passport is invalid.
# According to the above rules, your improved system would report 2 valid passports.
# Count the number of valid passports - those that have all required fields. Treat cid as optional

# ================================================
# v2

# You can continue to ignore the cid field, but each other field has strict rules about what values are valid for automatic validation:

# byr (Birth Year) - four digits; at least 1920 and at most 2002.
# iyr (Issue Year) - four digits; at least 2010 and at most 2020.
# eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
# hgt (Height) - a number followed by either cm or in:
# If cm, the number must be at least 150 and at most 193.
# If in, the number must be at least 59 and at most 76.
# hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
# ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
# pid (Passport ID) - a nine-digit number, including leading zeroes.
# cid (Country ID) - ignored, missing or not.
# Your job is to count the passports where all required fields are both present and valid according to the above rules. Here are some example values:

# byr valid:   2002
# byr invalid: 2003

# hgt valid:   60in
# hgt valid:   190cm
# hgt invalid: 190in
# hgt invalid: 190

# hcl valid:   #123abc
# hcl invalid: #123abz
# hcl invalid: 123abc

# ecl valid:   brn
# ecl invalid: wat

# pid valid:   000000001
# pid invalid: 0123456789
# Here are some invalid passports:

# eyr:1972 cid:100
# hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

# iyr:2019
# hcl:#602927 eyr:1967 hgt:170cm
# ecl:grn pid:012533040 byr:1946

# hcl:dab227 iyr:2012
# ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

# hgt:59cm ecl:zzz
# eyr:2038 hcl:74454a iyr:2023
# pid:3556412378 byr:2007
# Here are some valid passports:

# pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
# hcl:#623a2f

# eyr:2029 ecl:blu cid:129 byr:1989
# iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

# hcl:#888785
# hgt:164cm byr:2001 iyr:2015 cid:88
# pid:545766238 ecl:hzl
# eyr:2022

# iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
# Count the number of valid passports - those that have all required fields and valid values. Continue to treat cid as optional. 
# In your batch file, how many passports are valid?

inputData = '''ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in'''

invalidPassportsData='''eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007'''

validPassportsData='''pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719'''

def isFourDigitStr(content: str) -> bool:
    return len(content) == 4 and content.isdigit()

def validateByr(content: str) -> bool:
    if not isFourDigitStr(content):
        return False
    num = int(content)
    return num >= 1920 and num <= 2002

def validateIyr(content: str) -> bool:
    if not isFourDigitStr(content):
        return False
    num = int(content)
    return num >= 2010 and num <= 2020

def validateEyr(content: str) -> bool:
    if not isFourDigitStr(content):
        return False
    num = int(content)
    return num >= 2020 and num <= 2030

def validateHgt(content: str) -> bool:
    digits = lambda x: content[:-2].isdigit()

    if content.endswith('cm') and digits(content):
        num = int(content[:-2])
        return num >= 150 and num <= 193
    elif content.endswith('in') and digits(content):
        num = int(content[:-2])
        return num >= 59 and num <= 76
    return False

def validateHcl(content: str) -> bool:
    if not content.startswith('#'):
        return False

    rest = content[1:]
    def areAf(x: str) -> bool:
        for i in x:
            if i.isdigit():
                continue

            charCode = ord(i)
            if charCode < ord('a') or charCode > ord('f'):
                return False

        return True
    return len(rest) == 6 and areAf(rest)

def validateEcl(content: str) -> bool:
    valid = ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']
    return content in valid

def validatePid(content: str) -> bool:
    return len(content) == 9 and content.isdigit()

mandatoryFields: Dict[str, Callable[[str], bool]] = {
    'byr': validateByr, 
    'iyr': validateIyr,
    'eyr': validateEyr,
    'hgt': validateHgt,
    'hcl': validateHcl,
    'ecl': validateEcl,
    'pid': validatePid
    }

def mapEntry(content: str) -> Dict[str, str]:
    lines = content.split('\n')
    entries: Iterator[List[str]] = map(lambda line: line.split(' '), lines)
    out: Dict[str, str] = {}
    
    for params in entries:
        for e in params:
            keyPair = e.split(':')
            key = keyPair[0]
            value = keyPair[1]
            if key in out:
                raise Exception(f'Invalid entry! {e}')
            out[key] = value

    return out

def parseData(contents: str) -> List[Dict[str, str]]:
    passportEntries: List[str] = contents.split('\n\n')
    passports: List[Dict[str,str]] = list(map(mapEntry, passportEntries))
    return passports

def validatePassport(passport: Dict[str, str]) -> bool:
    missingFields: List[str] = []
    invalidFields: List[str] = []
    for mandatoryField in mandatoryFields.keys():
        if mandatoryField not in passport:
            missingFields.append(mandatoryField)
        else:
            if not mandatoryFields[mandatoryField](passport[mandatoryField]):
                invalidFields.append(mandatoryField)

    if len(missingFields) == 0 and len(invalidFields) == 0:
        return True
    else:
        print(f'passport {passport} INVALID, report: ', end='')
        if len(missingFields) != 0:
            print(f'missing fields: {missingFields}, ', end='')
        if len(invalidFields) != 0:
            print(f'invalid fields: ', end='')
            for i in invalidFields:
                print(f'{i}, provided: {passport[i]}, ', end='')
        print()
        return False


def controlPassports(passports: List[Dict[str,str]]) -> int:
    numberOfValidPassports = 0
    for p in passports:
        if validatePassport(p):
            numberOfValidPassports += 1
    print(f'number of valid passports in batch: {numberOfValidPassports}')
    return numberOfValidPassports

assert controlPassports(parseData(inputData)) == 2
assert controlPassports(parseData(invalidPassportsData)) == 0
assert controlPassports(parseData(validPassportsData)) == 4

with open('inputData.txt') as f:
    fileContent = f.read()
    data = parseData(fileContent)
    assert controlPassports(data) == 198
