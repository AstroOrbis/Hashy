package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func MD5encrypt(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

func verboseMD5encrypt(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return "MD5 hash for " + text + ": " + hash
}

func SHA1encrypt(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

func verboseSHA1encrypt(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return "SHA1 hash for " + text + ": " + hash
}
