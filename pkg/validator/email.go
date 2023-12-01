package validator

import "fmt"

var _ Interface = (*IsMailValid)(nil)

type IsMailValid struct {
	Email string
}

func (i *IsMailValid) Validate() error {
	if !emailRegexp.MatchString(i.Email) {
		return fmt.Errorf("is email valid :%w", ErrInvalidEmail)
	}

	return nil
}
