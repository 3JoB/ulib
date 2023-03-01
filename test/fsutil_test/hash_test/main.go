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
	sha1, err := f.SHA1()
	E(err)
	fmt.Printf("SHA1: %v\n", sha1)

	// SHA256
	sha256, err := f.SHA256()
	E(err)
	fmt.Printf("SHA256: %v\n", sha256)

	// MD5
	md5, err := f.MD5()
	E(err)
	fmt.Printf("MD5: %v\n", md5)

	// crc32
	crc32, err := f.CRC32()
	E(err)
	fmt.Printf("CRC32: %v\n", crc32)

	// With Os
	y, err := os.Open("GMakefile.yml")
	E(err)

	r_md5, err := hash.NewWithOs(y).MD5()
	E(err)
	fmt.Printf("MD5: %v\n", r_md5)
}

func E(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-2)
	}
}
