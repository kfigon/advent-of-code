use std::{str::FromStr, fs};

static EXAMPLE: &'static str = "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..";

#[test]
fn p1_ex() {
    assert_eq!(4361, p1(EXAMPLE));
}

#[test]
fn p1_test() {
    assert_eq!(498559, p1(&fs::read_to_string("d3.txt").unwrap()));
}

struct Grid(Vec<Vec<char>>);
impl Grid {
    fn neighbors(&self, row: usize, col: usize) -> [Option<char>; 8] {
        [
            self.get(row.checked_sub(1), col.checked_sub(1)), self.get(row.checked_sub(1), Some(col)), self.get(row.checked_sub(1), col.checked_add(1)),   
            self.get(Some(row), col.checked_sub(1)),                         self.get(Some(row), col.checked_add(1)),
            self.get(row.checked_add(1), col.checked_sub(1)), self.get(row.checked_add(1), Some(col)), self.get(row.checked_add(1), col.checked_add(1))
        ]
    }

    fn get(&self, r: Option<usize>, c: Option<usize>) -> Option<char> {
        match (r,c) {
            (Some(r), Some(c)) => self.0.get(r).and_then(|row| row.get(c).copied()),
            _ => None,
        }
    }

    fn rows(&self) -> usize {
        self.0.len()
    }

    fn cols(&self) -> usize {
        self.0.get(0).map(|row| row.len()).unwrap_or_default()
    }
}
impl FromStr for Grid {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let d = Grid(s.lines()
                .map(|line| line.chars().collect::<Vec<char>>())
                .collect::<Vec<_>>());

        for (id,r) in d.0.iter().enumerate() {
            if r.len() != d.cols() {
                return Err(format!("not all rows have equal columns, mismatch at row num {id}"));
            }
        }
        Ok(d)
    }
}

fn p1(s: &str) -> usize {
    let g: Grid = s.parse().expect("invalid input");
    let mut nums = vec![];

    for r in 0..g.rows() {
        let mut dig = String::new();
        let mut has_symbol_around = false;
        for c in 0..g.cols() {
            let character = g.get(Some(r), Some(c)).unwrap();

            if character.is_ascii_digit() {
                dig.push(character);

                if !has_symbol_around {
                    has_symbol_around = g.neighbors(r, c).iter().any(|nei| match nei {
                        None => false,
                        Some(c) => *c != '.' && !c.is_ascii_digit() ,
                    });
                }
            } else {
                if !dig.is_empty() && has_symbol_around {
                    nums.push(dig.parse().unwrap());
                }

                dig.clear();
                has_symbol_around = false;
            }
        }
        
        if !dig.is_empty() && has_symbol_around {
            nums.push(dig.parse().unwrap());
        }
    }

    nums.iter().sum()
}