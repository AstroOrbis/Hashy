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
	hashTypeFlag := flag.String("hash", "", "Hash algorithm to use. Available algorithms: md5, sha1, sha256, sha512")
	verboseFlag := flag.Bool("v", false, "Returns verbose output.")
	hashFlag := flag.String("string", "", "String to hash")

	// change usage message to display the mandatory parameters
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -hash <algorithm> -string <input> [-v]\n", os.Args[0])
		flag.PrintDefaults()
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "\nExample: %s -hash md5 -string e\n", os.Args[0])

	}

	flag.Parse()

	if len(*hashTypeFlag) == 0 || len(*hashFlag) == 0 {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "The flags -hash and -string are required\n\n")
		flag.Usage()
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
