package service

type CipherMock struct {
	PasswordToEncrypt  string
	HashToReturn       string
	ErrToReturn        error
	HashToValidate     string
	PasswordToValidate string
}

func (c CipherMock) Encrypt(password string) (string, error) {
	c.PasswordToEncrypt = password
	return c.HashToReturn, c.ErrToReturn
}

func (c CipherMock) Validate(hash, password string) error {
	c.HashToValidate = hash
	c.PasswordToValidate = password
	return c.ErrToReturn
}
