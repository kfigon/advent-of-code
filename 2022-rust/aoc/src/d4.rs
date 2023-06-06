
#[cfg(test)]
mod tests{
    use super::*;

    const EXAMPLE_DATA: &str = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";

    #[test]
    fn example_1() {
        assert_eq!(p1(&EXAMPLE_DATA.to_string()), 2)
    }

    #[test]
    fn example_2() {
        assert_eq!(p2(&EXAMPLE_DATA.to_string()), 4)
    }

    #[test]
    fn file_1() {
        assert_eq!(p1(&std::fs::read_to_string("d4.txt").unwrap()), 562)
    }

    #[test]
    fn file_2() {
        assert_eq!(p2(&std::fs::read_to_string("d4.txt").unwrap()), 924)
    }
}

#[derive(Debug)]
struct Range {
    min: usize,
    max: usize,
}

fn p1(input: &String) -> usize {
    fn fully_contains(a: &Range, b: &Range) -> bool {
        a.min <= b.min && a.max >= b.max
    }

    fn overlap(pair: &(Range, Range)) -> bool {
        fully_contains(&pair.0, &pair.1) || fully_contains(&pair.1, &pair.0)
    }

    input.lines()
    .map(parse_line)
    .filter(overlap)
    .count()
}

fn parse_line(line: &str) -> (Range, Range) {
    let mut pair = line.split(",")
     .map(|v| {
        let mut two = v.split("-")
        .filter_map(|pairs| pairs.parse::<usize>().ok());
        
        match (two.next(), two.next()) {
            (Some(min), Some(max)) => Some(Range{min: min, max: max}),
            _ => None
        }
    });
    
    match (pair.next(), pair.next()) {
        (Some(Some(r1)), Some(Some(r2))) => (r1,r2),
        _ =>  panic!("invalid line {}", line),
    }
}

fn p2(input: &String) -> usize {
    fn not_overlaps(a: &Range, b: &Range) -> bool{
        a.max < b.min || a.min > b.max
    }

    fn not_overlap_at_all(pair: &(Range, Range)) -> bool {
        !not_overlaps(&pair.0, &pair.1) && !not_overlaps(&pair.1, &pair.0)
    }
    
    input.lines()
    .map(parse_line)
    .filter(not_overlap_at_all)
    .count()
}