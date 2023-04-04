package regexp

import (
	"github.com/3JoB/ulib/internal/keyword/regexp"
)

// Use the regexp method to find whether a keyword exists.
func Find(text, keyword string) bool {
	return regexp.Match(text, keyword)
}

// Use the compilation method of regexp to find whether the keyword exists.
func Compile(keyword string) (*regexp.KeywordCompile, error) {
	k := &regexp.KeywordCompile{}
	if err := k.Init(keyword); err != nil {
		return nil, err
	}
	return k, nil
}
