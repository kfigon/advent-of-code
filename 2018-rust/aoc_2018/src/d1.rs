use std::{collections::HashSet, num::ParseIntError};


fn p1(input: &str) -> Result<i32, String>{
    input
    .lines()
    .map(|v| v.parse::<i32>())
    .try_fold(0, |acc, v| v.map(|i| i+acc))
    .map_err(|err| err.to_string())
}

#[derive(PartialEq,Debug)]
enum Error {
    NotFound,
    ParseErr(ParseIntError),
}

impl From<std::num::ParseIntError> for Error {
    fn from(e: std::num::ParseIntError) -> Self {
        Error::ParseErr(e)
    }
}

fn p2(input: &str) -> Result<i32, Error> {
    let mut freq: HashSet<i32> = HashSet::new();
    let mut current_v = 0;

    freq.insert(current_v);
    let list = input.lines()
        .map(|v| v.parse::<i32>())
        .collect::<Result<Vec<i32>, ParseIntError> >()?;

    let loop_guard = 10_000_000;
    let mut i = 0;
    for v in list.into_iter().cycle() {
        if i > loop_guard {
            break;
        }

        current_v += v;

        if !freq.insert(current_v) {
            return Ok(current_v);
        }

        i+=1;
    }

    return Err(Error::NotFound);
}

#[cfg(test)]
mod test {
    use std::fs;
    use super::*;
    
    #[test]
    fn test_p1() {
        let exp = 454;
        let data = fs::read_to_string("d1.txt").unwrap();
        assert_eq!(Ok(exp), p1(&data));
    }

    #[test]
    fn test_p2() {
        let exp = 566;
        let data = fs::read_to_string("d1.txt").unwrap();
        assert_eq!(Ok(exp), p2(&data));
    }

    #[test]
    fn test_p2_e1() {
        let data = "+1, -1".replace(", ", "\n");
        assert_eq!(Ok(0), p2(&data));
    }

    #[test]
    fn test_p2_e2() {
        let data = "3, +3, +4, -2, -4".replace(", ", "\n");
        assert_eq!(Ok(10), p2(&data));
    }

    #[test]
    fn test_p2_e3() {
        let data = "-6, +3, +8, +5, -6".replace(", ", "\n");
        assert_eq!(Ok(5), p2(&data));
    }

    #[test]
    fn test_p2_e4() {
        let data = "+7, +7, -2, -7, -4".replace(", ", "\n");
        assert_eq!(Ok(14), p2(&data));
    }

    #[test]
    fn test_p2_e5() {
        let data = "1, -2, +3, +1".replace(", ", "\n");
        assert_eq!(Ok(2), p2(&data));
    }
}