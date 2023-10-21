use std::{fs, collections::HashSet};

enum Direction {
    North,
    South,
    East,
    West,
}

fn parse(c: char) -> Direction {
    match c {
        '^' => Direction::North,
        '>' => Direction::East,
        'v' => Direction::South,
        '<' => Direction::West,
        _ => unreachable!()
    }
}

#[derive(Debug, PartialEq, Clone, Copy, Hash, Eq)]
struct Coord {
    x: i32,
    y: i32,
}


#[test]
fn p1_test() {
    assert_eq!(2592, p1(&fs::read_to_string("d3.txt").unwrap()))
}

#[test]
fn p2_test() {
    assert_eq!(2360, p2(&fs::read_to_string("d3.txt").unwrap()))
}

#[test]
fn p2_ex() {
    assert_eq!(11, p2("^v^v^v^v^v"))
}

struct Delivery {
    last: Coord,
    houses: HashSet<Coord>
}

impl Delivery {
    fn new() -> Self {
        Self { last: Coord { x: 0, y: 0 }, houses: HashSet::new() }
    }
    fn visit(&mut self, c: &Coord) {
        self.houses.insert(*c);
        self.last = *c;
    }

    fn mv(&mut self, d: Direction) {
        let new = match d {
            Direction::North => Coord { x: self.last.x, y: self.last.y-1 },
            Direction::South => Coord { x: self.last.x, y: self.last.y+1 },
            Direction::East => Coord { x: self.last.x+1, y: self.last.y },
            Direction::West => Coord { x: self.last.x-1, y: self.last.y },
        }; 
        self.visit(&new);
    }
}

fn p1(s: &str) -> usize {
    let mut d = Delivery::new();
    d.visit(&Coord { x: 0, y: 0 });
    s.chars().map(parse).for_each(|step| d.mv(step));

    d.houses.len()
}

fn p2(s: &str) -> usize {
    let mut santa = Delivery::new();
    let mut robo_santa = Delivery::new();
    santa.visit(&Coord { x: 0, y: 0 });
    robo_santa.visit(&Coord { x: 0, y: 0 });
    
    s.chars()
        .map(parse)
        .enumerate()
        .for_each(|(i, step)| if i % 2 == 0 {
                santa.mv(step);
            } else {
                robo_santa.mv(step)
            }
        );
    
    santa.houses.extend(robo_santa.houses);
    santa.houses.len()
}