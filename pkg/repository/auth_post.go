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

func (s *AuthPost) GetUser(username, password string) (structs.User, error) {
	var user structs.User
	fmt.Println(username, password)
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)

	row := s.db.QueryRow(query, username, password)

	err := row.Scan(&user.Id)
	fmt.Println(user.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were found for the given username and password.
			return structs.User{}, fmt.Errorf("user not found")
		}
		return structs.User{}, err
	}

	return user, nil
}
