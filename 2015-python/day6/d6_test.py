from d6 import *
import unittest

class TestRectange(unittest.TestCase):
    def testSize(self):
        tt = [
            {'in': Rectangle(Coordinate(0,0), Coordinate(0,1)), 'exp': 2},
            {'in': Rectangle(Coordinate(0,0), Coordinate(999,999)), 'exp': 1000000},
            {'in': Rectangle(Coordinate(0,0), Coordinate(2,2)), 'exp': 9},
            {'in': Rectangle(Coordinate(0,0), Coordinate(999,0)), 'exp': 1000},
            {'in': Rectangle(Coordinate(499,499), Coordinate(500,500)), 'exp': 4}
        ]
        for tc in tt:
            with self.subTest(str(tc['in'])):
                self.assertEqual(tc['exp'], tc['in'].size())

class TestParser(unittest.TestCase):
    def testParser(self):
        tt = [
            {'in': 'turn on 0,0 through 999,999', 'exp': (Operation.TURN_ON, Rectangle(Coordinate(0,0),Coordinate(999,999)))},
            {'in': 'toggle 0,0 through 999,0', 'exp': (Operation.TOGGLE, Rectangle(Coordinate(0,0),Coordinate(999,0)))},
            {'in': 'turn off 499,499 through 500,500', 'exp': (Operation.TURN_DOWN, Rectangle(Coordinate(499,499),Coordinate(500,500)))}
        ]
        for tc in tt:
            with self.subTest(str(tc['in'])):
                operation, rectangle = parseCommand(tc['in'])
                expectedOp, expectedRect = tc['exp']
                self.assertEqual(expectedOp, operation)
                self.assertEqual(expectedRect, rectangle)

    def testInvalidInput(self):
        self.assertRaises(ValueError, parseCommand, None)
        self.assertRaises(ValueError, parseCommand, "")
        self.assertRaises(ValueError, parseCommand, "some random")
        self.assertRaises(ValueError, parseCommand, "toggle 692,139 asd 698,880")
        self.assertRaises(ValueError, parseCommand, "turn ofx 499,499 through 500,500")
        self.assertRaises(ValueError, parseCommand, "turn ofx 4991499 through 500,500")

class TestGrid(unittest.TestCase):
    def test1(self):
        g = Grid()
        g.add(*parseCommand('turn on 0,0 through 999,999'))
        g.add(*parseCommand('toggle 0,0 through 999,0'))
        g.add(*parseCommand('turn off 499,499 through 500,500'))
        self.assertEqual(998_996, g.applyChangeLog())

    def test2(self):
        g = Grid()
        g.add(*parseCommand('turn on 0,0 through 0,0'))
        self.assertEqual(1, g.applyChangeLog2())

    def test3(self):
        g = Grid()
        g.add(*parseCommand('toggle 0,0 through 999,999'))
        self.assertEqual(2000000, g.applyChangeLog2())

    def test4(self):
        g = Grid()
        g.add(*parseCommand('turn on 0,0 through 0,0'))
        g.add(*parseCommand('toggle 0,0 through 0,0'))
        self.assertEqual(3, g.applyChangeLog2())

    def test5(self):
        g = Grid()
        g.add(*parseCommand('turn on 0,0 through 0,0'))
        g.add(*parseCommand('toggle 0,0 through 999,999'))
        self.assertEqual(2000001, g.applyChangeLog2())
    
    def test6(self):
        g = Grid()
        g.add(*parseCommand('turn on 0,0 through 0,0'))
        g.add(*parseCommand('turn off 0,0 through 0,0'))
        g.add(*parseCommand('turn off 0,0 through 0,0'))
        g.add(*parseCommand('turn off 0,0 through 0,0'))
        self.assertEqual(0, g.applyChangeLog2())
    

if __name__ == "__main__":
    unittest.main()