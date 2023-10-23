use std::{fs, str::FromStr};

#[test]
fn p1_test() {
    assert_eq!(569999, p1(&fs::read_to_string("d6.txt").unwrap()))
}

#[test]
fn p2_test() {
    assert_eq!(17836115, p2(&fs::read_to_string("d6.txt").unwrap()))
}

#[derive(Debug, PartialEq)]
struct Command {
    startX: i32,
    startY: i32,
    endX: i32,
    endY: i32,
    i: Instruction,
}

#[derive(Debug, PartialEq)]
enum Instruction {
    On,
    Off,
    Toggle,
}

impl FromStr for Command {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parse_pair = |x: &str| -> Option<(i32,i32)> {
            let p = x.split(",").collect::<Vec<_>>();
            match p[..] {
                [a,b] => match (a.parse::<i32>(), b.parse::<i32>()) {
                    (Ok(x), Ok(y)) => Some((x,y)),
                    _ => None,
                }
                _ => None,
            }
        };

        let parts = s.split(" through ").collect::<Vec<_>>();
        if parts.len() != 2 {
            return Err(format!("invalid str: {s}"));
        }

        let second_pair = parse_pair(parts[1]).ok_or(format!("invalid range: {s}"))?;
        
        let first = parts[0];
        
        let instr = if first.starts_with("turn on ") {
            first.strip_prefix("turn on ").and_then(parse_pair).map(|v| Command{ startX: v.0, startY: v.1, endX: second_pair.0, endY: second_pair.1, i: Instruction::On })
        } else if first.starts_with("toggle ") {
            first.strip_prefix("toggle ").and_then(parse_pair).map(|v| Command { startX: v.0, startY: v.1, endX: second_pair.0, endY: second_pair.1, i: Instruction::Toggle })
        } else if first.starts_with("turn off ") {
            first.strip_prefix("turn off ").and_then(parse_pair).map(|v| Command { startX: v.0, startY: v.1, endX: second_pair.0, endY: second_pair.1, i: Instruction::Off })
        } else {
            None
        };

        instr.ok_or(format!("Invalid cmd: {s}"))
    }
}

struct Lights {
    data: Vec<Vec<u8>>
}

impl Lights {
    fn new() -> Self {
        Self { data: vec![vec![0;1000]; 1000] }
    }

    fn process_p1(&mut self, cmd: &Command) {
        for r in cmd.startY..=cmd.endY {
            for c in cmd.startX..=cmd.endX {
                match cmd.i {
                    Instruction::On => self.data[r as usize][c as usize] = 1,
                    Instruction::Off => self.data[r as usize][c as usize] = 0,
                    Instruction::Toggle => if self.data[r as usize][c as usize] == 0 {
                        self.data[r as usize][c as usize] = 1
                    } else {
                        self.data[r as usize][c as usize] = 0
                    }
                }
            }
        }
    }

    fn process_p2(&mut self, cmd: &Command) {
        for r in cmd.startY..=cmd.endY {
            for c in cmd.startX..=cmd.endX {
                match cmd.i {
                    Instruction::On => self.data[r as usize][c as usize] += 1,
                    Instruction::Off => if self.data[r as usize][c as usize] > 0 {
                        self.data[r as usize][c as usize] -= 1;
                    },
                    Instruction::Toggle => self.data[r as usize][c as usize] += 2,
                }
            }
        }
    }

    fn how_many_lit(&self) -> i32 {
        self.data.iter()
            .flat_map(|rows| rows)
            .map(|v| *v as i32)
            .sum()
    }
}

fn p1(s: &str) -> i32 {
    let mut l = Lights::new();
    let r = s.lines().map(|line| line.parse::<Command>()).collect::<Result<Vec<_>, _>>().unwrap();
    r.iter().for_each(|c| l.process_p1(c));

    l.how_many_lit()
}

fn p2(s: &str) -> i32 {
    let mut l = Lights::new();
    let r = s.lines().map(|line| line.parse::<Command>()).collect::<Result<Vec<_>, _>>().unwrap();
    r.iter().for_each(|c| l.process_p2(c));

    l.how_many_lit()
}