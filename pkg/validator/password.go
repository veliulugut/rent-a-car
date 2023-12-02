package validator

import "fmt"

var _ Interface = (*IsPasswordValid)(nil)

type IsPasswordValid struct {
	Password string
}

func (i *IsPasswordValid) Validate() error {
	if !passRegexp.MatchString(i.Password) {
		return fmt.Errorf("is password valid :%w", ErrrInvalidPassword)
	}

	return nil
}
