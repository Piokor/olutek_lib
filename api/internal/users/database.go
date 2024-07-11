package users

import "github.com/Piokor/olutek_lib/internal/database"

func InsertUser(username, password string, dbService *database.DbService) (*User, error) {
	var id int64
	insertStr := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err := dbService.Client.QueryRow(insertStr, username, password).Scan(&id)
	if err != nil {
		return nil, err
	}
	user := User{Id: id, Username: username}
	return &user, nil
}

func SelectUser(username, password string, dbService *database.DbService) (*User, error) {
	queryStr := `
		SELECT (id, username) FROM users 
		WHERE username=$1 AND password=$2
	`
	user := User{}
	err := dbService.Client.Get(&user, queryStr, username, password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
