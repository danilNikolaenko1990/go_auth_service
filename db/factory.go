package db

import "auth_service/domain"

type UserRepositoryFactory struct{}

func (u UserRepositoryFactory) CreateUserRepo(cs ConnectionSettings) domain.UserRepository {
	dbConnection := GetConnectionInstance(cs)
	return NewPgUserRepository(dbConnection)
}
