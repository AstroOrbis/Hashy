mod lib;
use lib::*;

fn easyselect(prompt: &str, choices: Vec<String>) -> String {
    let choice = inquire::Select::new(prompt, choices).prompt().unwrap();

    choice
}

fn main() {
    println!("Hello, world! in md5: {}", md5("Hello, World!".to_string()));
}
