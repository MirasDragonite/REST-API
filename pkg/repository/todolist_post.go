package repository

import (
	"database/sql"
	"fmt"

	structs "rest"
)

type TodoListPost struct {
	db *sql.DB
}

func NewToDoListPost(db *sql.DB) *TodoListPost {
	return &TodoListPost{db: db}
}

func (r *TodoListPost) Create(userId int, list structs.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1,$2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", userListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}
	return id, tx.Commit()
}

func (r *TodoListPost) GetAll(userId int) ([]structs.TodoList, error) {
	var lists []structs.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id WHERE ul.user_id=$1", todoListsTable, userListsTable)
	row, err := r.db.Query(query, userId)

	for row.Next() {
		var list structs.TodoList

		err = row.Scan(&list.Id, &list.Title, &list.Description)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, err
}
