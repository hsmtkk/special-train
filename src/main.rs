mod armstrong;

fn main() {
    for i in 1..1000 {
        if armstrong::is_armstrong(i) {
            println!("{}", i);
        }
    }
}
