package service

import (
	"auth-service/cipher"
	"auth-service/domain"
	"errors"
	"github.com/go-pg/pg"
	"time"
)

const (
	InternalErrorErrorCode        = "internal_error"
	AlreadyHasAttributesErrorCode = "non_unique_attributes"
	NonUniqueErrorMessage         = "Email, phone or login already exists"
)

type UserCreator struct {
	UserRepo domain.UserRepository
	Cipher   cipher.Cipher
}

type CreationResponse struct {
	Registered   bool
	ErrorMessage string
	ErrorCode    string
}

type CreationRequest struct {
	Login    string
	Email    string
	Phone    string
	Password string
}

func (uc *UserCreator) Create(req CreationRequest) CreationResponse {
	user, err := uc.UserRepo.FindByUniqueAttributes(req.Login, req.Email, req.Phone)
	if err != nil && !errors.Is(err, pg.ErrNoRows) {
		return CreationResponse{ErrorCode: InternalErrorErrorCode, ErrorMessage: err.Error()}
	}
	if user != nil {
		return CreationResponse{
			ErrorCode:    AlreadyHasAttributesErrorCode,
			ErrorMessage: NonUniqueErrorMessage,
		}
	}
	hash, err := uc.Cipher.Encrypt(req.Password)
	if err != nil {
		return CreationResponse{ErrorCode: InternalErrorErrorCode, ErrorMessage: err.Error()}
	}

	err = uc.UserRepo.Insert(&domain.User{
		Login:        req.Login,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: hash,
		CreatedAt:    time.Now().Format(time.RFC3339),
	})

	if err != nil {
		return CreationResponse{ErrorCode: InternalErrorErrorCode, ErrorMessage: err.Error()}
	}

	return CreationResponse{Registered: true}
}
