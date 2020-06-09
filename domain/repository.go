package domain

type UserRepository interface {
	FindByLogin(login string) (*User, error)
	FindByUniqueAttributes(login, email, phone string) (*User, error)
	Insert(*User) error
}
