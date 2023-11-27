package bcrypt

import "testing"

func TestBcrypt_Generate(t *testing.T) {
	secret := "31veli31"
	bc := NewBcrypt(secret, 10)

	if _, err := bc.Generate("testpass"); err != nil {
		t.Error(err)
	}
}

func TestBcrypt_Compare(t *testing.T) {
	secret := "31veli31"
	bc := NewBcrypt(secret, 10)

	var (
		hashed string
		pass   = "abcabcoglusad"
		err    error
	)

	if hashed, err = bc.Generate(pass); err != nil {
		t.Error(err)
	}

	t.Run("Compare", func(t *testing.T) {
		if err = bc.Compare(hashed, pass); err != nil {
			t.Error(err)
		}
	})

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("WrongPass", func(t *testing.T) {
			if err = bc.Compare(hashed, "veli123"); err == nil {
				t.Errorf("expected: hashedPassword is not the hash of the given password")
			}
		})
	})
}
