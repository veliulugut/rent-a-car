package bcrypt

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"rent-a-car/pkg/passwd"
)

var _ passwd.Interface = (*Bcrypt)(nil)

type Bcrypt struct {
	secret string
	cost   int
}

const (
	errGenerate = "bcrypt / generate: %w"
	errCompare  = "bcrypt / compare: %w"
)

func NewBcrypt(s string, c int) *Bcrypt {
	return &Bcrypt{
		secret: s,
		cost:   c,
	}
}

func (b *Bcrypt) Generate(password string) (string, error) {
	var (
		generated []byte
		err       error
	)

	if generated, err = bcrypt.GenerateFromPassword([]byte(b.secret+password+b.secret), b.cost); err != nil {
		return "", fmt.Errorf(errGenerate, err)
	}

	return string(generated), nil
}

func (b *Bcrypt) Compare(hashed, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(b.secret+password+b.secret)); err != nil {
		return fmt.Errorf(errCompare, err)
	}

	return nil
}
