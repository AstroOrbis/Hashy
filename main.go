package main

import (
	"flag"
	"fmt"
	"os"
)

// References for anyone wanting to help!
// https://gobyexample.com/command-line-flags
//https://shapeshed.com/unix-exit-codes/

func main() {
	md5flag := flag.Bool("md5", false, "Returns a md5 hash.")
	sha1flag := flag.Bool("sha1", false, "Returns a sha1 hash.")
	sha256flag := flag.Bool("sha256", false, "Returns a sha256 hash.")
	sha512flag := flag.Bool("sha512", false, "Returns a sha512 hash.")
	verboseflag := flag.Bool("v", false, "Returns verbose output.")

	stringflag := flag.String("string", "", "Hashes the given string.")
	// StrArg := strings.Join(flag.Args(), " ")

	flag.Parse()

	// Huge thanks to "bereal" on StackOverflow!
	// https://stackoverflow.com/a/69963331/17029825
	flags := []*bool{md5flag, sha1flag, sha256flag, sha512flag}
	seenSetFlag := false
	for _, f := range flags {
		if *f {
			if seenSetFlag {
				fmt.Println("Please only use one hash-type flag.")
				os.Exit(1)
			}
			seenSetFlag = true
		}
	}

	if *stringflag == "" {
		fmt.Println("Please input a valid string. Example: ./Hashy -md5 -string e")
		os.Exit(1)
	}

	fmt.Println(switcher(*md5flag, *sha1flag, *sha256flag, *sha512flag, *verboseflag, *stringflag))
}
