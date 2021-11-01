package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// https://gobyexample.com/command-line-flags

func main() {
	md5flag := flag.Bool("md5", false, "Returns a md5 hash.")
	sha1flag := flag.Bool("sha1", false, "Returns a sha1 hash.")

	verboseflag := flag.Bool("v", false, "Returns verbose output.")

	flag.Parse()

	StrArg := strings.Join(flag.Args(), " ")

	// Chcek for multiple args
	if strings.Contains(StrArg, " ") {
		fmt.Println("Please only specify 1 string to hash.")
	}

	// Check for multiple hashtypes
	if *md5flag && *sha1flag {
		fmt.Println("Please use only 1 hashtype flag.")
		os.Exit(1) //https://shapeshed.com/unix-exit-codes/
	}

	// Check for lack of flags
	if !*md5flag && !*sha1flag {
		fmt.Println("Please provide one hashtype flag and one (or multiple) argument(s). All flags can be listed with the -h flag.")
	}

	// Checks for verbosity
	if *verboseflag {

		if *md5flag {
			fmt.Println(verboseMD5encrypt(StrArg))
		}

		if *sha1flag {
			fmt.Println(verboseSHA1encrypt(StrArg))
		}

	} else {

		if *md5flag {
			fmt.Println(MD5encrypt(StrArg))
		}

		if *sha1flag {
			fmt.Println(SHA1encrypt(StrArg))
		}

	}

}
