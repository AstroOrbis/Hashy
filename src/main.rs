use hashy::*;

fn easyselect(prompt: &str, choices: Vec<String>) -> String {
	inquire::Select::new(prompt, choices).prompt().unwrap()
}

fn easyinput(prompt: &str) -> String {
	inquire::Text::new(prompt).prompt().unwrap()
}

fn main() {
	let input: String = easyinput("Enter a string to hash: ");
	let hash_types: Vec<String> = vec![
		String::from("MD5"),
		String::from("SHA1"),
		String::from("SHA256"),
		String::from("SHA512"),
	];
	let hash_type: String = easyselect("Select a hash type:", hash_types);
	let hash: String = match hash_type.as_str() {
		"MD5" => md5(input.clone()),
		"SHA1" => sha1(input.clone()),
		"SHA256" => sha256(input.clone()),
		"SHA512" => sha512(input.clone()),
		_ => String::from("Error: Invalid hash type"),
	};
	println!("{} hash of {}: {}", hash_type, input, hash);
}
