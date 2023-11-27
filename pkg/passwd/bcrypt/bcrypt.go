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
		return "", fmt.Errorf("bcrypt / generate :%w", err)
	}

	return string(generated), nil
}

func (b *Bcrypt) Compare(hashed, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(b.secret+password+b.secret)); err != nil {
		return fmt.Errorf("bcrypt / compare :%w", err)
	}

	return nil
}
