package service

import (
	"auth_service/domain"
	"errors"
	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	userLogin    = "some_login"
	userPassword = "some_pass"
)

func TestLoginChecker_IsLoggedIn_GivenUserExistsAndHasCorrectPassword_ReturnsSuccessResult(t *testing.T) {
	cipherMock := CipherMock{}
	userRepoMock := &UserRepositoryMock{}
	userRepoMock.UserToReturn = &domain.User{}
	loginChecker := LoginChecker{
		Cipher:   cipherMock,
		UserRepo: userRepoMock,
	}

	result, err := loginChecker.IsLoggedIn(userLogin, userPassword)

	assert.True(t, result)
	assert.Nil(t, err)
}

func TestLoginChecker_IsLoggedIn_GivenUserDoesntExists_ReturnsError(t *testing.T) {
	cipherMock := CipherMock{}
	userRepoMock := &UserRepositoryMock{}
	userRepoMock.ErrorToReturn = pg.ErrNoRows
	loginChecker := LoginChecker{
		Cipher:   cipherMock,
		UserRepo: userRepoMock,
	}

	result, err := loginChecker.IsLoggedIn(userLogin, userPassword)

	assert.False(t, result)
	assert.EqualError(t, err, "user_not_found")
}

func TestLoginChecker_IsLoggedIn_GivenUserRepositoryReturnsSomeUnpredictableErroe_ReturnsError(t *testing.T) {
	cipherMock := CipherMock{}
	userRepoMock := &UserRepositoryMock{}
	userRepoMock.ErrorToReturn = errors.New("some err")
	loginChecker := LoginChecker{
		Cipher:   cipherMock,
		UserRepo: userRepoMock,
	}

	result, err := loginChecker.IsLoggedIn(userLogin, userPassword)

	assert.False(t, result)
	assert.EqualError(t, err, "internal_error")
}

func TestLoginChecker_IsLoggedIn_GivenUserHasWrongPassword_ReturnsError(t *testing.T) {
	cipherMock := CipherMock{}
	cipherMock.ErrToReturn = errors.New("wrong passw")
	userRepoMock := &UserRepositoryMock{}
	userRepoMock.UserToReturn = &domain.User{}
	loginChecker := LoginChecker{
		Cipher:   cipherMock,
		UserRepo: userRepoMock,
	}

	result, err := loginChecker.IsLoggedIn(userLogin, userPassword)

	assert.False(t, result)
	assert.EqualError(t, err, "invalid_password")
}
