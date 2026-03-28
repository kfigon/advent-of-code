use std::{collections::{HashMap, HashSet, hash_map::Entry::{Occupied, Vacant}}, str::FromStr};


#[derive(Debug, PartialEq, Eq, Hash, Copy, Clone)]
struct ID(i32);

#[derive(Debug, PartialEq)]
struct Descriptor {
    id: ID,
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

        Ok(Descriptor{ id: ID(id), location: Location { x, y }, size: Size { width, height } })
    }
}

type FirstID = ID;
fn build_fabric(descs: &Vec<Descriptor>) -> (HashMap<Location, (i32, FirstID)>, HashSet<ID>) {
    let mut overlap_counts: HashMap<Location, (i32, FirstID)> = HashMap::new();

    let mut all_ids = HashSet::<ID>::new();
    let mut overlapped = HashSet::<ID>::new();

    for v in descs {
        all_ids.insert(v.id);
        for loc in v.iter() {
            match overlap_counts.entry(loc) {
                Occupied(mut occupied_entry) => {
                    let (cnt, first_id) = occupied_entry.get_mut();
                    overlapped.insert(v.id);
                    overlapped.insert(*first_id);
                    *cnt+=1;
                },
                Vacant(vacant_entry) => {
                    vacant_entry.insert((1, v.id));
                },
            }
        }
    }

    let non_overlapped = all_ids.difference(&overlapped).copied().collect();
    (overlap_counts, non_overlapped)
}

fn p1(input: &str) -> Result<i32, String> {
    let descs = input.lines().map(&str::parse::<Descriptor>).collect::<Result<Vec<_>, _>>()?;
    let (fabric, _) = build_fabric(&descs);
    Ok(fabric.into_values().filter(|v| v.0 > 1).count() as i32)
}

fn p2(input: &str) -> Result<ID, String> {
    let descs = input.lines().map(&str::parse::<Descriptor>).collect::<Result<Vec<_>, _>>()?;
    let (_, non_overlapping) = build_fabric(&descs);

    match non_overlapping.iter().next() {
        Some(id) => Ok(*id),
        None => Err("Not found".to_string()),
    }
}

#[cfg(test)]
mod test {
    use std::fs;

    use super::*;
    
    #[test]
    fn parser_test() {
        assert_eq!(Ok(Descriptor{ id: ID(1), location: Location { x: 1, y: 3 }, size: Size { width: 4, height: 4 }}), "#1 @ 1,3: 4x4".parse::<Descriptor>());
        assert_eq!(Ok(Descriptor{ id: ID(2), location: Location { x: 3, y: 1 }, size: Size { width: 4, height: 4 }}), "#2 @ 3,1: 4x4".parse::<Descriptor>());
        assert_eq!(Ok(Descriptor{ id: ID(3), location: Location { x: 5, y: 5 }, size: Size { width: 2, height: 2 }}), "#3 @ 5,5: 2x2".parse::<Descriptor>());
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
        assert_eq!(Ok(ID(3)), p2(data));
    }

    #[test]
    fn p2_test() {
       assert_eq!(Ok(ID(382)), p2(&fs::read_to_string("d3.txt").unwrap()));
    }
}