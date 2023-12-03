use std::{str::FromStr, fs, collections::{HashMap, HashSet}};

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

#[test]
fn p2_test() {
    assert_eq!(72246648, p2(&fs::read_to_string("d3.txt").unwrap()));
}

struct Grid(Vec<Vec<char>>);
impl Grid {
    fn neighbors(&self, row: usize, col: usize) -> [Option<(Coord, char)>; 8] {
        let get = |r: Option<usize>, c: Option<usize>| -> Option<(Coord,char)> {
            self.get(r, c).map(|v| (Coord{r: r.unwrap(),c:c.unwrap()}, v))
        };

        [
            get(row.checked_sub(1), col.checked_sub(1)), get(row.checked_sub(1), Some(col)), get(row.checked_sub(1), col.checked_add(1)),   
            get(Some(row), col.checked_sub(1)),                         get(Some(row), col.checked_add(1)),
            get(row.checked_add(1), col.checked_sub(1)), get(row.checked_add(1), Some(col)), get(row.checked_add(1), col.checked_add(1))
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

#[derive(Debug, PartialEq, Eq, Hash, Clone)]
struct Coord {
    r: usize,
    c: usize
}

fn p1(s: &str) -> usize {
    let (_, symbols_with_numbers_around) = process(s);

    symbols_with_numbers_around.into_iter()
        .map(|v| v.1)
        .filter(|v| !v.is_empty())
        .flat_map(|v| v)
        .sum()
}

fn process(s: &str) -> (Grid, HashMap<Coord, HashSet<usize>>) {
    let g: Grid = s.parse().expect("invalid input");
    let mut symbols_with_numbers_around: HashMap<Coord, HashSet<usize>> = HashMap::new();

    for r in 0..g.rows() {
        let mut dig = String::new();
        let mut symbols_near_dig: HashSet<Coord> = HashSet::new();

        for c in 0..g.cols() {
            let character = g.get(Some(r), Some(c)).unwrap();

            if character.is_ascii_digit() {
                dig.push(character);

                g.neighbors(r, c).iter()
                .filter_map(|nei| match nei {
                    Some((coord, c)) if *c != '.' && !c.is_ascii_digit() => Some((coord.clone(), *c)),
                    _ => None,
                })
                .for_each(|v| { symbols_near_dig.insert(v.0); });
            
            } else if !dig.is_empty() {
                symbols_near_dig.iter()
                    .for_each(|s| {
                        let nums = symbols_with_numbers_around.entry(s.clone()).or_default();
                        nums.insert(dig.parse().unwrap());
                    });
                    symbols_near_dig.clear();
                dig.clear();
            }
        }
        
        if !dig.is_empty() {
            symbols_near_dig.into_iter()
                .for_each(|s| {
                    let nums = symbols_with_numbers_around.entry(s).or_default();
                    nums.insert(dig.parse().unwrap());
                });
        }
    }

    (g, symbols_with_numbers_around)
}

fn p2(s: &str) -> usize {
    let (grid, symbols_with_numbers_around) = process(s);
    
    symbols_with_numbers_around.into_iter()
        .filter(|v| match grid.get(Some(v.0.r), Some(v.0.c)) {
            Some('*') => true,
            _ => false,
        })
        .map(|v| v.1)
        .filter(|v| v.len() == 2)
        .map(|v| v.iter().product::<usize>())
        .sum()
}