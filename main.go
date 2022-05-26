package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// References for anyone wanting to help!
// https://gobyexample.com/command-line-flags
// https://shapeshed.com/unix-exit-codes/

func main() {
	hashTypeFlag := flag.String("hash", "", "Hash flag")
	verboseFlag := flag.Bool("v", false, "Returns verbose output.")
	hashFlag := flag.String("string", "", "Hashes the given string.")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Please input a valid string. \nExample: ./Hashy -hash md5 -string e")
		os.Exit(1)
	}

	if len(*hashFlag) < 1 {
		fmt.Println("No string specified to hash. \nUse the -string parameter to specify a string to hash")
		os.Exit(1)
	}

	hashType := *hashTypeFlag
	compiledHash := ""

	// Supported hash types: md5, sha1, sha256, sha512
	switch strings.ToLower(hashType) {
	case "md5":
		compiledHash = MD5hash(*hashFlag)
	case "sha1":
		compiledHash = SHA1hash(*hashFlag)
	case "sha256":
		compiledHash = SHA256hash(*hashFlag)
	case "sha512":
		compiledHash = SHA512hash(*hashFlag)
	default:
		fmt.Println("Unspecified hash type: " + hashType + "\nPlease use either MD5, SHA1, SHA256, or SHA512")
	}

	if *verboseFlag {
		fmt.Println(strings.ToUpper(hashType) + " hash for " + *hashFlag + ": " + compiledHash)
	} else {
		fmt.Println(compiledHash)
	}
}
