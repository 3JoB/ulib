package main

import (
	"fmt"
	"os"

	"github.com/3JoB/ulib/fsutil/hash"
)

func main() {
	// With Path
	f := hash.NewWithPath("GMakefile.yml")
	// SHA1
	sha1 := f.SHA1()
	fmt.Printf("SHA1: %v\n", sha1)

	// SHA1 HMAC
	f = hash.NewWithPath("GMakefile.yml")
	sha1h := f.HMAC("test").SHA1()
	fmt.Printf("SHA1 HMAC: %v\n", sha1h)

	// SHA256
	f = hash.NewWithPath("GMakefile.yml")
	sha256 := f.SHA256()
	fmt.Printf("SHA256: %v\n", sha256)

	// SHA256 HMAC
	f = hash.NewWithPath("GMakefile.yml")
	sha256h := f.HMAC("test").SHA256()
	fmt.Printf("SHA256 HMAC: %v\n", sha256h)

	// MD5
	f = hash.NewWithPath("GMakefile.yml")
	md5 := f.MD5()
	fmt.Printf("MD5: %v\n", md5)

	// MD5 HMAC
	f = hash.NewWithPath("GMakefile.yml")
	md5h := f.MD5()
	fmt.Printf("MD5 HMAC: %v\n", md5h)

	// CRC32
	f = hash.NewWithPath("GMakefile.yml")
	crc32 := f.CRC32()
	fmt.Printf("CRC32: %v\n", crc32)

	// CRC32 HMAC
	f = hash.NewWithPath("GMakefile.yml")
	crc32h := f.HMAC("test").CRC32()
	fmt.Printf("CRC32 HMAC: %v\n", crc32h)

	// With Os
	y, err := os.Open("GMakefile.yml")
	E(err)

	r_md5 := hash.NewWithOs(y).MD5()
	fmt.Printf("MD5 (With OS): %v\n", r_md5)
}

func E(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-2)
	}
}
