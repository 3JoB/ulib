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
	e.E = errors.New(e.Err)
	return fmt.Sprintf("%v: %v", e.Op, e.Err)
}

func (e *Err) Unwrap() error {
	e.E = errors.New(e.Err)
	return e.E
}
