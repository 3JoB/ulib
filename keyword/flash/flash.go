package flash

import (
	"github.com/3JoB/ulib/internal/keyword/flash"
)

// Use Flash's compilation method to find out whether keywords exist.
func NewPool(keyword ...string) *flash.Flash {
	if len(keyword) == 0 {
		return nil
	}
	return flash.NewPool(keyword...)
}

// Use the Flash method to find whether the keyword exists.
func Search(text string, keyword ...string) bool {
	return flash.Search(text, keyword...)
}
