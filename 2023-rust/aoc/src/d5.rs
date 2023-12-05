use std::{collections::HashMap, str::FromStr, fs};

const EXAMPLE: &'static str = "seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4";

#[test]
fn p1_ex() {
    assert_eq!(35, p1(EXAMPLE));
}

#[test]
fn p1_test() {
    todo!("fix ranges, file data is too big");
    assert_eq!(35, p1(&fs::read_to_string("d5.txt").unwrap()));
}

#[derive(Debug)]
struct Map(HashMap<usize, usize>);

impl TryFrom<&[&str]> for Map {
    type Error = String;

    fn try_from(lines: &[&str]) -> Result<Self, Self::Error> {
        let mut m = HashMap::new();
        
        for line in lines {
            let vals = line.split_ascii_whitespace()
                .map(|v| v.parse())
                .collect::<Result<Vec<usize>, _>>()
                .map_err(|err| format!("{err}: no number found {line}"))?;

            match vals[..] {
                [dest_start, source_start, range] => {
                    for i in 0..range {
                        m.insert(source_start+i, dest_start+i);
                    }
                },
                _ => return Err(format!("invalid line: {line}")),
            }
        }

        Ok(Map(m))
    }
}

impl Map {
    fn get(&self, v: usize) -> usize {
        *self.0.get(&v).unwrap_or(&v)
    }
}

#[derive(Debug)]
struct Maps {
    seeds: Vec<usize>,
    seed_to_soil: Map,
    soil_to_fertilizer: Map,
    fert_to_water: Map,
    water_to_ligt: Map,
    light_to_temp: Map,
    temp_to_humid: Map,
    humid_to_location: Map,
}

impl FromStr for Maps {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let s = s.replace("\r", "");
        let parts = s.split("\n\n").collect::<Vec<_>>();

        if parts.len() != 8 {
            return Err(format!("invalid data format, expected 8 parts, got {}", parts.len()));
        }

        let seeds = parts[0]
            .split_once(": ")
            .ok_or(format!("cant find seeds"))?
            .1.split_whitespace()
            .map(|v| v.parse())
            .collect::<Result<Vec<usize>, _>>()
            .map_err(|e| format!("{e}: number error when parsing seeds"))?;

        let mut seed_to_soil: Option<Map> = None;
        let mut soil_to_fertilizer: Option<Map> = None;
        let mut fert_to_water: Option<Map> = None;
        let mut water_to_ligt: Option<Map> = None;
        let mut light_to_temp: Option<Map> = None;
        let mut temp_to_humid: Option<Map> = None;
        let mut humid_to_location: Option<Map> = None;


        for chunks in &parts[1..] {
            let lines = chunks.lines().collect::<Vec<_>>();
            match &lines[..] {
                ["seed-to-soil map:", rest @ ..] => seed_to_soil = Some(rest.try_into()?),
                ["soil-to-fertilizer map:", rest @ ..] => soil_to_fertilizer = Some(rest.try_into()?),
                ["fertilizer-to-water map:", rest @ ..] => fert_to_water = Some(rest.try_into()?),
                ["water-to-light map:", rest @ ..] => water_to_ligt = Some(rest.try_into()?),
                ["light-to-temperature map:", rest @ ..] => light_to_temp = Some(rest.try_into()?),
                ["temperature-to-humidity map:", rest @ ..] => temp_to_humid = Some(rest.try_into()?),
                ["humidity-to-location map:", rest @ ..] => humid_to_location = Some(rest.try_into()?),
                _ => return Err(format!("invalid mapping: {}", lines[0])),
            }
        }


        match (seed_to_soil, soil_to_fertilizer, fert_to_water, water_to_ligt, light_to_temp, temp_to_humid, humid_to_location) {
            (Some(seed_to_soil), Some(soil_to_fertilizer), Some(fert_to_water), Some(water_to_ligt), Some(light_to_temp), Some(temp_to_humid), Some(humid_to_location)) => Ok(Maps { seeds, seed_to_soil, soil_to_fertilizer, fert_to_water, water_to_ligt, light_to_temp, temp_to_humid, humid_to_location }),
            _ => Err(format!("not all elements are found")),
        }
    }
}

fn p1(s: &str) -> usize {
    let m: Maps = s.parse().unwrap();
    m.seeds.into_iter()
        .map(|seed| {
            let soil = m.seed_to_soil.get(seed);
            let fert = m.soil_to_fertilizer.get(soil);
            let water = m.fert_to_water.get(fert);
            let light = m.water_to_ligt.get(water);
            let temp = m.light_to_temp.get(light);
            let humid = m.temp_to_humid.get(temp);
            let loc = m.humid_to_location.get(humid);

            loc
        })
        .min().unwrap()
}