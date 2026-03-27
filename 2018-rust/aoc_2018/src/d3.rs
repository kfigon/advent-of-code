use std::{collections::HashMap, str::FromStr};

#[derive(Debug, PartialEq)]
struct Descriptor {
    id: i32,
    location: Location,
    size: Size,
}

#[derive(Debug, PartialEq, Eq, Hash)]
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
        let parts = s.split_whitespace().collect::<Vec<_>>();
        let id = parts.get(0);
        let loc = parts.get(2);
        let size = parts.get(3);

        let (idRaw, locRaw, sizeRaw) = match (id, loc, size) {
            (Some(idV), Some(locV), Some(sizeV)) => (idV, locV, sizeV),
            _ => return Err(format!("some parts are missing. id: {:?}, loc: {:?}, size: {:?} for string: {}", id, loc, size, s)),
        };

        let id = idRaw.replace("#", "").parse::<i32>().map_err(|err| format!("invalid id: {} for {}", err, s))?;
        let binding = locRaw.replace(":", "");
        let locRaw = binding.split_once(",");
        let location = match locRaw {
            Some((a, b)) => {
                match (a.parse::<i32>(), b.parse::<i32>()) {
                    (Ok(x), Ok(y)) => Location{x,y},
                    _ => return Err(format!("invalid loc - not numbers for {}", s)),
                }
            },
            _ => return Err(format!("invalid loc: {:?} for {}", locRaw, s)),
        };


        let binding = sizeRaw.replace(":", "");
        let sizeRaw = binding.split_once("x");
        let size = match sizeRaw {
            Some((a,b)) => {
                match (a.parse::<i32>(), b.parse::<i32>()) {
                    (Ok(x), Ok(y)) => Size{width: x, height: y},
                    _ => return Err(format!("invalid size -  not numbers for {}", s)),
                }
            },
            _ => return Err(format!("invalid size: {:?} for {}", size, s)),
        };

        Ok(Descriptor {id, location, size})
    }
}

fn build_fabric(descs: &Vec<Descriptor>) -> HashMap<Location, i32> {
    let mut fabric: HashMap<Location, i32> = HashMap::new();

    for v in descs {
        for w in v.location.x..v.location.x + v.size.width{
            for h in v.location.y..v.location.y + v.size.height{
                fabric
                    .entry(Location { x: w, y: h })
                    .and_modify(|vc| { *vc +=1; } )
                    .or_insert(1);
            }
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