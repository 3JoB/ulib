package err

import (
	"errors"

	"github.com/3JoB/ulib/litefmt"
)

type Err struct {
	Op  string
	Err string
	E   error
}

func (e *Err) Error() string {
	if e.E == nil {
		e.E = errors.New(e.Err)
	}
	return litefmt.Sprint(e.Op, "", e.Err)
}

func (e *Err) Unwrap() error {
	if e.E == nil {
		e.E = errors.New(e.Err)
	}
	return e.E
}
