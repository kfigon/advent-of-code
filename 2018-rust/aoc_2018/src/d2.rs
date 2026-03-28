use std::{collections::{HashMap, HashSet}};

fn p1(input: &str) -> i32 {
    let count_freqs = |l: &str| -> HashMap<char, i32> {
        l.chars().fold(HashMap::<char, i32>::new(), |mut acc, c|{
            acc.entry(c).and_modify(|counter| *counter += 1).or_insert(1);
            acc
         })
    };

    let x = input.lines()
        .map(count_freqs)
        .fold((0,0), |acc, v|{
            let has2 = v.iter().any(|kv| *kv.1 == 2);
            let has3 = v.iter().any(|kv| *kv.1 == 3);
            match (has2, has3) {
                (true, true) => (acc.0 + 1, acc.1 + 1),
                (true, false) => (acc.0 + 1, acc.1),
                (false, true) => (acc.0, acc.1 + 1),
                (false, false) => acc,
            }
        });

    x.0 * x.1
}

fn p2(input: &str) -> Option<String> {
    let lines: Vec<&str> = input.lines().collect();
    if lines.is_empty() {
        return None;
    }
    let first_len = lines.get(0).unwrap().len(); // byte size, good enought

    // check horizontal
    for c_idx in 0..first_len {
        let mut set: HashSet<String> = HashSet::new();
        for line in &lines {
        
           let modified_line = format!("{}{}", &line[..c_idx],  &line[c_idx+1..]);
           if !set.insert(modified_line.clone()) {
            return Some(modified_line);
           }
        }    
    }
    None
}

#[cfg(test)]
mod test {
    use std::fs;

    use super::*;

    #[test]
    fn p1_ex() {
        let example = "abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab";
        assert_eq!(12, p1(example));
    }

    #[test]
    fn p1_test() {
        let exp = 6150;
        assert_eq!(exp, p1(&fs::read_to_string("d2.txt").unwrap()));
    }
    
    #[test]
    fn p2_ex() {
        let data="abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz";
        assert_eq!(Some("fgij".to_string()), p2(data));
    }

    #[test]
    fn p2_test() {
        assert_eq!(Some("rteotyxzbodglnpkudawhijsc".to_string()), p2(&fs::read_to_string("d2.txt").unwrap()));
    }
}