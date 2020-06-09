package domain

type User struct {
	TableName    struct{} `sql:"app_user"`
	Id           int64    `sql:",pk"`
	Login        string
	Email        string
	Phone        string
	PasswordHash string
	CreatedAt    string
}
