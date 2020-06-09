package service

import (
	"auth_service/domain"
)

type UserRepositoryMock struct {
	ErrorToReturn     error
	ErrorDuringInsert error
	RequestedLogin    string
	RequestedEmail    string
	RequestedPhone    string
	InsertedUser      *domain.User
	UserToReturn      *domain.User
}

func (u *UserRepositoryMock) FindByLogin(login string) (*domain.User, error) {
	return u.UserToReturn, u.ErrorToReturn
}

func (u *UserRepositoryMock) FindByUniqueAttributes(login, email, phone string) (*domain.User, error) {
	u.RequestedEmail = email
	u.RequestedLogin = login
	u.RequestedPhone = phone
	return u.UserToReturn, u.ErrorToReturn
}

func (u *UserRepositoryMock) Insert(user *domain.User) error {
	u.InsertedUser = user
	return u.ErrorDuringInsert
}
