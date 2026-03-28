use std::{collections::HashMap, str::FromStr};

#[derive(Debug, PartialEq)]
struct Descriptor {
    id: i32,
    location: Location,
    size: Size,
}

impl Descriptor {
    fn iter(&self) -> impl Iterator<Item = Location> {
        let x0 = self.location.x;
        let y0 = self.location.y;
        let w = self.size.width;
        let h = self.size.height;

        (y0..y0 + h).flat_map(move |y| {
            (x0..x0 + w).map(move |x| Location { x, y })
        })
    }
}

#[derive(Debug, PartialEq, Eq, Hash, Copy, Clone)]
struct Location {
    x: i32,
    y: i32,
}

#[derive(Debug, PartialEq)]
struct Size {
    width: i32,
    height: i32,
}

impl FromStr for Descriptor {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parts: Vec<&str> = s.split(&['#', '@', ',', ':', 'x', ' '][..])
                                   .filter(|x| !x.is_empty())
                                   .collect();

        let id: i32 = parts.get(0).ok_or(format!("missing id in {}", s))?.parse().map_err(|err|format!("failed parsing id for {}: {}", s, err))?;
        let x: i32 = parts.get(1).ok_or(format!("missing x in {}", s))?.parse().map_err(|err|format!("failed parsing x for {}: {}", s, err))?;
        let y: i32 = parts.get(2).ok_or(format!("missing y in {}", s))?.parse().map_err(|err|format!("failed parsing y for {}: {}", s, err))?;
        let width: i32 = parts.get(3).ok_or(format!("missing width in {}", s))?.parse().map_err(|err|format!("failed parsing width for {}: {}", s, err))?;
        let height: i32 = parts.get(4).ok_or(format!("missing height in {}", s))?.parse().map_err(|err|format!("failed parsing height for {}: {}", s, err))?;

        Ok(Descriptor{ id, location: Location { x, y }, size: Size { width, height } })
    }
}

fn build_fabric(descs: &Vec<Descriptor>) -> HashMap<Location, i32> {
    let mut fabric: HashMap<Location, i32> = HashMap::new();

    for v in descs {
        for loc in v.iter() {
            fabric
                .entry(loc.clone())
                .and_modify(|vc| { *vc +=1; } )
                .or_insert(1);
        }
    }

    fabric
}

fn p1(input: &str) -> Result<i32, String> {
    let descs = input.lines().map(&str::parse::<Descriptor>).collect::<Result<Vec<_>, _>>()?;
    let fabric = build_fabric(&descs);
    Ok(fabric.into_values().filter(|&v| v > 1).count() as i32)
}

fn p2(input: &str) -> Result<i32, String> {
    let descs = input.lines().map(&str::parse::<Descriptor>).collect::<Result<Vec<_>, _>>()?;
    let fabric = build_fabric(&descs);
    
    for v in descs {
        let mut allOnes = true;
        for w in v.location.x..v.location.x + v.size.width{
            for h in v.location.y..v.location.y + v.size.height{
                match fabric.get(&Location { x: w, y: h }) {
                    Some(&a) if a != 1 => {
                        allOnes = false;
                        break;
                    }
                    _ => continue,
                }
            }
        }
        if allOnes {
            return Ok(v.id);
        }
    }
    Err("not found".to_string())
}

#[cfg(test)]
mod test {
    use std::fs;

    use super::*;
    
    #[test]
    fn parser_test() {
        assert_eq!(Ok(Descriptor{ id: 1, location: Location { x: 1, y: 3 }, size: Size { width: 4, height: 4 }}), "#1 @ 1,3: 4x4".parse::<Descriptor>());
        assert_eq!(Ok(Descriptor{ id: 2, location: Location { x: 3, y: 1 }, size: Size { width: 4, height: 4 }}), "#2 @ 3,1: 4x4".parse::<Descriptor>());
        assert_eq!(Ok(Descriptor{ id: 3, location: Location { x: 5, y: 5 }, size: Size { width: 2, height: 2 }}), "#3 @ 5,5: 2x2".parse::<Descriptor>());
    }

    #[test]
    fn p1_ex() {
        let data = "#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2";
        assert_eq!(Ok(4), p1(data));
    }
    
    #[test]
    fn p1_test() {
       assert_eq!(Ok(116920), p1(&fs::read_to_string("d3.txt").unwrap()));
    }

    #[test]
    fn p2_ex() {
        let data = "#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2";
        assert_eq!(Ok(3), p2(data));
    }

    #[test]
    fn p2_test() {
       assert_eq!(Ok(382), p2(&fs::read_to_string("d3.txt").unwrap()));
    }
}