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

struct Grid(Vec<Vec<u8>>);
impl TryFrom<&str> for Grid {
    type Error = String;

    fn try_from(value: &str) -> Result<Self, Self::Error> {
        Ok(Self(value.lines()
            .map(|v| v.as_bytes().to_vec())
            .collect())
        )
    }
}

impl Grid {
    // A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
    // A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
    fn next(&mut self) {
        let mut changes: HashMap<(usize, usize), u8> = HashMap::new();
        for (r, rows) in self.0.iter().enumerate() {
            for (c, &el) in rows.iter().enumerate() {
                let nei = self.lit_neighbors(r, c);

                match(el, nei) {
                    (b'#', n) if n != 2 && n != 3 => changes.insert((r,c), b'.'),
                    (b'.', n) if n == 3 => changes.insert((r,c), b'#'),
                    _ => None,
                };
            }
        }

        for (key, el) in changes {
            self.0[key.0][key.1] = el;
        }
    }

    fn lit_neighbors(&self, r: usize, c: usize) -> usize {
        let get = |ro: Option<usize>, co: Option<usize>| match (ro, co) {
            (Some(rx), Some(cx)) => self.0.get(rx).and_then(|cols| cols.get(cx).cloned()),
            _ => None,
        };

        let nei = [
            get(r.checked_sub(1), c.checked_sub(1)), get(r.checked_sub(1), Some(c)), get(r.checked_sub(1), c.checked_add(1)), 
            get(Some(r), c.checked_sub(1)),                                                 get(Some(r), c.checked_add(1)), 
            get(r.checked_add(1), c.checked_sub(1)), get(r.checked_add(1), Some(c)), get(r.checked_add(1), c.checked_add(1)), 
        ];

        nei.iter().flatten().filter(|&&v| v == b'#').count()
    }

    fn count_on(&self) -> usize {
        self.0.iter()
            .flat_map(|v| v)
            .filter(|&&v| v == b'#')
            .count()
    }
}

fn p1(s: &str, steps: usize) -> usize {
    let mut g: Grid = s.try_into().unwrap();

    for _ in 0..steps {
        g.next();
    }
    g.count_on()
}

fn p2(s: &str, steps: usize) -> usize {
    todo!()
}