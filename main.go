package main

import (
	"flag"
	"fmt"
	"os"
)

// https://gobyexample.com/command-line-flags

func main() {
	md5flag := flag.Bool("md5", false, "Returns a md5 hash.")
	sha1flag := flag.Bool("sha1", false, "Returns a sha1 hash.")
	sha256flag := flag.Bool("sha256", false, "Returns a sha256 hash.")
	sha512flag := flag.Bool("sha512", false, "Returns a sha512 hash.")
	verboseflag := flag.Bool("v", false, "Returns verbose output.")

	stringflag := flag.String("string", "", "Hashes the given string.")
	// StrArg := strings.Join(flag.Args(), " ")

	flag.Parse()

	// Check for multiple hashtypes and lack of flags
	hashtypeflaglist := 0
	if *md5flag {
		hashtypeflaglist++
	}
	if *sha1flag {
		hashtypeflaglist++
	}
	if *sha256flag {
		hashtypeflaglist++
	}
	if *sha512flag {
		hashtypeflaglist++
	}

	if hashtypeflaglist > 1 {
		fmt.Println("Please use only 1 hashtype flag.")
		os.Exit(1) //https://shapeshed.com/unix-exit-codes/
	}

	if hashtypeflaglist == 0 {
		fmt.Println("Please provide one hashtype flag and one (or multiple) argument(s). All flags can be listed with the -h flag.")
	}

	// TODO: Update the switcher
	if *verboseflag {

		if *md5flag {
			fmt.Println(verboseMD5encrypt(*stringflag))
		}

		if *sha1flag {
			fmt.Println(verboseSHA1encrypt(*stringflag))
		}

		if *sha256flag {
			fmt.Println(verboseSHA256encrypt(*stringflag))
		}

		if *sha512flag {
			fmt.Println(verboseSHA512encrypt(*stringflag))
		}

	} else {

		if *md5flag {
			fmt.Println(MD5encrypt(*stringflag))
		}

		if *sha1flag {
			fmt.Println(SHA1encrypt(*stringflag))
		}

		if *sha256flag {
			fmt.Println(SHA256encrypt(*stringflag))
		}

		if *sha512flag {
			fmt.Println(SHA512encrypt(*stringflag))
		}

	}

}
