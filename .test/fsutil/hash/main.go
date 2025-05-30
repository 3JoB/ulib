package main

import (
	"fmt"
	"os"

	"github.com/3JoB/ulib/fsutil/hash"
)

func main() {
	// SHA1
	sha1 := hash.New("GMakefile.yml", &hash.Opt{Crypt: hash.SHA1})
	fmt.Printf("SHA1: %v\n", sha1)

	// SHA1 HMAC
	sha1h := hash.New("GMakefile.yml", &hash.Opt{Crypt: hash.SHA1, HMACKey: "test"})
	fmt.Printf("SHA1 HMAC: %v\n", sha1h)

	// SHA256
	sha256 := hash.New("GMakefile.yml", &hash.Opt{Crypt: hash.SHA256})
	fmt.Printf("SHA256: %v\n", sha256)

	// SHA256 HMAC
	sha256h := hash.New("GMakefile.yml", &hash.Opt{Crypt: hash.SHA256, HMACKey: "test"})
	fmt.Printf("SHA256 HMAC: %v\n", sha256h)

	// MD5
	md5 := hash.New("GMakefile.yml", &hash.Opt{Crypt: hash.MD5})
	fmt.Printf("MD5: %v\n", md5)

	// MD5 HMAC
	md5h := hash.New("GMakefile.yml", &hash.Opt{Crypt: hash.MD5, HMACKey: "test"})
	fmt.Printf("MD5 HMAC: %v\n", md5h)

	// CRC32
	crc32 := hash.New32("GMakefile.yml", &hash.Opt{Crypt: hash.CRC32})
	fmt.Printf("CRC32: %v\n", crc32)

	// CRC32 HMAC
	crc32h := hash.New32("GMakefile.yml", &hash.Opt{Crypt: hash.CRC32, HMACKey: "test"})
	fmt.Printf("CRC32 HMAC: %v\n", crc32h)
}

func E(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-2)
	}
}
