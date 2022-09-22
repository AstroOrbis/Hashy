use clap::{App, Arg, ArgMatches};

fn main() {
    let matches: ArgMatches = App::new("Hashy")
        .version("0.1.0")
        .author("Astro Orbis <astroorbis@gmail.com>")
        .about("A simple hashing utility!")
        .arg(
            Arg::with_name("string")
                .short('s')
                .long("string")
                .takes_value(true)
                .help("String to hash"),
        )
        .get_matches();



}
