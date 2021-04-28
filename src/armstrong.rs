use std::convert::TryInto;

pub fn is_armstrong(n: i64) -> bool {
    let mut s = 0;
    let p: u32 = n.to_string().len().try_into().unwrap();
    for n in number_to_digit(n) {
        s += n.pow(p);
    }
    return n == s;
}

#[test]
fn test_is_armstrong(){
    assert!(is_armstrong(9));
    assert!(!is_armstrong(10));
    assert!(is_armstrong(153));
    assert!(!is_armstrong(154));
}

fn number_to_digit(n: i64) -> Vec<i64> {
    let mut result: Vec<i64> = Vec::new();
    for d in n.to_string().chars() {
        let m = d.to_digit(10).unwrap();
        result.push(m.into());
    }
    return result;
}

#[test]
fn test_number_to_digit(){
    let want = vec![1,2,3,4,5];
    let got = number_to_digit(12345);
    assert_eq!(want,got);
}