package main

import (
	"github.com/3JoB/ulib/litefmt"
)

var dc []string = []string{
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"ccccc",
}

func main() {
	e := litefmt.PSprint(dc...)
	litefmt.Println(e)
	e1 := litefmt.PSprint(dc...)
	litefmt.Println(e1)
	e2 := litefmt.PSprint(dc...)
	litefmt.Println(e2)
	e3 := litefmt.PSprint(dc...)
	litefmt.Println(e3)
	e4 := litefmt.PSprint(dc...)
	litefmt.Println(e4)
}
