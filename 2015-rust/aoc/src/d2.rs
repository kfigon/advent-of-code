use std::{fs, str::FromStr};

#[test]
fn p1_test() {
    assert_eq!(1606483, p1(&fs::read_to_string("d2.txt").unwrap()))
}

#[test]
fn p2_test() {
    assert_eq!(3842356, p2(&fs::read_to_string("d2.txt").unwrap()))
}

struct Dimentions {
    l: usize,
    w: usize,
    h: usize,
}

impl Dimentions {
    fn area(&self) -> usize {
        self.sides().iter().map(|v| v*2).sum()
    }

    fn smallest_size(&self) -> usize {
        *self.sides().iter().min().unwrap()
    }

    fn sides(&self) -> [usize; 3] {
        [self.l*self.w, self.w*self.h, self.h*self.l]
    }

    fn ribbon(&self) -> usize {
        let mut vals = [self.w, self.h, self.l];
        vals.sort();
        let smallest = &vals[0..2];
        2* smallest[0] + 2*smallest[1] + self.h*self.w*self.l 
    }
}

impl FromStr for Dimentions {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let chunks = s.split("x")
            .map(|v| v.parse::<usize>())
            .collect::<Result<Vec<_>, _>>().map_err(|e| format!("{}: {}", e.to_string(), s))?;

        match chunks[..] {
            [l, w,h] => Ok(Self { l, w, h }),
            _ => Err(format!("invalid data: {:?}", chunks))
        }
    }
}

fn parse(s: &str) -> Result<Vec<Dimentions>, String> {
    s.lines().map(&str::parse).collect()
}

fn p1(s: &str) -> usize {
    parse(s).unwrap().iter().map(|d| d.area() + d.smallest_size()).sum()
}

fn p2(s: &str) -> usize {
    parse(s).unwrap().iter().map(&Dimentions::ribbon).sum()
}