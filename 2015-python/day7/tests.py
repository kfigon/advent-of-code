from main import *
import unittest

class TestProcessor(unittest.TestCase):
    def setUp(self):
        self.p = Processor()

    def testEmpty(self):
        self.assertRaises(ValueError, self.p.readRegister, 'x')
        self.assertEqual(0, len(self.p.registersName()))

    def testLoad(self):
        self.p.doLoadOperation(123, 'x')
        self.assertEqual(123, self.p.readRegister('x'))

    def testAnd(self):
        self.p.doLoadOperation(123, 'x')
        self.p.doLoadOperation(456, 'y')
        self.p.doOperation(Operations.AND, self.p.readRegister('x'), self.p.readRegister('y'), 'd')

        self.assertEqual(123, self.p.readRegister('x'))
        self.assertEqual(456, self.p.readRegister('y'))
        self.assertEqual(72, self.p.readRegister('d'))

    def testStres(self):
        self.p.doLoadOperation(123, 'x')
        self.p.doLoadOperation(456, 'y')
        self.p.doOperation(Operations.AND, self.p.readRegister('x'), self.p.readRegister('y'), 'd')
        self.p.doOperation(Operations.OR, self.p.readRegister('x'), self.p.readRegister('y'), 'e')
        self.p.doOperation(Operations.LSHIFT, self.p.readRegister('x'), 2, 'f')
        self.p.doOperation(Operations.RSHIFT, self.p.readRegister('y'), 2, 'g')
        self.p.doNot(self.p.readRegister('x'), 'h')
        self.p.doNot(self.p.readRegister('y'), 'i')

        self.assertEqual(72, self.p.readRegister('d'))
        self.assertEqual(507, self.p.readRegister('e'))
        self.assertEqual(492, self.p.readRegister('f'))
        self.assertEqual(114, self.p.readRegister('g'))
        self.assertEqual(65412, self.p.readRegister('h'))
        self.assertEqual(65079, self.p.readRegister('i'))
        self.assertEqual(123, self.p.readRegister('x'))
        self.assertEqual(456, self.p.readRegister('y'))

class TestParser(unittest.TestCase):
    def test1(self):
        self.fail("todo")

if __name__ == '__main__':
    unittest.main()