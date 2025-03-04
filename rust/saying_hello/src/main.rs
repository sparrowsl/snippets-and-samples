use std::io::{self, Write};

fn main() {
    print!("What is your name? ");
    io::stdout().flush().unwrap(); // Flush output to ensure prompt appears before input

    let mut name = String::new();
    _ = io::stdin().read_line(&mut name);

    println!("Hello, {}, nice to meet you!", name.trim());
}
