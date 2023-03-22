package keyword

import (
	"github.com/3JoB/ulib/keyword/bms"
	"github.com/3JoB/ulib/keyword/flash"
	"github.com/3JoB/ulib/keyword/regexp"
)

// Use the regexp method to find whether a keyword exists.
func RegexpFind(text, keyword string) bool {
	return regexp.Match(text, keyword)
}

// Use the compilation method of regexp to find whether the keyword exists.
func RegexpCompile(keyword string) (*regexp.KeywordCompile, error) {
	k := &regexp.KeywordCompile{}
	if err := k.Init(keyword); err != nil {
		return nil, err
	}
	return k, nil
}

// Use Boyer-Moore's compilation method to find the existence of keywords
func BoyerMooreCompile(keyword string) (*bms.BMSCompile, error) {
	k := &bms.BMSCompile{}
	if err := k.Init(keyword); err != nil {
		return nil, err
	}
	return k, nil
}

// Use the Boyer-Moore method to find the existence of keywords.
func BoyerMooreFind(text, keyword string) bool {
	return bms.Search(text, keyword)
}

// Use Flash's compilation method to find out whether keywords exist.
func FlashNewPool(keyword ...string) *flash.Flash {
	if len(keyword) == 0 {
		return nil
	}
	return flash.NewPool(keyword...)
}

// Use the Flash method to find whether the keyword exists.
func FlashSearch(text string, keyword ...string) bool {
	return flash.Search(text, keyword...)
}
