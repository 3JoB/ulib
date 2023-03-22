// This is not a public method package, it is recommended to use an external wrapper.
package flash

import (
	"github.com/dav009/flash"
)

type Flash struct {
	k flash.Keywords
}

func NewPool(keyword ...string) *Flash {
	if len(keyword) == 0 {
		return nil
	}
	f := &Flash{}
	f.k = flash.NewKeywords()
	for _, k := range keyword {
		f.k.Add(k)
	}
	return f
}

func (f *Flash) Add(keyword ...string) {
	if len(keyword) != 0 {
		for _, k := range keyword {
			f.k.Add(k)
		}
	}
}

func (f *Flash) Find(text string) bool {
	return len(f.k.Extract(text)) != 0
}

func Search(text string, keyword ...string) bool {
	if len(keyword) == 0 {
		return false
	}
	key := flash.NewKeywords()
	for _, k := range keyword {
		key.Add(k)
	}
	return len(key.Extract(text)) != 0
}
