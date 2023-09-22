package repository

import (
	"database/sql"
	"fmt"

	structs "rest"
)

type AuthPost struct {
	db *sql.DB
}

func NewAuthPost(db *sql.DB) *AuthPost {
	return &AuthPost{db: db}
}

func (s *AuthPost) CreateUser(user structs.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name,username,password_hash) VALUES ($1,$2,$3) RETURNING id", usersTable)
	row := s.db.QueryRow(query, user.Name, user.UserName, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}
