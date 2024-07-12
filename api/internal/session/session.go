package session

import (
	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/google/uuid"
)

func StartSession(userId int64, dbService *database.DbService) (*Session, error) {
	session := New(userId)
	err := insert(session, dbService)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func GetUserId(sessionId uuid.UUID, dbService *database.DbService) (int64, error) {
	return selectUserId(sessionId, dbService)
}
