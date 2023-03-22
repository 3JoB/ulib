package main

import (
	"fmt"

	"github.com/3JoB/ulib/keyword"
)

func main() {
	text := "hello world"
	kw := "world"
	flash_pool := keyword.FlashNewPool(kw)
	fmt.Printf("Flash: %v\n", flash_pool.Find(text))
	fmt.Printf("BoyerMooreFind: %v\n", keyword.BoyerMooreFind(text, kw))
	bm, _ := keyword.BoyerMooreCompile(kw)
	fmt.Printf("BoyerMooreCompile: %v\n", bm.Find(text))
	re, _ := keyword.RegexpCompile(kw)
	reb, _ := re.Find(text)
	fmt.Printf("RegexpFind: %v\n", reb)
	fmt.Printf("RegexpFind: %v\n", keyword.RegexpFind(text, kw))
}