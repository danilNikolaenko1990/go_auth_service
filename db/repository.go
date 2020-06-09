package db

import (
	"auth_service/domain"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type PgUserRepository struct {
	conn *pg.DB
}

func NewPgUserRepository(connection *pg.DB) domain.UserRepository {
	return &PgUserRepository{conn: connection}
}

func (r *PgUserRepository) FindByLogin(login string) (*domain.User, error) {
	user := &domain.User{}
	err := r.conn.Model(user).Where("login=?", login).Select()
	return user, err
}

func (r *PgUserRepository) FindByUniqueAttributes(login, email, phone string) (*domain.User, error) {
	user := &domain.User{}
	err := r.conn.Model(user).
		WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			q = q.WhereOr("login = ?", login).
				WhereOr("email = ?", email).
				WhereOr("phone = ?", phone)
			return q, nil
		}).Select()

	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *PgUserRepository) Insert(user *domain.User) error {
	return r.conn.Insert(user)
}
