package service

import (
	"auth_service/cipher"
	"auth_service/domain"
	"errors"
	"github.com/go-pg/pg"
)

const InvalidPasswordErrorCode = "invalid_password"
const InternalErrorCode = "internal_error"
const UserNotFound = "user_not_found"

type LoginChecker struct {
	UserRepo domain.UserRepository
	Cipher   cipher.Cipher
}

func (c *LoginChecker) IsLoggedIn(login, password string) (bool, error) {
	user, err := c.UserRepo.FindByLogin(login)

	if errors.Is(err, pg.ErrNoRows) {
		return false, errors.New(UserNotFound)
	}
	if err != nil {
		return false, errors.New(InternalErrorCode)
	}

	err = c.Cipher.Validate(user.PasswordHash, password)
	if err != nil {
		return false, errors.New(InvalidPasswordErrorCode)
	}

	return true, nil
}
