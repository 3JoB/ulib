// This is not a public method package, it is recommended to use an external wrapper.
package bms

import (
	"github.com/cubicdaiya/bms"
)

type BMSCompile struct {
	Keyword string
	Table   map[rune]int
}

// Init initializes the BMSCompile instance with the given keyword and constructs a skip table for Boyer-Moore search.
func (c *BMSCompile) Init(keyword string) {
	c.Keyword = keyword
	c.Table = make(map[rune]int)
	c.Table = bms.BuildSkipTable(keyword)
}

// Find checks if the specified text contains the compiled keyword using the skip table for efficient searching.
func (c *BMSCompile) Find(text string) bool {
	if c.Table == nil {
		return false
	}
	return bms.SearchBySkipTable(text, c.Keyword, c.Table) != 0
}

// Search checks if the keyword exists in the provided text using the Boyer-Moore string search algorithm.
func Search(text, keyword string) bool {
	return bms.Search(text, keyword) != 0
}
