from typing import Dict

class Operations:
    LOAD = 0
    AND = 1
    OR = 2
    LSHIFT = 3
    RSHIFT = 4
    NOT = 5

class Processor:
    def __init__(self):
        self.registers: Dict[str, int] = {}
    def readRegister(self, regName: str) -> int:
        if regName not in self.registers:
            raise ValueError(f'Invalid register {regName}')
        return self.registers[regName]

    def registersName(self):
        return self.registers.keys()
    def doOperation(self, operation: int, arg1: int, arg2: int, targetRegister: str):
        if operation == Operations.AND:
            self.doLoadOperation(arg1 & arg2, targetRegister)
        elif operation == Operations.OR:
            self.doLoadOperation(arg1 | arg2, targetRegister)
        elif operation == Operations.LSHIFT:
            self.doLoadOperation(arg1 << arg2, targetRegister)
        elif operation == Operations.RSHIFT:
            self.doLoadOperation(arg1 >> arg2, targetRegister)
    def doLoadOperation(self, val: int, targetRegister: str):
        self.registers[targetRegister] = (val & 0xFFFF)
    def doNot(self, sourceValue: int, targetRegister: str):
        self.doLoadOperation(~sourceValue, targetRegister)

