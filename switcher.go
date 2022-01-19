package main

func switcher(md5flag bool, sha1flag bool, sha256flag bool, sha512flag bool, verboseflag bool, stringflag string) string {
	if verboseflag {
		return verboseSwitcher(md5flag, sha1flag, sha256flag, sha512flag, stringflag)
	} else {
		return nonVerboseSwitcher(md5flag, sha1flag, sha256flag, sha512flag, stringflag)
	}

}

func verboseSwitcher(md5flag bool, sha1flag bool, sha256flag bool, sha512flag bool, stringflag string) string {
	text := ""
	if md5flag {
		text = verboseMD5hash(stringflag)
	}

	if sha1flag {
		text = verboseSHA1hash(stringflag)
	}

	if sha256flag {
		text = verboseSHA256hash(stringflag)
	}

	if sha512flag {
		text = verboseSHA512hash(stringflag)
	}
	return text
}

func nonVerboseSwitcher(md5flag bool, sha1flag bool, sha256flag bool, sha512flag bool, stringflag string) string {
	text := ""
	if md5flag {
		text = MD5hash(stringflag)
	}

	if sha1flag {
		text = SHA1hash(stringflag)
	}

	if sha256flag {
		text = SHA256hash(stringflag)
	}

	if sha512flag {
		text = SHA512hash(stringflag)
	}
	return text
}
