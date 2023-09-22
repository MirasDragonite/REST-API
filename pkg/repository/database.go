package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	usersTable     = "users"
	todoListsTable = "todo_lists"
	userListsTable = "users_lists"
	todoItemsTable = "todo_items"
	listsItemTable = "lists_item"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	sts := ` DROP TABLE IF EXISTS lists_items;
	DROP TABLE IF EXISTS todo_items;
	DROP TABLE IF EXISTS users_lists;
	DROP TABLE IF EXISTS todo_lists;
	DROP TABLE IF EXISTS users;
	CREATE TABLE users(id INTEGER PRIMARY KEY, name TEXT, username TEXT, password_hash TEXT);
	CREATE TABLE todo_lists(id INTEGER PRIMARY KEY, title TEXT, description TEXT);
	CREATE TABLE users_lists(id INTEGER PRIMARY KEY, user_id INTEGER , list_id INTEGER,FOREIGN KEY (user_id) REFERENCES users (id), FOREIGN KEY (list_id) REFERENCES todo_lists (id));
	CREATE TABLE todo_items(id INTEGER PRIMARY KEY, title TEXT, description TEXT, done INTEGER DEFAULT 0 );
	CREATE TABLE lists_items(id INTEGER PRIMARY KEY, item_id INTEGER , list_id INTEGER ,FOREIGN KEY (item_id) REFERENCES todo_items (id), FOREIGN KEY (list_id) REFERENCES todo_lists (id));

	`
	_, err = db.Exec(sts)
	if err != nil {
		return nil, err
	}
	return db, nil
}
