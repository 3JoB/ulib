package bms

import (
	"github.com/3JoB/ulib/internal/keyword/bms"
)

// Use Boyer-Moore's compilation method to find the existence of keywords
func Compile(keyword string) (*bms.BMSCompile, error) {
	k := &bms.BMSCompile{}
	if err := k.Init(keyword); err != nil {
		return nil, err
	}
	return k, nil
}

// Use the Boyer-Moore method to find the existence of keywords.
func Find(text, keyword string) bool {
	return bms.Search(text, keyword)
}
