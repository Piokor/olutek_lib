package users

import (
	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/Piokor/olutek_lib/internal/session"
	"github.com/google/uuid"
)

func RegisterUser(username, password string, dbService *database.DbService) (*User, *uuid.UUID, error) {
	user, err := InsertUser(username, password, dbService)
	if err != nil {
		return nil, nil, err
	}
	session, err := session.StartSession(user.Id, dbService)
	if err != nil {
		return nil, nil, err
	}
	return user, &session.Id, nil
}

func LoginUser(username, password string, dbService *database.DbService) (*User, *uuid.UUID, error) {
	user, err := SelectUser(username, password, dbService)
	if err != nil {
		return nil, nil, err
	}
	session, err := session.StartSession(user.Id, dbService)
	if err != nil {
		return nil, nil, err
	}
	return user, &session.Id, nil
}
