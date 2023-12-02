package validator

import (
	"log"
	"regexp"
)

var (
	emailRegexp *regexp.Regexp
	passRegexp  *regexp.Regexp
)

func init() {
	var err error

	if emailRegexp, err = regexp.Compile("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$"); err != nil {
		log.Fatalln("compile regexp / email :", err)
	}

	if passRegexp, err = regexp.Compile("^[a-zA-Z0-9]+[!@#$%^&*()-_=+,.?/:;{}<>~][a-zA-Z0-9!@#$%^&*()-_=+,.?/:;{}<>~]{5,19}$"); err != nil {
		log.Fatalln("compile regexp / password :", err)
	}

}

func ValidateAll(vlds ...Interface) error {
	for i := range vlds {
		if err := vlds[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}
