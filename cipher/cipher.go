package cipher

import "golang.org/x/crypto/bcrypt"

type Cipher interface {
	Encrypt(password string) (string, error)
	Validate(hash, password string) error
}

type BcryptCipher struct{}

func (b BcryptCipher) Encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (b BcryptCipher) Validate(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
