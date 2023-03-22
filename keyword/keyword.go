package keyword

import (
	"github.com/3JoB/ulib/keyword/bms"
	"github.com/3JoB/ulib/keyword/flash"
	"github.com/3JoB/ulib/keyword/regexp"
)

func RegexpFind(text, keyword string) bool {
	return regexp.Match(text, keyword)
}

func RegexpCompile(keyword string) (*regexp.KeywordCompile, error) {
	k := &regexp.KeywordCompile{}
	if err := k.Init(keyword); err != nil {
		return nil, err
	}
	return k, nil
}

// Boyer-Moore
func BoyerMooreCompile(keyword string) (*bms.BMSCompile, error) {
	k := &bms.BMSCompile{}
	if err := k.Init(keyword); err != nil {
		return nil, err
	}
	return k, nil
}

func BoyerMooreFind(text, keyword string) bool {
	return bms.Search(text, keyword)
}

func FlashNewPool(keyword ...string) *flash.Flash {
	if len(keyword) == 0 {
		return nil
	}
	return flash.NewPool(keyword...)
}

func FlashSearch(text string, keyword ...string) bool {
	return flash.Search(text, keyword...)
}
