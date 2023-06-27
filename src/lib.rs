use md5 as md5lib;
use sha1::{Digest, Sha1};
use sha2::{Sha256, Sha512};

pub fn md5(s: String) -> String {
    let digest = md5lib::compute(s);
    format!("{:x}", digest)
}

pub fn sha1(s: String) -> String {
    let mut hasher = Sha1::new();
    hasher.update(s);
    let result = hasher.finalize();
    format!("{:x}", result)
}

pub fn sha256(s: String) -> String {
    let mut hasher = Sha256::new();
    hasher.update(s);
    let result = hasher.finalize();
    format!("{:x}", result)
}

pub fn sha512(s: String) -> String {
    let mut hasher = Sha512::new();
    hasher.update(s);
    let result = hasher.finalize();
    format!("{:x}", result)
}

/////////////////////
////////Tests////////
/////////////////////

#[cfg(test)]
mod tests {
    use super::*;

    // MD5

    #[test]
    fn md5_1() {
        let input = String::from("Hello, World!");
        let expected_output = String::from("65a8e27d8879283831b664bd8b7f0ad4");
        assert_eq!(md5(input), expected_output);
    }

    #[test]
    fn md5_2() {
        let input = String::from("Rust is awesome!");
        let expected_output = String::from("514bcadd095418e9be47628a0da4d6a8");
        assert_eq!(md5(input), expected_output);
    }

    // Sha1

    #[test]
    fn sha1_1() {
        let input = String::from("Hello, World!");
        let expected_output = String::from("0a0a9f2a6772942557ab5355d76af442f8f65e01");
        assert_eq!(sha1(input), expected_output);
    }

    #[test]
    fn sha1_2() {
        let input = String::from("Rust is awesome!");
        let expected_output = String::from("031c18148eb1e55d82d8b6669bb5be0db5a1bd3a");
        assert_eq!(sha1(input), expected_output);
    }

    // Sha256

    #[test]
    fn sha256_1() {
        let input = String::from("Hello, World!");
        let expected_output =
            String::from("dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f");
        assert_eq!(sha256(input), expected_output);
    }

    #[test]
    fn sha256_2() {
        let input = String::from("Rust is awesome!");
        let expected_output =
            String::from("e13b847b65df6185756ae240f3d2aeff3b3062d0423394f5b5139725695dd30f");
        assert_eq!(sha256(input), expected_output);
    }

    // Sha512

    #[test]
    fn sha512_1() {
        let input = String::from("Hello, World!");
        let expected_output =
            String::from("374d794a95cdcfd8b35993185fef9ba368f160d8daf432d08ba9f1ed1e5abe6cc69291e0fa2fe0006a52570ef18c19def4e617c33ce52ef0a6e5fbe318cb0387");
        assert_eq!(sha512(input), expected_output);
    }

    #[test]
    fn sha512_2() {
        let input = String::from("Rust is awesome!");
        let expected_output =
            String::from("2c88c3efd1b0528908c11ce2d1162b9d005da4ee5265bc59ca10af289800dab1cc5a7ff659085d425485fcc741fe4ea325e9ef14493210472faf50b5152fbab2");
        assert_eq!(sha512(input), expected_output);
    }
}
