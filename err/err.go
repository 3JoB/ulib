package err

import (
	"errors"
	"fmt"
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
	return fmt.Sprintf("%v: %v", e.Op, e.Err)
}

func (e *Err) Unwrap() error {
	if e.E == nil {
		e.E = errors.New(e.Err)
	}
	return e.E
}
