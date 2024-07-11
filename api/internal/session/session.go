package session

import "github.com/Piokor/olutek_lib/internal/database"

func StartSession(userId int64, dbService *database.DbService) (*Session, error) {
	session := New(userId)
	err := insert(session, dbService)
	if err != nil {
		return nil, err
	}
	return session, nil
}
