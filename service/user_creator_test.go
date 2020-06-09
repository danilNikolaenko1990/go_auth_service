package service

import (
	"auth_service/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	UserLogin    = "karabas_barabas"
	UserEmail    = "email@lala.com"
	UserPhone    = "12345678"
	UserPass     = "qwerty123"
	UserPassHash = "$2a$10$LI0y1qJfF8mft7ErGhrhMupC6iO7zGGf3ajeXK/UkoTOeYZm6eVZq"
)

func TestUserCreator_Create_GivenUniqueUserAttributes_RegisterUser(t *testing.T) {
	cipherMock := CipherMock{}
	cipherMock.HashToReturn = UserPassHash
	userRepoMock := &UserRepositoryMock{}

	userCreator := UserCreator{
		Cipher:   cipherMock,
		UserRepo: userRepoMock,
	}
	resp := userCreator.Create(givenCreationRequest())

	assert.Equal(t, UserLogin, userRepoMock.InsertedUser.Login)

	assert.Equal(t, UserLogin, userRepoMock.InsertedUser.Login)
	assert.Equal(t, UserEmail, userRepoMock.InsertedUser.Email)
	assert.Equal(t, UserPhone, userRepoMock.InsertedUser.Phone)
	assert.Equal(t, UserPassHash, userRepoMock.InsertedUser.PasswordHash)
	assert.True(t, resp.Registered)
	assert.Empty(t, resp.ErrorMessage)
	assert.Empty(t, resp.ErrorCode)
}

func TestUserCreator_Create_GivenUserAttributeAlreadyExists_ReturnsError(t *testing.T) {
	cipherMock := CipherMock{}
	cipherMock.HashToReturn = UserPassHash
	userRepoMock := &UserRepositoryMock{}
	userRepoMock.UserToReturn = &domain.User{}

	userCreator := UserCreator{
		Cipher:   cipherMock,
		UserRepo: userRepoMock,
	}
	resp := userCreator.Create(givenCreationRequest())

	assert.False(t, resp.Registered)
	assert.Equal(t, "Email, phone or login already exists", resp.ErrorMessage)
	assert.Equal(t, "non_unique_attributes", resp.ErrorCode)
}

func TestUserCreator_Create_GivenUserCipherReturnsError_ReturnsError(t *testing.T) {
	cipherMock := CipherMock{}
	cipherMock.HashToReturn = ""
	cipherMock.ErrToReturn = errors.New("some error")
	userRepoMock := &UserRepositoryMock{}

	userCreator := UserCreator{
		Cipher:   cipherMock,
		UserRepo: userRepoMock,
	}
	resp := userCreator.Create(givenCreationRequest())

	assert.False(t, resp.Registered)
	assert.Equal(t, "some error", resp.ErrorMessage)
	assert.Equal(t, "internal_error", resp.ErrorCode)
}

func TestUserCreator_Create_GivenUserRepositoryReturnsErrorDuringInserting_ReturnsError(t *testing.T) {
	cipherMock := CipherMock{}
	cipherMock.HashToReturn = ""
	userRepoMock := &UserRepositoryMock{}
	userRepoMock.ErrorDuringInsert = errors.New("some error")

	userCreator := UserCreator{
		Cipher:   cipherMock,
		UserRepo: userRepoMock,
	}
	resp := userCreator.Create(givenCreationRequest())

	assert.False(t, resp.Registered)
	assert.Equal(t, "some error", resp.ErrorMessage)
	assert.Equal(t, "internal_error", resp.ErrorCode)
}

func TestUserCreator_Create_GivenUserRepositoryReturnsErrorDuringFindingByAttributes_ReturnsError(t *testing.T) {
	cipherMock := CipherMock{}
	cipherMock.HashToReturn = ""
	userRepoMock := &UserRepositoryMock{}
	userRepoMock.ErrorToReturn = errors.New("some error")

	userCreator := UserCreator{
		Cipher:   cipherMock,
		UserRepo: userRepoMock,
	}
	resp := userCreator.Create(givenCreationRequest())

	assert.False(t, resp.Registered)
	assert.Equal(t, "some error", resp.ErrorMessage)
	assert.Equal(t, "internal_error", resp.ErrorCode)
}

func givenCreationRequest() CreationRequest {
	return CreationRequest{
		Email:    UserEmail,
		Password: UserPass,
		Phone:    UserPhone,
		Login:    UserLogin,
	}
}
