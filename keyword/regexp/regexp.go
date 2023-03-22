// This is not a public method package, it is recommended to use an external wrapper.
package regexp

import (
	"errors"
	"fmt"

	"github.com/grafana/regexp"
)

var ErrNotCompiled error = errors.New("not compiled")

type KeywordCompile struct {
	Compile *regexp.Regexp
}

// DO NOT USE!
//
// This is not a public function!
func (c *KeywordCompile) Init(keyword string) (err error) {
	keyword = fmt.Sprintf("(^|\\s)%v(\\s|$)", keyword)
	c.Compile, err = regexp.Compile(keyword)
	return
}

func (c *KeywordCompile) Find(text string) (bool, error) {
	if c.Compile == nil {
		return false, ErrNotCompiled
	}
	return c.Compile.MatchString(text), nil
}

func Match(text, keyword string) bool {
	if match, err := regexp.MatchString(keyword, text); err == nil {
		return match
	} else {
		return false
	}
}
