use hashy::*;

fn easyselect(prompt: &str, choices: Vec<String>) -> String {
    inquire::Select::new(prompt, choices).prompt().unwrap()
}

fn main() {
    println!("Hello, world! in md5: {}", md5("Hello, World!".to_string()));
}
