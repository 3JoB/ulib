// This is not a public method package, it is recommended to use an external wrapper.
package bms

import (
	"github.com/cubicdaiya/bms"
)

type BMSCompile struct {
	Keyword string
	Table   map[rune]int
}

// DO NOT USE!
//
// This is not a public function!
func (c *BMSCompile) Init(keyword string) {
	c.Keyword = keyword
	c.Table = make(map[rune]int)
	c.Table = bms.BuildSkipTable(keyword)
}

func (c *BMSCompile) Find(text string) bool {
	if c.Table == nil {
		return false
	}
	return bms.SearchBySkipTable(text, c.Keyword, c.Table) != 0
}

func Search(text, keyword string) bool {
	return bms.Search(text, keyword) != 0
}
