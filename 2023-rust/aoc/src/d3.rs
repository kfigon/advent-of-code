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

#[test]
fn p2_ex() {
    assert_eq!(467835, p2(EXAMPLE));
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

    fn nei_id_to_coord(&self, r: usize, c: usize, idx: usize) -> Option<(usize, usize)> {
        let coord_within_range = |r: Option<usize>, c: Option<usize>| -> Option<(usize, usize)> { 
                match (r, c) {
                (Some(r), Some(c)) if r >= 0 && r < self.rows() && c >= 0 && c < self.cols() => Some((r,c)),
                _ => None,
            }
        };

        match idx {
            0 => coord_within_range(r.checked_sub(1), c.checked_sub(1)),
            1 => coord_within_range(r.checked_sub(1), Some(c)),
            2 => coord_within_range(r.checked_sub(1), c.checked_add(1)),

            3 => coord_within_range(Some(r), c.checked_sub(1)),
            4 => coord_within_range(Some(r), c.checked_add(1)),
            
            5 => coord_within_range(r.checked_add(1), c.checked_sub(1)),
            6 => coord_within_range(r.checked_add(1), Some(c)),
            7 => coord_within_range(r.checked_add(1), c.checked_add(1)),
            _ => None
        }
    }

    fn get_number_around_coord(&self, r: usize, c: usize) -> Option<usize> {
        let is_digit = |c: usize| -> bool {
            self.get(Some(r), Some(c)).filter(|v| v.is_ascii_digit()).is_some()
        };

        let mut starting_idx = Some(c);
        while let Some(idx) = starting_idx {
            if !is_digit(idx) {
                break;
            }
            starting_idx = idx.checked_sub(1);
        }

        let mut end_idx = Some(c);
        while let Some(idx) = end_idx {
            if idx < self.cols() && !is_digit(idx) {
                break;
            }
            end_idx = idx.checked_add(1);
        }
        let (num_start, num_end) = match (starting_idx, end_idx) {
            (None, None) => (0, self.cols()-1),
            (None, Some(e)) => (0, e-1),
            (Some(s), None) => (s+1, self.cols()),
            (Some(s), Some(e)) => (s+1, e-1),
        };
        let mut dig = String::new();
        for c in num_start..=num_end {
            let v = match self.get(Some(r), Some(c)) {
                None => return None,
                Some(v) => v,
            };
            dig.push(v);
        }
        dig.parse().ok()
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

fn p2(s: &str) -> usize {
    let g: Grid = s.parse().expect("invalid input");
    let mut nums = vec![];

    for r in 0..g.rows() {
        for c in 0..g.cols() {
            let character = g.get(Some(r), Some(c)).unwrap();
            if character != '*' {
                continue;
            }

            let nei = g.neighbors(r, c);
            let numeric_nei = nei.iter().enumerate()
                .filter_map(|(idx, c)| match c {
                    Some(v) if v.is_ascii_digit() => Some(idx),
                    _ => None,
                })
                .collect::<Vec<usize>>();

                // this can grab same num multiple times
            if numeric_nei.len() != 2 {
                continue;
            }
            
            for n in numeric_nei {
                let mut numz = vec![];
                match g.nei_id_to_coord(r, c, n) {
                    None => continue,
                    Some((r,c)) => match g.get_number_around_coord(r,c) {
                        None => continue,
                        Some(v) => numz.push(v),
                    }
                }
                nums.push(numz.iter().product());
            }

        }        
    }

    nums.iter().sum()
}