package users

import (
	"github.com/Piokor/olutek_lib/internal/database"
)

func CreateUser(username, password string, dbService *database.DbService) (*User, error) {
	user, err := InsertUser(username, password, dbService)
	if err != nil {
		return nil, err
	}
	return user, nil
}
