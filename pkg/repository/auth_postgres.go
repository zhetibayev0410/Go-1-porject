package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	GO_Project "github.com/zhetibayev0410/Go-1-project"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user GO_Project.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3)RETURNING id")

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

//func (r *AuthPostgres) GetUser(username, password string) (GO_Project.User, error) {
//	var user GO_Project.User
//	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
//	err := r.db.Get(&user, query, username, password)
//
//	return id, err
//}
