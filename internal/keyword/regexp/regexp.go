// This is not a public method package, it is recommended to use an external wrapper.
package regexp

import (
	"errors"

	"github.com/grafana/regexp"

	"github.com/3JoB/ulib/litefmt"
)

var ErrNotCompiled error = errors.New("not compiled")

type KeywordCompile struct {
	Compile *regexp.Regexp
}

// DO NOT USE!
//
// This is not a public function!
func (c *KeywordCompile) Init(keyword string) (err error) {
	keyword = litefmt.Sprint("(^|\\s)", keyword, "(\\s|$)")
	c.Compile, err = regexp.Compile(keyword)
	return
}

func (c *KeywordCompile) Find(text string) bool {
	if c.Compile == nil && text == "" {
		return false
	}
	return c.Compile.MatchString(text)
}

func Match(text, keyword string) bool {
	if match, err := regexp.MatchString(keyword, text); err == nil {
		return match
	} else {
		return false
	}
}
