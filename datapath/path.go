package datapath

import (
	"errors"
)

type Path []Part

func (p Path) String() string {
	return "" // TODO
}

func (p Path) Validate() error {
	errs := make([]error, 0, len(p))
	for _, part := range p {
		if err := part.Validate(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}