use std::{collections::HashMap, fs};

const EXAMPLE: &'static str = ".#.#.#
...##.
#....#
..#...
#.#..#
####..";

#[test]
fn p1_ex() {
    assert_eq!(4, p1(EXAMPLE, 4));
}

#[test]
fn p2_ex() {
    assert_eq!(17, p2(EXAMPLE, 5));
}

#[test]
fn p1_test() {
    assert_eq!(821, p1(&fs::read_to_string("d18.txt").unwrap(), 100));
}

#[test]
fn p2_test() {
    assert_eq!(886, p2(&fs::read_to_string("d18.txt").unwrap(), 100));
}

struct Grid<T: ElementAccessor>{
    els: Vec<Vec<u8>>,
    accessor: T,
}

fn parse<T: ElementAccessor>(value: &str, accessor: T) -> Grid<T> {
    Grid { 
        els: value.lines()
            .map(|v| v.as_bytes().to_vec())
            .collect(), 
        accessor: accessor 
    }
}


impl<T: ElementAccessor> Grid<T> {
    // A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
    // A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
    fn next(&mut self) {
        let mut changes: HashMap<(usize, usize), u8> = HashMap::new();
        for (r, rows) in self.els.iter().enumerate() {
            for (c, _) in rows.iter().enumerate() {
                let nei = self.lit_neighbors(r, c);
                let el = self.get_el(Some(r), Some(c));

                match(el, nei) {
                    (Some(b'#'), n) if n != 2 && n != 3 => changes.insert((r,c), b'.'),
                    (Some(b'.'), n) if n == 3 => changes.insert((r,c), b'#'),
                    _ => None,
                };
            }
        }

        for (key, el) in changes {
            self.accessor.set_el(&mut self.els, key.0, key.1, el);
        }
    }

    fn get_el(&self, r: Option<usize>, c: Option<usize>) -> Option<u8> {
        self.accessor.get_el(&self.els, r,c)
    }

    fn lit_neighbors(&self, r: usize, c: usize) -> usize {
        let top = r.checked_sub(1);
        let down = r.checked_add(1);
        let left = c.checked_sub(1);
        let right = c.checked_add(1);
        let nei = [
            self.get_el(top, left), self.get_el(top, Some(c)), self.get_el(top, right), 
            self.get_el(Some(r), left),                             self.get_el(Some(r), right), 
            self.get_el(down, left), self.get_el(down, Some(c)), self.get_el(down, right), 
        ];

        nei.iter().flatten().filter(|&&v| v == b'#').count()
    }

    fn count_on(&self) -> usize {
        self.els.iter()
            .flat_map(|v| v)
            .filter(|&&v| v == b'#')
            .count()
    }
}

trait ElementAccessor {
    fn get_el(&self, v: &Vec<Vec<u8>>, r: Option<usize>, c: Option<usize>) -> Option<u8>;
    fn set_el(&self, v: &mut Vec<Vec<u8>>, r: usize, c: usize, value: u8);
}

struct DefaultAccesor;
impl ElementAccessor for DefaultAccesor {
    fn get_el(&self, v: &Vec<Vec<u8>>, r: Option<usize>, c: Option<usize>) -> Option<u8> {
        match (r, c) {
            (Some(rx), Some(cx)) => v.get(rx).and_then(|cols| cols.get(cx).cloned()),
            _ => None,
        }
    }

    fn set_el(&self, v: &mut Vec<Vec<u8>>, r: usize, c: usize, value: u8) {
        v[r][c] = value;
    }
}
struct StuckAccessor;
impl ElementAccessor for StuckAccessor {
    fn get_el(&self, v: &Vec<Vec<u8>>, r: Option<usize>, c: Option<usize>) -> Option<u8> {
        let num_rows = v.len()-1;
        let num_cols = v[0].len()-1;
        match (r, c) {
            (Some(0), Some(0)) => Some(b'#'),
            (Some(0), Some(cols)) if cols == num_cols => Some(b'#'),
            (Some(rows), Some(0)) if rows == num_rows => Some(b'#'),
            (Some(rows), Some(cols)) if rows == num_rows && cols == num_cols => Some(b'#'),
            (Some(rx), Some(cx)) => v.get(rx).and_then(|cols| cols.get(cx).cloned()),
            _ => None,
        }
    }

    fn set_el(&self, v: &mut Vec<Vec<u8>>, r: usize, c: usize, value: u8) {
        let num_rows = v.len()-1;
        let num_cols = v[0].len()-1;

        v[r][c] = match (r,c) {
            (0, 0) => b'#',
            (0, cols) if cols == num_cols => b'#',
            (rows, 0) if rows == num_rows => b'#',
            (rows, cols) if rows == num_rows && cols == num_cols => b'#',
            _ => value,
        };
    }
}

fn p1(s: &str, steps: usize) -> usize {
    solve(parse(s, DefaultAccesor{}), steps)
}

fn p2(s: &str, steps: usize) -> usize {
    solve(parse(s, StuckAccessor{}), steps)
}

fn solve<T: ElementAccessor>(mut g: Grid<T>, steps: usize) -> usize {
    for _ in 0..steps {
        g.next();
    }
    g.count_on()
}