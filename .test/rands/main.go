package main

import (
	"fmt"

	"github.com/3JoB/ulib/rands"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	cr := rands.CRands(data, 3)
	fmt.Println(cr)
	r := rands.Rands(data, 3)
	fmt.Println(r)
	sdata := []string{"a", "b", "c", "d", "e", "f", "g"}
	cs := rands.CRandString(sdata, 3)
	fmt.Println(cs)
	s := rands.RandString(sdata, 3)
	fmt.Println(s)
}
